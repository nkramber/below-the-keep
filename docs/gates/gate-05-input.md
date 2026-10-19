# Gate 5: the input

## Goal

A key reader turns the bytes of the terminal into key events. The reader runs in the one extra goroutine of P-32, and it obeys the panic rule of D-56. This gate makes the package `internal/input`.

## Scope

In:

- `internal/input` (P-29): the key reader goroutine, and the parser of the bytes into key events.
- The keys of P-25: the arrow keys, the letters, the digits, `.`, `>`, `?`, Esc, and Ctrl-C.
- A temporary demo in `main.go`: it shows the name of each key. Gate 6 replaces it.

Out:

- The actions of the keys (gates 6, 10, 12, and 13).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] A table-driven test of the parser covers each key of P-25, and the arrow keys in both of their escape forms.
- [ ] A test: an Esc byte alone is the Esc key.
- [ ] A test: an unknown escape sequence gives no key and no error.
- [ ] A test: two keys in one read give two key events, in order.
- [ ] The reader goroutine defers a `recover`, and it sends a panic to `main` as an error (D-56).
- [ ] When the input closes, the reader sends this fact to `main`, and the goroutine ends.
- [ ] The reader reads from an `io.Reader`, so a test can give it bytes with no terminal.
- [ ] A manual check: the demo shows the correct name for each key of P-25.
- [ ] `internal/input` imports no package of the project (P-30).
- [ ] The owner ran the five checks of hard rule 5, and the output shows no failure.
- [ ] The limits of `docs/code-standards.md` hold, and each exported name that does not state its whole job has a doc comment (D-89).

## Open questions

None yet.

## Decisions

None yet.

## Findings

| ID | Severity | Location | Finding | Status |
|---|---|---|---|---|

## Status

Planned.
