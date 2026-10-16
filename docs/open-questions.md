# Open questions

This file holds each question with no answer yet. When the owner answers a question, write the answer as a row in `docs/decisions.md`. Then move the question to "Answered" below, with the `D-` id.

Ask a question when the work reaches the point that needs the answer. Ask in batches of max four questions (hard rule 8 of `AGENTS.md`).

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-2 | What is the game style? For example: a roguelike dungeon crawl, a party RPG with turn-based fights, or a story RPG with menus. What tone, and what world? | Each later gate rests on it. | Gate 1, first. |
| OQ-3 | What is the core loop of the gameplay? What does the player do minute to minute, and what makes the player progress? How does a run end: a win, a death, or no end? | It sets the scope of the code gates. | Gate 1, after OQ-2. |
| OQ-4 | What is the combat model? Turn order, party size, the stats, and the actions. | It sets the entity model and the rules engine. | Gate 1, after OQ-3. |
| OQ-5 | What does the screen show? A top-down tile map, panes of menus and text, or both? | It sets the render model and the input model. | Gate 1. |
| OQ-6 | Which glyphs can the game show? ASCII alone, or also Unicode box-drawing characters and symbols? | No width library is permitted (D-21). A wide rune, for example an emoji, uses two columns and breaks the column math. | Gate 1. |
| OQ-7 | Which colors? None, the 16 ANSI colors, 256 colors, or 24-bit color? | Each terminal of macOS supports a different set. The answer sets the escape sequences. | Gate 1. |
| OQ-8 | Can the player save and load a game? If so, which format: `encoding/json` or `encoding/gob`? | The save format becomes a contract. A later change needs a migration. | Gate 1. |
| OQ-9 | Where does the game content live: in Go code, or in data files that `embed` loads? Which paths? | The content is game data, and the game data belongs to the owner (D-11). The guard must cover the data paths (D-19). | Gate 1. |
| OQ-10 | Does the game use a seeded random source from `math/rand/v2`, so that a run can repeat? | A seed makes a test repeatable and makes a bug report repeatable. | Gate 1. |
| OQ-11 | Which entity model? For example: structs with composition and small interfaces, or an entity-component-system. | The entity structure is the first part of "perfect bones" (D-1). | Gate 1. |
| OQ-12 | What is the error policy? Wrap with `%w`, sentinel errors, custom error types, `panic` for programmer errors alone, and the restore of the terminal after a panic. | The error handling is a part of "perfect bones" (D-1). | Gate 1. |
| OQ-13 | Which numbers apply to the code? For example: the max lines of a function, the max parameters, the max depth of nested blocks, and the max size of a file. What are the rules of the package layout? | A review needs a rule that it can measure. Gate 1 adds the numbers to the checklist of `AGENTS.md`. | Gate 1. |
| OQ-14 | What is the test policy? Which parts must have tests, table-driven tests or not, and a coverage target or not? | It sets a success criterion of each code gate. | Gate 1. |
| OQ-15 | Do we use a dev tool outside the Go toolchain, for example `staticcheck`? Since Go 1.24, a `tool` directive adds `require` lines to `go.mod`. | It can conflict with D-21. A binary that stays out of `go.mod` does not. | Gate 1. |
| OQ-16 | Does a GitHub Actions workflow run `go vet`, `go test`, and `gofmt -l` on each pull request? | The game is for macOS alone (D-20). The `syscall` constants of macOS do not compile on Linux, so the runner must be macOS. The agent can write the workflow (D-11). | Gate 1 decides. A code gate adds it. |
| OQ-17 | Which Go version goes in the `go` line of `go.mod`? The toolchain on this Mac is 1.27.2. | It sets the language features that the code can use. | The first code gate. |
| OQ-18 | How does the game write logs? The game owns the terminal, so a log on stdout breaks the screen. For example: `log/slog` to a file, behind a debug flag. | Without a log, a bug in the game loop is hard to find. | Gate 1. |
| OQ-19 | What is the minimum size of the terminal, and what does the game do on a resize? | It sets the layout rules and the use of `SIGWINCH`. | Gate 1. |
| OQ-20 | Is a `Makefile` or a shell script tool config, which the agent can edit, or code of the owner? | D-11 does not name these files, so they belong to the owner now. | The first gate that adds such a file. |
| OQ-21 | Which license does the public repo use? With no license, the default copyright applies, and other people have no right to reuse the code. | The repo is public (D-29). | Before the first code gate merges. |
| OQ-22 | Who writes the `README.md`, and when? | It is the first page of the public repo. | Gate 1. |
| OQ-23 | What is the name of the game? | The name goes in the `README.md` and on the title screen. The module path can also use it. | Gate 1. |

## Asked, no answer yet

No question waits now.

## Answered

| # | Question | Answer |
|---|---|---|
| OQ-1 | What external libraries, if any, must we have? | D-21: the standard library alone. |
