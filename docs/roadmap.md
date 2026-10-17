# Roadmap

This file lists each gate, in order. A gate is one roadmap stage with success criteria. Each gate has one gate file in `docs/gates/` and one pull request (D-10, D-25).

## The rules of the order

- A gate starts only after the gate before it merges.
- A change of the order needs an owner decision in `docs/decisions.md`.
- Gate 1 wrote the list of the code gates (D-17, D-73).

## The gates

| Gate | Name | Kind | Status | Goal | Gate file |
|---|---|---|---|---|---|
| 0 | The process and the session rules | docs | merged | The rules of the owner and the agent, and the procedures of a session. | `docs/gates/gate-00-process.md` |
| 1 | The game design and the roadmap | docs | merged | The game design, the code standards, and the list of the code gates. | `docs/gates/gate-01-design.md` |
| 2 | The terminal | code | planned | The game takes the terminal and gives it back in its first state on each exit path. | `docs/gates/gate-02-terminal.md` |
| 3 | The logs | code | planned | Each run writes a log file, and the logs never hold more than 10 MB. | `docs/gates/gate-03-logs.md` |
| 4 | The screen buffer | code | planned | A grid of cells with 256 colors, and a render that writes only the changed cells. | `docs/gates/gate-04-screen.md` |
| 5 | The input | code | planned | A key reader turns the bytes of the terminal into key events. | `docs/gates/gate-05-input.md` |
| 6 | The game loop | code | planned | The scenes and the loop, with resizes, a pause for a small window, and a clean quit. | `docs/gates/gate-06-loop.md` |
| 7 | The content | code | planned | The JSON content loads at the start, and each bad file gives a clear error. | `docs/gates/gate-07-content.md` |
| 8 | The world | code | planned | The entity model: the map, the actors, the items, the hero, and a floor. | `docs/gates/gate-08-world.md` |
| 9 | The cave maker | code | planned | One seed makes the same 5 cave floors, with items and monsters. | `docs/gates/gate-09-cave.md` |
| 10 | The walk | code | planned | The hero walks on the play screen and goes down the 5 floors. | `docs/gates/gate-10-walk.md` |
| 11 | The fights | code | planned | Attacks, the monster AI, the death of the hero, and the end screen for a loss. | `docs/gates/gate-11-fights.md` |
| 12 | The items | code | planned | Pick up, drop, equip, use, the drops, and the regen. | `docs/gates/gate-12-items.md` |
| 13 | The levels and the win | code | planned | Experience points, the level menu, the Warden, the artifact, and the win. | `docs/gates/gate-13-levels.md` |
| 14 | The balance and the release | code | planned | Play tests, tuned numbers, the instructions to play, and the first release. | `docs/gates/gate-14-release.md` |

The Status column has three values: planned, active, and merged. The first session of a gate writes "active" on the gate branch (D-35). The pull request of a gate writes "merged" before it opens (D-32). So `main` shows only "planned" or "merged".
