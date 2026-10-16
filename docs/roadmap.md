# Roadmap

This file lists each gate, in order. A gate is one roadmap stage with success criteria. Each gate has one gate file in `docs/gates/` and one pull request (D-10, D-25).

## The rules of the order

- A gate starts only after the gate before it merges.
- A change of the order needs an owner decision in `docs/decisions.md`.
- Gate 1 writes the list of the code gates. Until then, the rows after gate 1 are empty (D-17).

## The gates

| Gate | Name | Kind | Status | Gate file |
|---|---|---|---|---|
| 0 | The process and the session rules | docs | merged | `docs/gates/gate-00-process.md` |
| 1 | The game design and the roadmap | docs | planned | `docs/gates/gate-01-design.md` |
| 2 and later | Gate 1 sets them | code | planned | Gate 1 writes them |

The Status column has three values: planned, active, and merged. The pull request of a gate writes "merged" before it opens (D-32).
