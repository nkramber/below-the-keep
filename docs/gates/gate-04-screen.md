# Gate 4: the screen buffer

## Goal

The game draws into a grid of cells, and a render writes only the cells that changed. Each cell has a glyph, a foreground color, and a background color from the 256 colors. This gate makes the package `internal/screen`.

## Scope

In:

- `internal/screen` (P-29): the grid of cells, the render to an `io.Writer`, a box with borders of box-drawing characters, and a line of text at a position.
- A new grid size after a resize, with a full redraw.
- A temporary demo in `main.go`: a box and the colors of P-24. Gate 6 replaces it.

Out:

- The layout of the game screen (gate 6), and the draw of the map (gate 10).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A test: the first render of a new grid writes each cell.
- [ ] A test: a second render with no change writes no byte.
- [ ] A test: a render after one changed cell writes the move to that cell and that cell alone.
- [ ] A test: a cell writes its foreground and its background as 256-color escape sequences (D-44, D-49).
- [ ] A test: a box has the borders of P-23 at the correct cells.
- [ ] A test: a line of text stops at the edge of the grid.
- [ ] A test: a new size makes the next render write each cell.
- [ ] A manual check: the demo shows the box and the colors of P-24 in Terminal.app, with the background on each cell.
- [ ] `internal/screen` imports no package of the project (P-30).
- [ ] The owner ran the five checks of hard rule 5, and the output shows no failure.
- [ ] The limits of `docs/code-standards.md` hold, and each exported name has a doc comment.

## Open questions

None yet.

## Decisions

None yet.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Planned.
