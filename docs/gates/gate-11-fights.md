# Gate 11: the fights

## Goal

The hero and the monsters attack each other, and the monsters act on each turn. When the hero dies, the end screen shows the loss. This gate adds the fights and the monster AI to `internal/rules`.

## Scope

In:

- The hit roll and the damage roll of P-5, from the play source (D-53).
- The monster turn in the order of P-3, the sight of P-13, the chase of P-14, and the wander of P-15.
- The death of a monster, and the death of the hero.
- The end screen for a loss (P-26): the cause of death, the floor, the turns, and the seed.
- The messages of the fights, with the text of the content (P-27).

Out:

- The experience points of a kill (gate 13), and the drops (gate 12).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A table-driven test: the hit chance is min 5 and max 95 (P-5).
- [ ] A table-driven test: the damage of a hit is min 1 (P-5).
- [ ] A test: a fixed seed gives a fixed list of outcomes.
- [ ] A test: a wall blocks the sight of a monster, and a hero 9 tiles away is out of sight (P-13).
- [ ] A test on a small fixed map: a chase takes the first step of a shortest path. When another monster blocks the step, the monster waits (P-14).
- [ ] A test: the monsters act in the order of the monster list (P-3).
- [ ] A test: a monster with 0 hit points leaves the floor.
- [ ] A manual check: when the hero dies, the end screen shows the cause, the floor, the turns, and the seed (P-26).
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
