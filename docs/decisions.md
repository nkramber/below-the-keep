# Decisions

This file holds each decision, with its date and its source. It is the source of truth. The other docs cite a decision by its `D-` id. Do not ask again a question that has a row here.

The Source column has three values:

- **owner statement:** the owner said it without a question.
- **owner answer:** the owner selected it in answer to a question of the agent.
- **session proposal:** the agent proposed it, and no question asked it. The merge of the gate that adds the row is the sign-off of the owner. The owner can change it at any time.

When a later decision changes an earlier one, keep both rows. Mark the ID cell of the old row "(amended by D-x)" or "(superseded by D-x)".

Take the next free number for a new row. Never reuse a number.

| # | Date | Source | Question | Decision |
|---|---|---|---|---|
| D-1 | 2026-10-16 | owner statement | The goals of the project | Three goals, in order. (1) A functional game with perfect bones: the entity structure, the error handling, and each other base part. (2) As few external libraries as possible. (3) The most readable code possible. |
| D-2 | 2026-10-16 | owner statement | The language | Go. |
| D-3 | 2026-10-16 | owner statement | The roles | The owner owns the repo and signs off on each decision. The agent is the project manager. It sets the gates and the success criteria of each gate. It proposes the game style and the gameplay, and the owner approves each one. The owner expects many questions from the agent, in each session. |
| D-4 | 2026-10-16 | owner statement | The number one rule | The agent never edits a file that holds production code. The agent gives no code example unless the owner asks for one. An example is generic, and it does not solve the current problem. It guides the owner in the correct direction, as bait on a hook guides a fish. A syntax question gets the generic form of the language feature. |
| D-5 | 2026-10-16 | owner statement | The purpose of a code review | Two purposes. (1) The code does its stated purpose. (2) The code obeys the rules of the owner: simple and readable code, small methods with helpers, minimal duplication, separation of concerns, and the other best practices. |
| D-6 | 2026-10-16 | owner statement | The work on a finding | For each bug, error, complexity problem, or other finding, the owner does the work, one finding at a time. |
| D-7 | 2026-10-16 | owner statement | The language of the chat | All chat output of the agent follows ASD-STE100. |
| D-8 | 2026-10-16 | owner answer | The language of the docs | All docs also follow ASD-STE100. One style everywhere is easier to read, and decktome does the same. |
| D-9 | 2026-10-16 | owner statement | The review tools | No Gitar. No review by a second AI provider. The agent of the session does each review. |
| D-10 | 2026-10-16 | owner statement | The session rules | One gate in each pull request. `AGENTS.md` and `docs/SESSION_HANDOFF.md` hold the session rules, in the style of the decktome repo. Each session can end with a prompt for the next session that names `docs/SESSION_HANDOFF.md`. |
| D-11 | 2026-10-16 | owner answer | The write scope of the agent | Docs and tool config. The docs are `AGENTS.md` and each `.md` file under `docs/`. The tool config is `.claude/settings.json`, the files under `.github/`, `.gitignore`, and `.editorconfig`. Each other file belongs to the owner: Go code, tests, `go.mod`, `go.sum`, and the game data. The rejected option: docs alone. |
| D-12 | 2026-10-16 | owner answer | The git and GitHub write steps | The owner does each write step: branch, commit, push, pull request, and merge. The agent uses git and gh to read alone. No change enters the history before the owner sees it in the diff. |
| D-13 | 2026-10-16 | owner answer | The form of the findings | First a short numbered list: the title and the severity of each finding, with no fix. Then the owner works on one finding at a time. The agent checks each fix before the next finding. |
| D-14 | 2026-10-16 | owner answer | The record of the findings | The gate file holds a findings table: id, severity, location, finding, and status. One gate, one file, one pull request. |
| D-15 | 2026-10-16 | owner answer | The checks | The owner runs `go build`, `go vet`, `go test`, and `gofmt -l`, and pastes the output. The agent reads the files and the diff alone. The rejected option: the agent runs read-only checks. |
| D-16 | 2026-10-16 | owner answer | The merge method | Squash. Each gate becomes one commit on `main`, and the history reads gate by gate. |
| D-17 | 2026-10-16 | owner answer | The order of the first gates | Gate 0 is the process and the session rules, with docs alone. Gate 1 is the game design and the roadmap, with docs alone. The code starts at gate 2. |
| D-18 | 2026-10-16 | owner answer | The first commit of an empty repo | The owner pushes one empty commit to `main`. The branch of gate 0 starts from it. Gate 0 then goes through a pull request, as each later gate does. |
| D-19 | 2026-10-16 | owner answer | The guard | `.claude/settings.json` holds deny rules for the agent. The Edit rules block the Edit, Write, and NotebookEdit tools on each `.go` file and on the module files. The Bash rules block the git and gh write commands. They also block the go commands that build, test, run, format, or change modules. The agent keeps the git reads, the gh reads, `go doc`, and `go version`. The rules apply to the agent alone, not to the terminal of the owner. Per the Claude Code docs of 2026-10-16, the rules do not block a script that writes a file. Hard rule 1 of `AGENTS.md` covers that gap. |
| D-20 | 2026-10-16 | owner answer | The operating systems | macOS alone. The rejected options: macOS and Linux, and macOS, Linux, and Windows. |
| D-21 | 2026-10-16 | owner answer | The external libraries (OQ-1) | The standard library alone. `go.mod` holds no `require` line. Research of 2026-10-16 on Go 1.27.2, darwin/arm64: `syscall` holds `TIOCGETA`, `TIOCSETA`, `TIOCGWINSZ`, `SIGWINCH`, and the `Termios` type. So the raw mode, the window size, and the resize signal need no module. `syscall` holds no type for the window size, so the code must define one. The calls need the `unsafe` package, in one small file for macOS. The doc of `syscall` asks new code to prefer `golang.org/x/sys`. The owner accepts this trade-off. The rejected option: `golang.org/x/term` v0.47.0, which needs `golang.org/x/sys` v0.49.0. The render, the input, the save files, the random numbers, the logs, and the tests all have a standard package. A width library for wide runes is not necessary when the glyph set stays narrow (OQ-6). |
| D-22 | 2026-10-16 | session proposal | The severity of a finding | Three levels. Blocker: the code fails its purpose, crashes, loses data, leaves the terminal broken, or breaks a hard rule. Major: a rule of the owner fails, for example duplication, a long function, mixed concerns, or an error path with no check. Minor: a name, a comment, an order, or a small point of readability. |
| D-23 | 2026-10-16 | session proposal | The status of a finding | Four values. Open: no fix yet. Fixed: the agent read the fix and found no problem. Kept: the owner keeps the code as it is, and the table holds the reason. Moved: a later gate owns it, and the table names that gate. |
| D-24 | 2026-10-16 | session proposal | The hint ladder | Three steps, one step up only on request of the owner. (1) A question that points to the problem. (2) The name of the concept, the pattern, or the standard package. (3) A generic example, on an explicit request alone, under D-4. |
| D-25 | 2026-10-16 | session proposal | The names of files and branches | The roadmap is `docs/roadmap.md`. Each gate file is `docs/gates/gate-NN-<slug>.md`, with a two-digit number. The branch of a gate is `gate-NN-<slug>`. A finding id is `F-NN-n`. |
| D-26 | 2026-10-16 | session proposal | Sessions and gates | A session works on one gate alone. A gate can use many sessions, because the owner writes the code and a gate can take days. A pause ends with a continue prompt. A merge ends with a start prompt. decktome uses the same model: many clean sessions can work on one pull request. |
| D-27 | 2026-10-16 | session proposal | The size of the hand-off | The hand-off keeps the three most recent session records. Each older record moves to `docs/session-archive.md`, word for word. |
| D-28 | 2026-10-16 | session proposal | The rules file of the harness | `AGENTS.md` is the one rules file. The repo holds no `CLAUDE.md`. Per the Claude Code docs of 2026-10-16, Claude Code reads `AGENTS.md` at the start of a session when no `CLAUDE.md` exists, from version 2.1.277. A `CLAUDE.md` with the line `@AGENTS.md` is the fallback. |
| D-29 | 2026-10-16 | session proposal | The public repo | The GitHub repo is public (read 2026-10-16). No file holds an email address, a personal address, or a secret. |
