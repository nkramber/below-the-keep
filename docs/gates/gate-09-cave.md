# Gate 9: the cave maker

## Goal

The map source makes each floor as a cave, with items and monsters on it. One seed always gives the same 5 floors. This gate makes the package `internal/cave`.

## Scope

In:

- `internal/cave` (P-29): the cave maker of P-19, and the check of the connections (D-68).
- The start tile and the stairs (P-1), and the artifact and the Warden on floor 5.
- The items of P-12 and the monsters of P-17, from the map source alone (D-53).

Out:

- The draw of the map (gate 10).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A test with fixed seeds: each border tile is a wall (P-18).
- [ ] A test with fixed seeds: each open tile connects to the start tile in 4 directions (D-68, D-69).
- [ ] A test with fixed seeds: the open area is 40% or more of the map (P-19).
- [ ] A test: the same seed gives the same 5 floors, two times in a row (D-53).
- [ ] A test on a small fixed map: the stairs are on the open tile with the longest path from the start (P-1).
- [ ] A test: floor 5 has the artifact and the Warden next to it, and no stairs (P-1).
- [ ] A test: each floor has 5 items and the number of monsters of P-17, from the kinds of that floor or of the floor above it.
- [ ] A test: no item and no monster is on a wall or on the start tile, and no tile holds two monsters.
- [ ] A golden test compares the map of one fixed seed with a text file in `testdata` (P-37).
- [ ] `internal/cave` imports only `world` and `content` (P-30).
- [ ] The owner ran the five checks of hard rule 5, and the output shows no failure.
- [ ] The limits of `docs/code-standards.md` hold, and each exported name that does not state its whole job has a doc comment (D-89).

## Open questions

None yet.

## Decisions

None yet.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Planned.
