# Gate 1: the game design and the roadmap

## Goal

The owner approves the game style, the gameplay, and the rules of the bones of the code. Then the agent writes the list of the code gates, with one gate file for each. This gate holds no code (D-17).

## Scope

In:

- A new doc, `docs/game-design.md`: the style, the gameplay, the screen, and the content.
- A new doc, `docs/code-standards.md`: the entity model, the error policy, the test policy, and the numbers of the review checklist.
- `docs/roadmap.md`: the list of the code gates.
- One gate file in `docs/gates/` for each code gate.
- The checklist of `AGENTS.md`: a link to `docs/code-standards.md`.
- The guard: deny rules for the paths of the game data (OQ-9).

Out:

- Each Go file and `go.mod`. The first code gate owns them.

## The order of the work

Ask the questions in this order. A later answer rests on an earlier one.

1. The game: OQ-2, OQ-3, OQ-4, OQ-23.
2. The screen: OQ-5, OQ-6, OQ-7, OQ-19.
3. The content and the state: OQ-8, OQ-9, OQ-10.
4. The bones: OQ-11, OQ-12, OQ-13, OQ-14, OQ-18.
5. The tools: OQ-15, OQ-16, OQ-22.
6. The list of the code gates.

## A draft of the code gates

This list is a draft for discussion. The owner approved none of it. The bones come first, and each gate is small enough for one pull request.

1. The terminal: the raw mode on and off, the alternate screen, and the restore on each exit and after a panic.
2. The screen buffer: a grid of cells, and a render that writes only the changed cells.
3. The input: read the keys, parse the escape sequences, and make key events.
4. The game loop: the scenes, and the split of update and draw.
5. The entity model: per OQ-11.
6. The gameplay gates: per OQ-2 to OQ-4, for example the map, the fights, and the items.
7. The save and the load: per OQ-8.
8. The content, the balance, and the release.

## Success criteria

- [ ] Each question of "The order of the work" has a row in `docs/decisions.md`, or the owner moved it to a named gate.
- [ ] `docs/game-design.md` exists, and the owner accepted it.
- [ ] `docs/code-standards.md` exists, and it gives each number of OQ-13.
- [ ] The checklist of `AGENTS.md` links to `docs/code-standards.md`.
- [ ] `docs/roadmap.md` lists each code gate with its kind and a goal of one line.
- [ ] Each code gate has a gate file with success criteria that the owner can check.
- [ ] The gate file of gate 2 is complete, and gate 2 can start.
- [ ] The guard covers the paths of the game data, or OQ-9 put the content in Go code.
- [ ] "Next steps, in order" of `docs/SESSION_HANDOFF.md` names gate 2.

## Open questions

OQ-2 to OQ-16, OQ-18, OQ-19, OQ-22, and OQ-23.

## Decisions

None yet.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Planned.
