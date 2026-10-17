# Gate 12: the items

## Goal

The hero can pick up, drop, equip, and use items. Monsters drop items, and the hero gets back hit points over time. This gate adds the items to `internal/rules` and the letter choices to `internal/game`.

## Scope

In:

- The actions `g`, `d`, `e`, and `u` of P-25, and the cancel with Esc.
- The pack of P-10, and the first items of the hero (P-6).
- The effects of the potions and the scrolls (P-11).
- The drops, with the drop chance of each kind (P-16) and the play source (D-53).
- The regen of P-9.

Out:

- The level menu (gate 13).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A test: a pick up moves the item into the pack. A full pack refuses it with an outcome, not an error (P-10, D-55).
- [ ] A test: a drop puts the item on the tile of the hero.
- [ ] A test: an equip changes the damage range or the armor of the hero (D-41).
- [ ] A test: a potion heals up to the max hit points, and not above them.
- [ ] A test: a teleport scroll moves the hero to an open tile.
- [ ] A test: a fire scroll hits each monster within 3 tiles, and no monster further away.
- [ ] A test: the hero gets back 1 hit point every 10 turns, and not above the max (P-9).
- [ ] A test: a fixed seed gives a fixed list of drops.
- [ ] A manual check: a letter choice works, and Esc cancels it.
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
