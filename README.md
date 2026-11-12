# sUML — smart UML · version 1.1

Offline UML / architecture diagram editor. Nothing is uploaded.

| Folder | What |
| --- | --- |
| `web/` | The whole app: `index.html` (single file, no dependencies) + `netlify.toml` for Netlify hosting |
| `desktop/` | Windows exe wrapper (Go + WebView2). `web_index.html` is a copy of `web/index.html` embedded at build time — see `BUILD.md` |
| `logo/` | App logo as SVG and PNG (16–512 px); `desktop/icon.ico` is the Windows icon |

Edit `web/index.html`, then copy it to `desktop/web_index.html` before rebuilding the exe.
