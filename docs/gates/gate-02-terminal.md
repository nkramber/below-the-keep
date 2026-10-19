# Gate 2: the terminal

## Goal

The game takes the terminal, and it gives the terminal back in its first state on each exit path. This gate makes `go.mod`, the package `internal/terminal`, and the skeleton of `main.go`. The skeleton follows the error policy of D-55 and D-56. The game itself comes in later gates.

## Scope

In:

- `go.mod`: the module path `github.com/nkramber/below-the-keep` (D-42), and the line `go 1.27.2` (D-74).
- `internal/terminal` (P-29, D-79): the raw mode on and off, the alternate screen on and off, the cursor off and on, the window size, the resize events, and the type `terminal.Error` (P-31).
- `main.go`: the set-up in a clear order (P-36), the restore of the terminal, and the exit path for an error and for a panic (D-55, D-56).
- A temporary demo in `main.go`. It shows the window size and waits for one key. Gate 6 replaces it.
- `.gitignore`: the binary `below-the-keep`. The agent writes this file, because it is tool config.
- `LICENSE`: the MIT license (D-75). The owner writes this file (D-76).

Out:

- The logs (gate 3).
- The key reader goroutine and the key parser (gate 5).
- The game loop, and the signals `SIGTERM` and `SIGHUP` (gate 6).

## Success criteria

The decisions and the checks:

- [x] `docs/decisions.md` has a row for OQ-17 (D-74), OQ-21 (D-75), and OQ-24 (D-77).
- [ ] `go.mod` holds the module path `github.com/nkramber/below-the-keep`, the line `go 1.27.2` (D-74), and no `require` line (D-21).
- [ ] The owner ran the five checks of hard rule 5. `go vet`, `gofmt -l`, and `staticcheck` print nothing. `go build` and `go test` pass.
- [ ] `.gitignore` holds the binary. After `go build`, `git status` does not show `below-the-keep`.
- [x] `LICENSE` holds the MIT text. Its copyright line holds a name and no email address (D-29, D-75, D-76, D-78).

The package `terminal`:

- [ ] Only `internal/terminal` imports `syscall` or `unsafe` (D-21, P-29).
- [ ] `internal/terminal` imports no package of the project (P-30).
- [x] `terminal.Error` holds the operation and the cause, and it has an `Unwrap` method (D-56, P-31). A test shows that `errors.Is` finds the cause. The same test shows that the text has the form "terminal: operation: cause".
- [ ] `terminal` gives the window size in columns and rows. A manual check: the demo shows the same numbers as `stty size`, which prints the rows first.
- [ ] `terminal` gives a channel of resize events. A test sends `SIGWINCH` to its own process, and it receives one event.

The exit paths. Do these manual checks in a Terminal.app window:

- [ ] **A normal quit.** Run `stty -g` before the demo and after it. The two outputs are the same. After the demo, the cursor shows, and the old screen of the shell is back.
- [ ] **Ctrl-C.** In the demo, Ctrl-C ends the demo through the normal quit path. `stty -g` does not change.
- [ ] **An error before the raw mode.** `./below-the-keep < /dev/null` prints one line of the form "terminal: operation: cause" to stderr. `echo $?` prints `1` (D-77). `stty -g` does not change.
- [ ] **A panic after the raw mode.** Add a temporary `panic` after the raw mode, and run the game. The terminal comes back as after a normal quit. Stderr shows the panic value and the stack. `echo $?` prints `70` (D-77). Then remove the temporary line.

The skeleton of `main`. The agent checks these items in the review:

- [ ] `main` does the set-up in a clear order, and the code has no `init` function (P-36).
- [ ] Only `main` calls `os.Exit`, and only after the restore (D-56).
- [ ] A deferred `recover` in `main` restores the terminal and prints the stack (D-55).
- [ ] The limits of `docs/code-standards.md` hold (D-57, D-58).
- [ ] Each exported name that does not state its whole job has a doc comment (D-89), and each package has a package comment (P-39).

## Open questions

- OQ-17: the `go` line of `go.mod`. Answered by D-74.
- OQ-21: the license. Answered by D-75 and D-76.
- OQ-24: the exit code for an error, and the exit code for a panic. Answered by D-77.

## Decisions

D-74 to D-82, and D-84 to D-88. The error type of `terminal` follows D-84, D-85, and D-86.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|
| F-02-1 | minor | `internal/terminal/doc.go:1` | The first sentence uses a dash form. The Go form is "Package terminal" followed by a verb. | fixed |
| F-02-2 | minor | `internal/terminal/doc.go:1` | The agent flagged missing articles. Hard rule 6 covers chat, docs, and questions, not code comments. The finding was wrong. | kept |
| F-02-3 | minor | `internal/terminal/doc.go:2` | The comment gives the `syscall` rule now. It does not say how a failure is reported (D-56). | fixed |
| F-02-4 | minor | `internal/terminal/doc.go:2` | "It imports no other package" is false. The package imports standard library packages. | fixed |
| F-02-5 | minor | `internal/terminal/doc.go:1` | The comment has three sentences. P-39 asks for one or two. The agent told the owner to add the error rule, so the agent shares the cause. | fixed |

## Status

Active.
