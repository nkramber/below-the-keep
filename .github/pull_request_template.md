Gate NN: the name of the gate. Write the one concern of this pull request in two to four sentences.

## Gate

- Gate file: `docs/gates/gate-NN-<slug>.md`
- Branch: `gate-NN-<slug>`
- Base: `0000000`

## Success criteria

Copy the checklist of the gate file. Check each item, or name the decision that moved it.

## Findings

Name each finding of the gate file and its status: fixed, kept, or moved. Write `none` when the review found no finding.

## Checks

Paste the output of `go build ./...`, `go vet ./...`, `go test ./...`, and `gofmt -l .`. A gate of docs alone writes `Not applicable: docs alone`.

## Documentation

Start each entry with `Changed:` or `No change:`, and give a reason.

| File | Entry |
|---|---|
| hand-off | Changed: `docs/SESSION_HANDOFF.md` <reason> |
| gate file | Changed: `docs/gates/gate-NN-<slug>.md` <reason> |
| roadmap | <status> `docs/roadmap.md` <reason> |
| decisions | <status> `docs/decisions.md` <reason> |
| questions | <status> `docs/open-questions.md` <reason> |
| rules | <status> `AGENTS.md` <reason> |
| guard | <status> `.claude/settings.json` <reason> |
