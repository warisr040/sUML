# sUML (simply handle IT UML) — release notes

## Version 1.4
- **New name:** sUML now stands for **simply handle IT UML**, with a new icon: a small UML diagram with an “s” badge.
- View, Pen and Path are icon-only buttons (names show on hover).
- **Snip** (top bar, replaces the SVG / PNG buttons): works like Snipping Tool. Drag over part of the canvas, or click once for the whole view; the picture goes straight to the clipboard (paste with Ctrl+V) and a card offers Save PNG…, Save SVG…, Copy again and New snip. "Whole diagram" snips everything. Snips use the colours you see on screen and include visible path lines. Can also go on a shortcut slot.
- **Save As** can now save the whole diagram as a **PNG** or **SVG** picture (pick the type in the dialog); the diagram's own .suml.json file is not changed by that.
- **Tidier top bar:**
  - **View** menu (beside zoom) holds Fit, Grid, Snap and Smart, each with ON / OFF; a dot on the button shows Smart is on.
  - **Theme** moved into ⚙ Settings (with the picture background choice).
  - **Path** sits right of Pen: Pen · Path · Comment.
  - Labels fold away step by step on narrower windows so every button stays reachable at the exe's minimum width.

## Version 1.3
- **Groups:** select objects → Ctrl+G (or right-click → Group). Click any member to select the whole group and move it together; Alt+click picks one member; Ctrl+Shift+G ungroups. Lines between grouped shapes join the group. Copy / paste / duplicate makes a separate group.
- **Groups in Favorites:** right-click → *Add to Favorites as a group*. Click the favourite (or drag it onto the canvas, or put it on a shortcut slot) to insert a copy.
- **Text position** for every shape with a text box and every line label: Top / Mid / Bottom × Left / Center / Right (Properties panel). For lines, Left / Center / Right = near the start, middle or end; Top / Bottom = above or below the line.

## Version 1.2
- **Path lines** (⟿ Path in the top bar): trace a route through the diagram on its own layer, with its own window.
  - Click objects in order; the route follows existing lines between them, otherwise it finds its own way around objects.
  - One path line holds many scenarios; click one to show it. Shown / Hidden per path line and for all.
  - Effects: Running colour, Highlight, Blink, None. Six colour themes, each with two toned-down levels (Soft, Muted). Direction Forward / Reverse / Both, three speeds, numbered steps.
  - Objects on the route light up with the same highlight and theme.
  - Optional: include visible path lines in PNG / SVG export (still image).
- **Highlight** shapes and lines from the right-click menu (5 colours); **Clear all highlights** from any right-click menu.
- **Shortcut actions:** Draw line, Comment mode, Highlight on/off, Clear highlights, Path line on/off, Next scenario and Path lines window can go on Ctrl / Alt slots.
- **Draw line:** the end now snaps onto the destination's border, like the start.
- **Snap** puts shapes on the visible grid dots (20 units).
- **Comment pictures** open in a zoomable viewer (double-click in the box, or click the thumbnail): wheel / pinch zoom up to 16×, drag to pan, Fit, 100%.

## Version 1.1
- **Code panel** (button at the bottom right, or Ctrl+`): the diagram as editable text.
  - Shapes: `n1: rect "Start" @100,80 120x50 fill=#dbeafe layer=1`
  - Lines: `n1 -> n2 "label" route=ortho width=thick`  (arrows `->  -->  ..>  --  -|>  ..|>  <->  *--  o--`)
  - Apply (Ctrl+Enter) adds / changes / deletes shapes and lines; errors are listed by line number and nothing changes until fixed; Ctrl+Z undoes an Apply.
  - The text follows the diagram automatically while you aren't editing it.
- **Draw arrow fix:** the start of a drawn line now snaps onto the shape (side centres, centre, or the nearest point on its edge), with a green preview before you press.

## Version 1.0

Offline UML & architecture diagram tool. Single-file web app + Windows exe. Nothing is uploaded.

## Diagrams
- Shapes for General, Flowchart, Architecture (software / hardware), Use case, Dataflow (DFD), Activity, Class, Swimlanes
- Favorites (right-click a shape), foldable and re-orderable shape sections
- Tabs for several diagrams, unsaved-changes prompt, warning above 10 open diagrams
- 5 themes: Basic, Dark, Starry Night, Forest, Burning (export in white or theme colours)

## Lines
- Connect by dragging connection dots; free arrows; Draw arrow; Shift+drag detaches
- Attach anywhere on a box edge, centre or side centres; grid snap; joints (Ctrl+drag adds, double-click removes)
- Straight or right-angle routing; line hops at crossings
- Line size Thin / Medium / Thick; solid / dashed / dotted; UML heads
- Smart routing: stubs, avoid objects, avoid line overlap, fan-out of parallel lines, optional thicker overlapped lines
- Dragging a line moves the whole line

## Smart mode
- Connection window, place-next-and-connect (centre-aligned), fade unrelated objects (optional)
- Auto-connect loose line ends touched by a moved box
- Auto increment layer on add / move; lines take the smaller layer of their two boxes
- Smart OFF pauses all automatic behaviour

## Arrange
- Align (left / centre / right / top / middle / bottom), Distribute horizontally / vertically
- Auto-order a flow in 8 directions (3×3 picker)
- Auto-line alignment: Auto-route (keeps ends) and Straight
- Ctrl + drag a resize handle: symmetric resize around the centre

## Layers, marks, comments
- Object layers 0–1000, ink layers 0–50 (Layers panel), Move layer ±1 on right-click
- Important marks (lightning, !, check, cross, red / green / blue circle) and flow-order numbers 1–99 at a chosen corner
- Pen (freehand ink) with colour palette
- Comment mode: highlighted area + comment box with images (≤1000 px, paste screenshots) and styled text, drawing on comments, mark as done (Open / Closed), Comments panel with Shown / Hidden

## Files & app
- Save / Save As (.suml.json) with the system dialog, Open, SVG / PNG export
- Shortcut menu: Ctrl / Alt + 1–0, Q–T slots; config file next to the exe
- Keyboard shortcuts work on any keyboard layout (Thai, English, Chinese, Japanese, …)
- F1 help, F2 edit text
