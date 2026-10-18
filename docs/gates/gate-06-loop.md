# Gate 6: the game loop

## Goal

The game has scenes and one loop. The loop waits for a key, a resize, an error of the key reader, or a signal to end. Then it updates the scene and draws it. The demos of gates 2, 4, and 5 go away. This gate makes the first part of `internal/game`.

## Scope

In:

- `internal/game` (P-29): the scenes, the loop, and the layout of P-20 with empty parts.
- The title screen of P-26, an empty play screen, and the quit of P-25 with a yes or no question.
- The pause for a small window (D-47, P-20).
- The signals `SIGTERM` and `SIGHUP`: the game ends through the normal exit path.
- `main.go`: it connects `terminal`, `logs`, `screen`, `input`, and `game`.

Out:

- The map, the panel, and the messages of the play screen (gate 10).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A table-driven test gives a list of key events to the scenes, with no terminal. It checks the scene after each key.
- [ ] A manual check: the title screen shows the name, a new run, and quit (P-26).
- [ ] A manual check: a new run shows the empty play screen, with the borders of the map, the panel, and the messages (P-20).
- [ ] A manual check: `Q` asks a yes or no question. "Yes" goes back to the title screen. On the title screen, quit ends the game (P-25).
- [ ] A manual check: a window under 107 by 37 shows a pause message with the size that the game needs. A bigger window ends the pause (D-47, P-20).
- [ ] A manual check: a resize redraws the whole screen.
- [ ] A manual check: `kill` with `SIGTERM`, and the close of the Terminal.app tab, both end the game. After `SIGTERM`, `stty -g` does not change.
- [ ] A test: an error of the key reader reaches `main`, and `main` exits through the error path (D-55).
- [ ] `internal/game` imports only the packages of its row in P-30.
- [ ] No demo code stays in `main.go`.
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
