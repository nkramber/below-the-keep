# Gate 2: the terminal

## Goal

The game takes the terminal, and it gives the terminal back in its first state on each exit path. This gate makes `go.mod`, the package `internal/term`, and the skeleton of `main.go`. The skeleton follows the error policy of D-55 and D-56. The game itself comes in later gates.

## Scope

In:

- `go.mod`: the module path `github.com/nkramber/below-the-keep` (D-42), and the `go` line of OQ-17.
- `internal/term` (P-29): the raw mode on and off, the alternate screen on and off, the cursor off and on, the window size, the resize events, and the type `term.Error` (P-31).
- `main.go`: the set-up in a clear order (P-36), the restore of the terminal, and the exit path for an error and for a panic (D-55, D-56).
- A temporary demo in `main.go`. It shows the window size and waits for one key. Gate 6 replaces it.
- `.gitignore`: the binary `below-the-keep`. The agent writes this file, because it is tool config.
- The license (OQ-21). The license file is not on the lists of hard rule 1. So the owner adds it, or the owner first adds it to the docs list.

Out:

- The logs (gate 3).
- The key reader goroutine and the key parser (gate 5).
- The game loop, and the signals `SIGTERM` and `SIGHUP` (gate 6).

## Success criteria

The decisions and the checks:

- [ ] `docs/decisions.md` has a row for OQ-17 (the `go` line), OQ-21 (the license), and OQ-24 (the exit codes).
- [ ] `go.mod` holds the module path `github.com/nkramber/below-the-keep`, the `go` line of OQ-17, and no `require` line (D-21).
- [ ] The owner ran the five checks of hard rule 5. `go vet`, `gofmt -l`, and `staticcheck` print nothing. `go build` and `go test` pass.
- [ ] `.gitignore` holds the binary. After `go build`, `git status` does not show `below-the-keep`.

The package `term`:

- [ ] Only `internal/term` imports `syscall` or `unsafe` (D-21, P-29).
- [ ] `internal/term` imports no package of the project (P-30).
- [ ] `term.Error` holds the operation and the cause, and it has an `Unwrap` method (D-56, P-31). A test shows that `errors.Is` finds the cause. The same test shows that the text has the form "term: operation: cause".
- [ ] `term` gives the window size in columns and rows. A manual check: the demo shows the same numbers as `stty size`, which prints the rows first.
- [ ] `term` gives a channel of resize events. A test sends `SIGWINCH` to its own process, and it receives one event.

The exit paths. Do these manual checks in a Terminal.app window:

- [ ] **A normal quit.** Run `stty -g` before the demo and after it. The two outputs are the same. After the demo, the cursor shows, and the old screen of the shell is back.
- [ ] **Ctrl-C.** In the demo, Ctrl-C ends the demo through the normal quit path. `stty -g` does not change.
- [ ] **An error before the raw mode.** `./below-the-keep < /dev/null` prints one line of the form "term: operation: cause" to stderr. `echo $?` prints the error code of OQ-24. `stty -g` does not change.
- [ ] **A panic after the raw mode.** Add a temporary `panic` after the raw mode, and run the game. The terminal comes back as after a normal quit. Stderr shows the panic value and the stack. `echo $?` prints the panic code of OQ-24. Then remove the temporary line.

The skeleton of `main`. The agent checks these items in the review:

- [ ] `main` does the set-up in a clear order, and the code has no `init` function (P-36).
- [ ] Only `main` calls `os.Exit`, and only after the restore (D-56).
- [ ] A deferred `recover` in `main` restores the terminal and prints the stack (D-55).
- [ ] The limits of `docs/code-standards.md` hold (D-57, D-58).
- [ ] Each exported name has a doc comment, and each package has a package comment (P-39).

## Open questions

- OQ-17: the `go` line of `go.mod`.
- OQ-21: the license.
- OQ-24: the exit code for an error, and the exit code for a panic.

## Decisions

None yet.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Planned.
