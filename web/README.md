# simply handle IT UML (sUML) — website

Static site, no build step, nothing loaded from other servers.

| File | What it is |
|---|---|
| `index.html` | Front page |
| `app.html` | The diagram editor (the whole app, one file) |
| `manual.html` | User manual |
| `design.html` | Design document |
| `shots/` | Screenshots used on the front page |
| `netlify.toml` | Security headers (`connect-src 'none'`) and the `/app` short link |

## Deploy to Netlify
- **Drag & drop:** open https://app.netlify.com/drop and drop this whole folder.
- **Git:** push the repo, "Add new site → Import", leave the build command empty, publish directory `web`.

Diagrams live in the visitor's browser until they **Save / Save As** (`.suml.json`), **SVG** or **PNG**.
