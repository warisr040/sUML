// smart UML (sUML) desktop — wraps the offline index.html in a native Windows window (WebView2).
// The page is served only on the loopback interface (127.0.0.1); nothing is sent to the internet.
package main

import (
	_ "embed"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/jchv/go-webview2"
)

//go:embed web_index.html
var page []byte

const (
	configName = "sUML.config.json" // exactly one config file, next to sUML.exe
	wmClose    = 0x0010
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	setWindowLongPt = user32.NewProc("SetWindowLongPtrW")
	callWindowProc  = user32.NewProc("CallWindowProcW")
	postMessage     = user32.NewProc("PostMessageW")
	messageBox      = user32.NewProc("MessageBoxW")

	oldProc    uintptr
	closeMu    sync.Mutex
	allowClose bool
	closeAcked bool
)

func listen() net.Listener {
	// Fixed port keeps browser storage (tabs / autosave) stable between runs; fall back if busy.
	if l, err := net.Listen("tcp", "127.0.0.1:47391"); err == nil {
		return l
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	return l
}

func configPath() string {
	exe, err := os.Executable()
	if err != nil {
		return configName
	}
	return filepath.Join(filepath.Dir(exe), configName)
}

func dirExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.IsDir()
}

func msg(text string) {
	t, _ := syscall.UTF16PtrFromString(text)
	c, _ := syscall.UTF16PtrFromString("sUML")
	messageBox.Call(0, uintptr(unsafe.Pointer(t)), uintptr(unsafe.Pointer(c)), 0x40)
}

func main() {
	ln := listen()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; connect-src 'none'")
		w.Write(page)
	})
	go http.Serve(ln, mux)

	dataDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "sUML")
	// Keep data from the old app name (SitUML): move its WebView2 profile and config once.
	if old := filepath.Join(os.Getenv("LOCALAPPDATA"), "SitUML"); dirExists(old) && !dirExists(dataDir) {
		_ = os.Rename(old, dataDir)
	}
	if exe, err := os.Executable(); err == nil {
		oldCfg := filepath.Join(filepath.Dir(exe), "SitUML.config.json")
		if _, e1 := os.Stat(configPath()); os.IsNotExist(e1) {
			if _, e2 := os.Stat(oldCfg); e2 == nil {
				_ = os.Rename(oldCfg, configPath())
			}
		}
	}
	_ = os.MkdirAll(dataDir, 0o755)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		DataPath:  dataDir,
		WindowOptions: webview2.WindowOptions{
			Title:  "smart UML (sUML) 1.1",
			Width:  1440,
			Height: 900,
			IconId: 2,
			Center: true,
		},
	})
	if w == nil {
		msg("sUML needs the Microsoft Edge WebView2 Runtime.\n\nInstall it from:\nhttps://developer.microsoft.com/microsoft-edge/webview2/")
		return
	}
	defer w.Destroy()
	hwnd := uintptr(w.Window())

	// --- settings file next to the exe (single file, overwritten on save) ---
	_ = w.Bind("sumlLoadConfig", func() string {
		b, err := os.ReadFile(configPath())
		if err != nil {
			return ""
		}
		return string(b)
	})
	_ = w.Bind("sumlSaveConfig", func(s string) string {
		p := configPath()
		tmp := p + ".tmp"
		if err := os.WriteFile(tmp, []byte(s), 0o644); err != nil {
			return "cannot write " + p + " (" + err.Error() + ")"
		}
		if err := os.Rename(tmp, p); err != nil {
			_ = os.Remove(tmp)
			return err.Error()
		}
		return ""
	})

	// --- close button: let the page ask about unsaved tabs first ---
	_ = w.Bind("sumlCloseAck", func() {
		closeMu.Lock()
		closeAcked = true
		closeMu.Unlock()
	})
	_ = w.Bind("sumlQuit", func() {
		closeMu.Lock()
		allowClose = true
		closeMu.Unlock()
		postMessage.Call(hwnd, wmClose, 0, 0)
	})
	cb := syscall.NewCallback(func(h, m, wp, lp uintptr) uintptr {
		if m == wmClose {
			closeMu.Lock()
			ok := allowClose
			closeAcked = false
			closeMu.Unlock()
			if !ok {
				w.Eval("window.__sumlRequestClose && window.__sumlRequestClose()")
				// Safety net: if the page doesn't answer (e.g. still loading), close anyway.
				go func() {
					time.Sleep(2 * time.Second)
					closeMu.Lock()
					acked := closeAcked
					if !acked {
						allowClose = true
					}
					closeMu.Unlock()
					if !acked {
						postMessage.Call(h, wmClose, 0, 0)
					}
				}()
				return 0
			}
		}
		r, _, _ := callWindowProc.Call(oldProc, h, m, wp, lp)
		return r
	})
	gwlpWndProc := ^uintptr(3) // GWLP_WNDPROC = -4
	oldProc, _, _ = setWindowLongPt.Call(hwnd, gwlpWndProc, cb)

	w.SetSize(1100, 700, webview2.HintMin)
	w.SetSize(1440, 900, webview2.HintNone)
	w.Navigate("http://" + ln.Addr().String() + "/")
	w.Run()
}
