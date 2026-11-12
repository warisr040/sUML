# sUML (smart UML) — release notes

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
