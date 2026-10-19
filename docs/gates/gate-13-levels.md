# Gate 13: the levels and the win

## Goal

Kills give experience points, and each new level opens the level menu. On floor 5, the Warden guards the artifact. When the hero picks up the artifact, the run ends with a win. With this gate, a full run is possible.

## Scope

In:

- The experience points of each kind (P-16), and the level curve of P-8.
- The level menu of P-7, as a box over the map, with the keys `1`, `2`, and `3`.
- The pick up of the artifact, and the end screen for a win (P-26).
- The help screen with the keys of P-25.

Out:

- The tune of the numbers (gate 14).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A table-driven test: the level curve gives levels at 10, 30, 60, and so on, up to 550 for level 11 (P-8).
- [ ] A test: a kill gives the experience points of the kind of the monster.
- [ ] A test: a gain that crosses two levels opens the level menu two times.
- [ ] A test: each choice of the level menu changes its stat as P-7 says.
- [ ] A test: the pick up of the artifact ends the run with a win.
- [ ] A manual check: the help screen lists each key of P-25.
- [ ] A manual check: a full run with a fixed seed reaches the end screen for a win.
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
