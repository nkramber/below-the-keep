# Gate 7: the content

## Goal

The game content is in JSON files, and `embed` puts them in the binary. The loader turns the files into kinds, and each bad file gives a clear error before the game starts. This gate makes the package `internal/content`.

## Scope

In:

- `internal/content` (P-29): the JSON files, the `embed` line, the loader, and the kinds (D-51).
- The monster kinds of P-16, the item kinds of P-11, the messages of P-27, and the glyph list of P-23.
- `main.go`: the load of the content at the start (P-33).

Out:

- The use of the kinds in the world (gate 8).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] The owner wrote the JSON files in `internal/` (D-61). The files hold each monster kind of P-16 and each item kind of P-11.
- [ ] A test loads the embedded files with no error.
- [ ] A test checks each glyph of the content against the glyph list (D-48).
- [ ] A table-driven test with bad files in `testdata` covers these cases: bad JSON, a required field that is absent, an unknown field, a glyph that is not on the list, and a value out of its range.
- [ ] Each error of the loader is a `content.Error` that names the file and the field (P-31, P-33).
- [ ] A manual check: a bad file stops the game before the raw mode, with one clear line on stderr.
- [ ] `internal/content` imports no package of the project (P-30).
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
