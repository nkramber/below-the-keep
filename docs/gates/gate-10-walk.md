# Gate 10: the walk

## Goal

The play screen shows the map, the panel, and the messages. The hero walks in 4 directions and goes down the 5 floors. The monsters stand still until gate 11. This gate makes the first part of `internal/rules`.

## Scope

In:

- `internal/rules` (P-29): the move, the blocked move (P-2), the wait, and the stairs down.
- `internal/game`: the draw of the map with the glyphs of P-23 and the colors of P-24, the panel of P-21, and the messages of P-22.
- The seed: the flag, a random seed with no flag, the two sources of D-53, and the seed in the log.

Out:

- Attacks and the monster turn (gate 11).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A test: a move to an open tile moves the hero, and one turn passes.
- [ ] A test: a move into a wall returns the outcome "blocked", and no turn passes (P-2, D-55).
- [ ] A test: `>` on the stairs makes the next floor. `>` on another tile returns an outcome and no error.
- [ ] A manual check: a run with `-seed` shows the same floors two times in a row (D-53).
- [ ] A manual check: a run with no flag writes its seed to the log.
- [ ] A manual check: the panel shows the stats of P-6, the floor, and the turn (P-21).
- [ ] A manual check: the map uses the glyphs of P-23 and the colors of P-24.
- [ ] `internal/rules` imports only `world` and `content` (P-30).
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
