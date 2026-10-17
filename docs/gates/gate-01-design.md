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

## The code gates

The owner accepted 13 code gates, from gate 2 to gate 14 (D-73). `docs/roadmap.md` holds the list, and each gate has a gate file in `docs/gates/`.

## Success criteria

- [x] Each question of "The order of the work" has a row in `docs/decisions.md`, or the owner moved it to a named gate (D-36 to D-70, D-73).
- [x] `docs/game-design.md` exists, and the owner accepted it (D-71).
- [x] `docs/code-standards.md` exists, and it gives each number of OQ-13 (D-72).
- [x] The checklist of `AGENTS.md` links to `docs/code-standards.md` (owner check, 2026-10-16).
- [x] `docs/roadmap.md` lists each code gate with its kind and a goal of one line (owner check, 2026-10-16).
- [x] Each code gate has a gate file with success criteria that the owner can check (owner check, 2026-10-16).
- [x] The gate file of gate 2 is complete, and gate 2 can start (owner check, 2026-10-16).
- [x] The guard covers the paths of the game data, or OQ-9 put the content in Go code (D-61). `/permissions` showed 78 deny rules.
- [x] "Next steps, in order" of `docs/SESSION_HANDOFF.md` names gate 2.

## Open questions

OQ-2 to OQ-16, OQ-18, OQ-19, OQ-22, and OQ-23.

## Decisions

- D-35: the start check sets a gate to "active".
- D-36: the game style (OQ-2).
- D-37: the core loop (OQ-3).
- D-38: the combat model (OQ-4).
- D-39: the stats, the healing, the hunger, and the source of items.
- D-40: levels and experience points. This amends D-37.
- D-41: the level system. This amends D-39.
- D-42: the name of the game, and the module path (OQ-23).
- D-43: the glyphs (OQ-6).
- D-44: the colors (OQ-7).
- D-45: the size of the screen (OQ-5, OQ-19).
- D-46: the layout of the screen (OQ-5).
- D-47: the fit of the map, and the resize (OQ-19).
- D-48: the width rule of the glyphs.
- D-49: the background color.
- D-50: no save (OQ-8).
- D-51: the content in JSON files with `embed` (OQ-9).
- D-52: two seeded random sources (OQ-10).
- D-53: the split of the random sources. This amends D-52.
- D-54: the entity model (OQ-11).
- D-55: the error policy (OQ-12).
- D-56: the error type and the goroutines.
- D-57: the code limits and the package layout (OQ-13).
- D-58: two details of the limits (session proposal).
- D-59: the test policy (OQ-14).
- D-60: the logs (OQ-18).
- D-61: the deny rule for the data files.
- D-62: the folder, the size, and the level of the log.
- D-63: the split of the log into files.
- D-64: no CI (OQ-16).
- D-65: the README, and the docs list of hard rule 1 (OQ-22). This amends D-11.
- D-66: staticcheck as the fifth check (OQ-15). This amends D-15.
- D-67: the vision of the hero: the whole floor.
- D-68: the map maker: caves.
- D-69: 4 directions of movement.
- D-70: the guard rule for staticcheck.
- D-71: the owner accepted `docs/game-design.md`.
- D-72: the owner accepted `docs/code-standards.md`.
- D-73: the list of the code gates.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Merged (D-32). The status was "active" during the work (D-35).
