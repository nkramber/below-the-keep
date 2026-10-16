# Gate 0: the process and the session rules

## Goal

Set the rules of each session: the start, the work, the pause, the end of a gate, and the transition to the next gate. Write the base docs, the guard, and the template of a pull request. This gate holds no code (D-17).

## Scope

In:

- `AGENTS.md`
- `docs/SESSION_HANDOFF.md`
- `docs/decisions.md` and `docs/open-questions.md`
- `docs/roadmap.md`, this file, and `docs/gates/gate-01-design.md`
- `.claude/settings.json`, the guard (D-19)
- `.github/pull_request_template.md`

Out:

- The game design and the list of the code gates. Gate 1 owns them.
- Each Go file and `go.mod`.

## Success criteria

- [ ] `AGENTS.md` holds the goals, the roles, the hard rules, the session lifecycle, the review procedure, the glossary, and the file map.
- [ ] `docs/SESSION_HANDOFF.md` holds the resume point and the six procedures of a session.
- [ ] `docs/SESSION_HANDOFF.md` holds the template of the continue prompt and of the start prompt.
- [ ] `docs/decisions.md` holds each owner statement and each owner answer of 2026-10-16, with the date and the source.
- [ ] `docs/open-questions.md` holds each open question with the point at which to ask it.
- [ ] OQ-1 has its answer in D-21, and the answer gives the research behind it.
- [ ] The guard holds the deny rules of D-19, and `jq` reads it as valid JSON.
- [ ] Each doc follows ASD-STE100 (D-8).
- [ ] The owner read each file of the diff and accepted it.
- [ ] The owner accepted each session proposal, D-22 to D-29, or changed it.
- [ ] `main` holds one empty commit, and the pull request of gate 0 starts from it (D-18).

## Open questions

None. Gate 1 owns each open question of `docs/open-questions.md`, except OQ-17, OQ-20, and OQ-21.

## Decisions

D-1 to D-29.

## Findings

The owner reviews the docs of this gate. Each change that the owner asks for goes here.

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Active.
