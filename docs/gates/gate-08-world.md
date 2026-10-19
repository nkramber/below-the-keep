# Gate 8: the world

## Goal

The world holds the entity model of D-54: the map, the actors, the items, the hero, and a floor. It holds the state of the game, but no rule of the game. This gate makes the package `internal/world`.

## Scope

In:

- `internal/world` (P-29): the tiles and the map of 80 by 30 (P-18), `Actor`, `Item`, the hero type, and a floor.
- The pack of the hero with max 10 items, and the marks of the equipped items (P-10).
- The questions that the rules ask: the tile at a position, the actor at a position, and the items at a position.

Out:

- Each action and its outcome (gates 10 to 13).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] The types follow D-54. Each actor and each item holds its kind and its own state.
- [ ] A test: the pack refuses an eleventh item (P-10).
- [ ] A test: the armor of the hero is the armor value of the equipped armor, and the damage range is the range of the equipped weapon (D-41).
- [ ] A test: a position outside the map is not open, and a wall is not open.
- [ ] A test: the actor at a position, and the items at a position.
- [ ] The monster list of a floor keeps its order (P-3).
- [ ] No state at the package level (P-35).
- [ ] `internal/world` imports only `content` (P-30).
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
