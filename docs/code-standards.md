# The code standards of Below the Keep

This file holds the rules of the bones of the code: the libraries, the packages, the entity model, the error policy, the limits, the tests, and the logs. Each rule cites its decision in `docs/decisions.md`. A review checks the code against this file and against the checklist of `AGENTS.md`.

An item with a `P-` id is a proposal of the agent. The owner accepted this file and each proposal in it on 2026-10-16 (D-72).

## 1. The libraries (D-21)

- The code imports the standard library alone. `go.mod` holds no `require` line.
- The raw mode, the window size, and the resize signal come from `syscall`. The calls need `unsafe`, in one small file for macOS.
- `syscall` holds no type for the window size, so the code defines one.
- A new import from outside the standard library needs an owner decision first. A review marks it as a blocker.

## 2. The package layout (D-57)

- `main.go` is at the root of the module. The module path is `github.com/nkramber/below-the-keep` (D-42).
- Each other package is in `internal/`.
- The binary gets the name `below-the-keep`.

**P-29. The packages.** Each package has one job.

| Package | Its job |
|---|---|
| `internal/term` | The raw mode, the alternate screen, the window size, the resize signal, and the restore. The only package with `syscall` and `unsafe`. |
| `internal/screen` | A grid of cells: a glyph, a foreground, and a background. A render that writes only the changed cells to an `io.Writer`. |
| `internal/input` | It reads bytes from an `io.Reader` and parses them into key events. |
| `internal/content` | The JSON files, the `embed` line, the loader, the kinds of monsters and items, the messages, and the glyph list. |
| `internal/world` | The map, the actors, the items, and the state of a floor. The entity model of section 3. |
| `internal/cave` | The cave maker of P-19. It makes a floor of `world` from the map source. |
| `internal/rules` | Each action and its outcome: moves, attacks, items, levels, and the monster AI. |
| `internal/game` | The scenes, the loop, and the draw of the world into the grid of `screen`. |
| `internal/logs` | The log folder, one file for each run, and the limit of 10 MB. |

**P-30. The direction of the imports.** A package imports only the packages of its row. Go forbids a cycle, but this rule also keeps each concern in its part.

| Package | It can import |
|---|---|
| `term`, `screen`, `input`, `content`, `logs` | no package of the project |
| `world` | `content` |
| `cave` | `world`, `content` |
| `rules` | `world`, `content` |
| `game` | `rules`, `cave`, `world`, `content`, `screen`, `input` |
| `main` | each package |

So the rules never know the terminal, and the render never changes the world.

## 3. The entity model (D-54)

- The hero and each monster are an `Actor`. Each item is an `Item`.
- Each actor and each item holds its kind from the JSON content, and its own state, for example its current hit points.
- A hero type holds the hero `Actor` and the parts of the hero alone: the experience points, the level, the pack, and the equipped items.
- An interface appears only where two types share a behavior.

## 4. The error policy (D-55, D-56)

- Each package has its own error type. The type holds the name of the failed operation and the cause.
- The type has an `Unwrap` method, so `errors.Is` and `errors.As` still find the cause.
- A rule outcome is a result, not an error. For example, a move into a wall returns the outcome "blocked".
- An error is a failure of the program or of its environment, for example the terminal or a bad data file.
- Each error goes up to `main`. `main` restores the terminal, prints the error to stderr, and exits with a code that is not 0.
- A deferred `recover` in `main` does the same for a panic, and it also prints the stack.
- Each goroutine that the game starts defers a `recover`, and it sends the panic to `main` as an error.
- Only `main` calls `os.Exit`, and only after the restore of the terminal.

The reasons: a `recover` works only for a panic on the same goroutine (Go spec, "Handling panics"). `os.Exit` runs no deferred function (`go doc os.Exit`).

**P-31. The name and the text of the error type.** The type of each package has the name `Error`, so a caller reads `term.Error` or `content.Error`. Its text has the form "package: operation: cause".

**P-32. One extra goroutine.** The game has one goroutine next to `main`: the key reader of `input`. The resize signal needs no goroutine, because package `signal` relays it to a channel (`go doc os/signal.Notify`). A new goroutine needs an owner decision.

**P-33. The load of the content.** `main` loads the content at the start. A bad file is an error that names the file and the field. Then the game does not start.

## 5. The limits (D-57, D-58)

These limits apply to production code. Tests have no limit.

| Limit | Max |
|---|---|
| The lines of a function | 40 |
| The parameters of a function | 4 |
| The levels of nested blocks in a function | 3 |
| The lines of a file | 400 |

The length of a function counts each line between the braces of its body, also blank lines and comments (D-58).

**P-34. How to count the other limits.**

- The parameters: the receiver does not count. A variadic parameter counts as one.
- The levels: the body of the function is level 0. Each `if`, `else`, `for`, `switch`, `select`, `case`, and function literal adds one level to the code inside it.
- The lines of a file: each line, also blank lines and comments.

## 6. State and set-up

**P-35. No state at the package level.** A package-level variable is permitted only for a value that never changes after the start, for example the `embed` file system. The random sources go to each function that needs them as arguments.

**P-36. No `init` function.** `main` does each step of the set-up in a clear order.

## 7. The random sources (D-52, D-53)

- One seed makes two sources of `math/rand/v2`: the map source and the play source.
- The map source goes to `cave` and to the placement of items and monsters. The play source goes to `rules`.
- A test that uses random numbers uses a fixed seed.

## 8. The tests (D-59)

- These parts must have tests: the rules, the map maker, the content loader, the screen buffer, and the key parser.
- A function with many cases gets a table-driven test.
- The terminal gets a manual checklist in its gate file, because the raw mode needs a real terminal.
- The project has no coverage target. Each success criterion of a code gate names its tests. A review checks one test for each behavior of the gate.

**P-37. The place of the tests.** A test file is in the same package as the code that it tests. Fixed test files go in a `testdata` folder of the package. The go tool ignores a folder with that name (`go doc cmd/go`).

## 9. The logs (D-60, D-62, D-63)

- The game always writes a log, in `~/Library/Caches/below-the-keep/`. `os.UserCacheDir` gives the base folder.
- Each run writes a new file, named by its start time.
- At the start, the game deletes the oldest files until the total is under 10 MB (10 x 1024 x 1024 bytes).
- The log uses `slog.NewTextHandler` at the info level. A `-debug` flag adds the debug lines.

**P-38. One logger.** `main` makes the logger and sets it as the default with `slog.SetDefault`, one time, at the start. The other packages call the functions of `slog`.

## 10. The checks (D-15, D-66)

The owner runs these five checks before each commit (D-64), and pastes the output:

- `go build ./...`
- `go vet ./...`
- `go test ./...`
- `gofmt -l .`
- `staticcheck ./...`

The owner installs staticcheck with `go install honnef.co/go/tools/cmd/staticcheck@latest`, which does not change `go.mod`.

## 11. The comments

- Each exported name has a doc comment.

**P-39. The package comment.** Each package has one package comment. It gives the job of the package in one or two sentences.

## 12. Open for a code gate

- OQ-17: the `go` line of `go.mod`. The first code gate decides it.
- OQ-20: is a `Makefile` or a shell script tool config, or code of the owner?
- OQ-21: the license, before the first code gate merges.
