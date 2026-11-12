# smart UML (sUML)

Simple, offline UML / architecture diagram editor in a single `index.html`.
Nothing is uploaded — diagrams live in your browser until you **Save / Save As** (`.suml.json`), **SVG** or **PNG**.

Diagram kinds: architecture (software & hardware), use case, dataflow (DFD), activity, class, swimlanes.

## Deploy to Netlify
- **Drag & drop:** open https://app.netlify.com/drop and drop this whole folder.
- **Git:** push this folder to a repo, "Add new site → Import", leave the build command empty, publish directory `.`.

`netlify.toml` adds a Content-Security-Policy with `connect-src 'none'`, so the page cannot send data anywhere.

## Files
- `index.html` — the whole app (no dependencies)
- `netlify.toml` — Netlify config / security headers
