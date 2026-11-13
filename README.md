# sUML — simply handle IT UML · version 1.4

Offline UML / architecture diagram editor. Nothing is uploaded.

| Folder | What |
| --- | --- |
| `web/` | The website: `index.html` (front page), `app.html` (the whole editor, single file, no dependencies), `manual.html`, `design.html`, `shots/`, and `netlify.toml` for Netlify hosting |
| `desktop/` | Windows exe wrapper (Go + WebView2). `web_index.html` is a copy of `web/app.html` embedded at build time — see `BUILD.md` |
| `logo/` | App logo as SVG and PNG (16–512 px); `desktop/icon.ico` is the Windows icon |

Edit `web/app.html`, then copy it to `desktop/web_index.html` before rebuilding the exe.
