# Gate 3: the logs

## Goal

Each run writes a log file, and all the logs together never hold more than 10 MB. `main` sets the logger one time, so each later gate can write to the log. This gate makes the package `internal/logs`.

## Scope

In:

- `internal/logs` (P-29): the log folder, one file for each run, and the limit of 10 MB (D-62, D-63).
- `main.go`: the `-debug` flag, and the logger with `slog.SetDefault` (P-38).

Out:

- The seed in the log (gate 10).

## Success criteria

The first session of this gate reviews these criteria with the owner.

- [ ] `logs` makes the folder `below-the-keep` in the folder of `os.UserCacheDir` when the folder does not exist (D-62).
- [ ] Each run makes a new file. The name holds the start time, and the names sort in time order (D-63).
- [ ] At the start, `logs` deletes the oldest files until the total is under 10 MB (10 x 1024 x 1024 bytes) (D-63).
- [ ] A table-driven test in a temporary folder covers these cases: a total under the limit deletes nothing, and a total over the limit deletes the oldest file first. A file that does not belong to the game stays.
- [ ] The log uses `slog.NewTextHandler` at the info level. With `-debug`, the debug lines also appear (D-62). A manual check: run with and without the flag, and read the file.
- [ ] A failure to make the folder or the file is a `logs.Error` (P-31). `main` exits through the error path.
- [ ] `internal/logs` imports no package of the project (P-30).
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
