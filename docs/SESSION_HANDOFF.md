# Session hand-off

`AGENTS.md` is the entry point, and it sends you here. This file holds the resume point and the procedures of each session. It also holds the next steps, the facts that expire, and the three most recent session records (D-27).

Update this file inside the pull request of the gate, before the owner merges it. Never write a merge commit, a pull request number, or a check result that you did not read.

## RESUME HERE (2026-10-18, gate 2 active, fourth pause)

The owner resumes with the words `Resume gate 2` (D-88). Do the start check of section 1. Take the state from this section.

**Where the gate stands.** Gate 2 is active: the terminal. Branch: `gate-02-terminal`. Head at the pause: `5187f86` ("mid gate 2"). Gate file: `docs/gates/gate-02-terminal.md`. At the pause, the tests in `internal/terminal/error_test.go`, the decision D-89, and the doc changes of D-89 are not committed. The owner commits and pushes the branch before a move to another machine (D-12). After the pull, the head can differ from `5187f86`. Tell the owner, and go on.

**The finished work.**

- Decisions D-74 to D-89 are recorded. The ones for the code: the package is `terminal` (D-79), the fields of `Error` start with a small letter (D-84), the cause is never nil (D-85), and each `Error` method has a pointer receiver (D-86). D-89: an exported name needs a doc comment only when the name does not state its whole job.
- `go.mod` holds the module `github.com/nkramber/below-the-keep` and the line `go 1.27.2`.
- `internal/terminal/doc.go` holds the package comment.
- `internal/terminal/error.go` is complete and reviewed. It holds the constant `packageName`, the struct `Error` with the fields `operation` and `cause`, and the methods `Error()` and `Unwrap()`.
- `internal/terminal/error_test.go` is complete and reviewed. `TestError` checks the exact text `terminal: testOperation: test error for cause`. `TestUnwrap` checks that `errors.Is` finds the cause. The owner did a mutation check for each test, and both failed as they must (read 2026-10-18).
- The owner ran the five checks of hard rule 5 in the repo root (read 2026-10-18). `go build`, `go vet`, `gofmt -l`, and `staticcheck` printed nothing. `go test` printed `ok` for `internal/terminal`. The root has no `main.go` yet, so the build covers only `terminal`.
- `staticcheck` is in `~/go/bin`, and the owner put that folder in the `PATH`.
- `LICENSE` and `.gitignore` are written. The folder `go-practice/` is in the repo (D-87).
- The tutorials on functions and methods, and on writing tests, are done. The tutorial on tests ended with Part 3 of the tests for `Error`.

**The open work.** The owner writes the code of the gate, in this order:

1. **The raw mode in `internal/terminal`. This is the next work.** See "Next action".
2. The other parts of `internal/terminal`: the alternate screen, the cursor, the window size, and the resize events.
3. The skeleton of `main.go`, with the demo.
4. The five checks and the manual checks of the gate file.

**The tutorials in progress.** The skill `tutorial` (D-83) governs them. The owner says "Review" to ask for a review. The agent sends one part in each message. No tutorial is in progress now. The tutorial on the raw mode has not started.

**Open findings:** none. The findings of the tutorial reviews are fixed or kept. E-5 is kept (package-level test names with the prefix `test`). **Open questions of this gate:** none. OQ-25 waits for gate 3.

**Notes for the next session:**

- Before the first function that returns `&Error{...}`, tell the owner to read the question "Why is my nil error value not equal to nil?" at https://go.dev/doc/faq#nil_error (D-86).
- With a pointer error type, the target of `errors.As` is the address of a variable of type `*terminal.Error`. Explain it when the owner reaches that code.
- The owner asks many questions about Go words. Use the skill `explain-go`. The comparison of the parcel for wrapped errors, and the photocopy and the address slip for receivers, work well.
- The owner runs `go test -count=1 ./...` in `go-practice/`. The root pattern `./...` skips that folder.
- **How the owner wants the steps.** The owner said that vague steps are hard to follow. Name each piece: the file, the function, the Go feature, and the `go doc` topic. Do not write "decide" without the options. Do not write the code (hard rule 1), but name every part of it.
- **Doc comments (D-89).** Do not ask for a comment on a name that says its whole job. The owner found such requests noise.
- **Unsaved files.** A mutation check once passed because the file was not saved. Tell the owner to save the file before each run.
- **Package-level names in tests.** Go has no scope for one file. A name at the top of a `.go` file belongs to the whole package, tests too. The owner keeps the prefix `test` for such names.
- **A failure message.** The owner likes a message of the form `call(inputs) = result, want result`. The standard terms (`go doc errors`, read 2026-10-18): an error wraps another error, the wrapped error is a child, and `errors.Is` looks in the tree.

**Next action.** Do the start check of section 1. Tell the owner the state in max five lines. Then plan the raw mode with the owner, and wait. Do not send a tutorial part before the owner asks.

1. Read the gate file, P-29 and P-31 of `docs/code-standards.md`, and "Facts that expire" for the `syscall` names.
2. Ask the owner the open design questions with `AskUserQuestion`. The agent must not decide them. The questions: the form of the API (a function that returns a restore function, or a type with methods), the file layout (the calls that need `unsafe` go in one small file for macOS, P-29), which file descriptor to use, which flags the raw mode changes, and how an error of `syscall` becomes a `terminal.Error`. Record each answer in `docs/decisions.md` at once.
3. Check each fact about `syscall` with `go doc` before it goes into a part (hard rule 10).
4. Offer the tutorial on the raw mode, one part in each message (D-83). The first part holds the plan.
5. The owner then writes the code and the tests, and asks for a review.

## The procedures of a session

### 1. The start check

Do these steps before the first edit:

1. Read `AGENTS.md`, this file, and the gate file that the prompt names.
2. Read the decisions that the work touches, by `D-` id.
3. Read this conversation. Look for work on another gate.
4. When you find such work, stop. Answer only `Blocked: start a new clean session for this gate.`
5. A fork, a subagent, or a summary of another session is not a clean session. Stop for these too.
6. Run `git status -sb` and `git log --oneline -3`. Record the branch and the head.
7. When the branch is `main`, ask the owner to make the gate branch. Edit no file on `main`.
8. When the branch does not agree with the prompt, tell the owner, and stop. For the words `Resume gate <NN>` (see "The short resume" in section 3), the branch comes from "RESUME HERE". A different head is not a stop. Tell the owner.
9. When the gate file shows "planned", write "active" in the gate file and in the roadmap row (D-35). This is the first edit of the gate.
10. Read the open findings and the open questions of the gate file.
11. Tell the owner the state in max five lines. Then do the first action of the prompt.

### 2. The work of a session

A session can plan, answer, and review. Each of these jobs stays inside the one gate of the session.

- **Plan:** write and change the gate file with the owner. The owner approves each success criterion.
- **Answer:** give the generic form of a language feature (hard rule 2 of `AGENTS.md`).
- **Review:** follow "The review procedure" of `AGENTS.md`.

Obey these rules during the work:

- Ask each question when you think of it. Record each answer in `docs/decisions.md` at once.
- Move each answered question to "Answered" in `docs/open-questions.md`.
- Check a success criterion in the gate file only after the owner shows the evidence.
- When the owner asks for work of another gate, say so. Offer to record it as an open question, or as a finding that moves to that gate. Do not do the work.

### 3. A pause

A pause ends a session, and the gate continues in a new session. One of these events starts a pause:

- The owner says to stop, to pause, or to end the session.
- A context compaction occurred, and the owner agrees to a pause.
- The agent proposes a pause after a long session, and the owner agrees.

Do these steps:

1. Finish the current step. When you can not finish it, record it as open work.
2. Write the state in "RESUME HERE": the finished work, the open work, the open findings, and the next action.
3. Add or update the session record of today.
4. Tell the owner which doc files changed. The owner commits them (D-12).
5. Write the continue prompt below in the last message, and stop.

```
Continue gate <NN>: <the name of the gate>

Read `docs/SESSION_HANDOFF.md` first. Then read `docs/gates/gate-<NN>-<slug>.md`.
Branch: `gate-<NN>-<slug>`. Head: `<sha>`.
Open findings: <each F- id with its title, or `none`>.
Open questions of this gate: <each OQ- id with its subject, or `none`>.
First action: <the next concrete action>.
```

**The short resume (D-88).** The owner can write only `Resume gate <NN>`. This works only when "RESUME HERE" holds the exact state. So the pause (steps 1 to 3 above) must record the finished work, the open work, the tutorials in progress, the open findings, the open questions, and the next action. After a pause, "RESUME HERE" must let a new session start with no other input. Write no path of the repo on one machine in any file. Use paths from the root of the repo.

### 4. The gate end

The completion check passes when each item below is true:

- Each success criterion of the gate file holds, or an owner decision moved it.
- Each finding has the status fixed, kept, or moved.
- A code gate: the owner ran the five checks of hard rule 5, and the output shows no failure.
- A code gate: each import comes from the standard library (D-21).
- `docs/decisions.md` and `docs/open-questions.md` agree with the work.
- `docs/roadmap.md` and the gate file show the status "merged" (D-32).
- "RESUME HERE" and "Next steps, in order" name the next gate and its first action.

Do these steps:

1. Write "merged" as the status of the gate in `docs/roadmap.md` and in the gate file (D-32).
2. Write "RESUME HERE" and "Next steps, in order" for the next gate.
3. Do the completion check. Tell the owner the result of each item.
4. When an item fails, fix the docs part or name the owner part. Then do the check again.
5. The owner commits, pushes, and opens the pull request. The owner writes each commit message, the title, and the body (D-33).
6. The owner reads the diff on GitHub, and merges with squash (D-16).

The status text reaches `main` only through the merge, so it is true when it arrives (D-32). When the owner asks for a change after the pull request opens, edit the docs again. The owner then pushes the change.

### 5. The transition

The transition starts when the owner says that the pull request of this gate merged. Each of these messages is a trigger, and each similar message too:

- `merged`
- `merged #N`
- `PR #N is merged`
- `#N merged`

Two cases stop the transition. Ask the owner, and write no prompt before the answer:

- The message names no pull request, and this session has no gate. Ask which pull request it names.
- The owner merged before the completion check passed. Name each part that did not land. Then ask the owner for the next step.

Do these steps:

1. Run `gh pr view <N> --json state,mergedAt,mergeCommit`.
2. Make sure that the state is `MERGED`, and that the pull request is the one of this gate.
3. Read "Next steps, in order" below, and name the next gate.
4. Read `docs/open-questions.md`, and name each question of the next gate.
5. Name each check that needs `main` after this merge.
6. Write the start prompt below in the last message, and stop.

```
Start gate <NN>: <the name of the gate>

Gate <MM> merged to `main` as <sha> (pull request #<N>).
Owner: make the branch `gate-<NN>-<slug>` from `main` before the first edit.
Read `docs/SESSION_HANDOFF.md` first. Then read `docs/gates/gate-<NN>-<slug>.md`.
<Each check that needs `main` after this merge. Do it before the work of the gate.>
Open questions of this gate: <each OQ- id with its subject, or `none`>.
First action: <the first concrete action>.
```

The rules of the start prompt:

- The prompt carries one gate. A second gate needs a second session and a second prompt.
- The prompt never asks the next session to record this merge. Git holds the merge.
- Remove the line of checks when this merge needs no check.
- "Next steps, in order" holds the same first action. When the two do not agree, this file wins.

After the prompt, write this line:

`This session is bound to gate <MM> and is complete. Start a new clean session for gate <NN>.`

The session makes no edit for the next gate.

### 6. After a context compaction

A context compaction keeps the session bound to its gate.

1. Read "RESUME HERE" and the gate file again.
2. Tell the owner that a context compaction occurred.
3. Offer a pause (section 3).

### Traps for a new session

- A change of the guard shows in `/permissions` in the same session. On 2026-10-16, the owner saw 78 rules right after the agent added two rules.
- The guard denies a whole compound command when one part of it matches a rule. For example, a chain with `git branch --show-current` fails.
- Never make a `.go` file, also not as a probe of the guard.
- The guard blocks `git branch`. Use `git status -sb` to read the branch.
- The guard blocks `git -C`. Start a command with an absolute `cd` to the root of the repo. Read the root with `pwd` at the start of a session. The root differs on each machine, so write it in no file.
- The guard blocks `go env`, because hard rule 5 allows only `go doc` and `go version`. For the module cache, use the default path `~/go/pkg/mod`.
- The view of the owner shows the text before a tool call only as a short summary. Put each full answer in the last message of a turn, after the last tool call. Put each fact that a question needs in the `AskUserQuestion` box: in the question, an option, or a preview.
- A merge before the status mark leaves `main` stale. Pull request #1 did this, and the follow-up of D-31 fixed it. Steps 1 and 2 of section 4 come before the pull request opens.

## Next steps, in order

1. **Gate 2: the terminal** (D-73). Gate file: `docs/gates/gate-02-terminal.md`. First action: the start check, then plan the raw mode with the owner (see "RESUME HERE").
2. **Gates 3 to 14,** in the order of `docs/roadmap.md`. The first session of each gate reviews the success criteria of its gate file with the owner.

## Facts that expire

- The Go toolchain on this Mac: `go1.27.2 darwin/arm64` in `/usr/local/go` (read 2026-10-16).
- Go 1.27.2 `syscall` on darwin/arm64 holds `TIOCGETA`, `TIOCSETA`, `TIOCGWINSZ`, `SIGWINCH`, and the `Termios` type. It holds no type for the window size (read 2026-10-16 in the Go source).
- `golang.org/x/term` v0.47.0 of 2026-10-09 needs `golang.org/x/sys` v0.49.0 (read 2026-10-16). The repo does not use it (D-21).
- Claude Code 2.1.288 is on this Mac (read 2026-10-16). It reads `AGENTS.md` at the start of a session when no `CLAUDE.md` exists (D-28).
- The Edit deny rules cover the Edit, Write, and NotebookEdit tools, a Bash redirect, and the known file commands. They do not cover a script that writes a file (Claude Code docs, read 2026-10-16).
- The GitHub repo `nkramber/below-the-keep` is public (read 2026-10-16 with `gh repo view`). Its old name was `terminal-rpg` (D-42). Per the GitHub docs, the old URL redirects to the new one, until a new repo takes the old name. On 2026-10-16, `main` held `9969b29` (the empty root commit), `e275872` (#1), and `b2d2b15` (#2).
- The `origin` remote of the local clone is `https://github.com/nkramber/below-the-keep.git` (read 2026-10-16 in `.git/config`). The local folder has the name `below-the-keep` (D-42). Its place on disk differs on each machine (D-88).
- The guard holds 78 deny rules: 6 Edit, 30 git, 25 gh, 16 go, and 1 staticcheck (`/permissions`, read by the owner on 2026-10-16 and on 2026-10-18).
- On Go 1.27.2, `go mod init` writes the version of the local toolchain in the `go` line (`cmd/go/internal/modload/init.go`, read 2026-10-18).
- A directory that holds a `go.mod` is skipped by `./...` of the parent module (`go help packages`, read 2026-10-18). So `go-practice/` is outside the five checks of the root, except `gofmt -l .` (D-87).
- The Go FAQ (read 2026-10-18) advises a value receiver for a small struct that a method only reads. The Code Review Comments say "when in doubt, use a pointer". The owner chose pointers (D-86).
- `staticcheck` is in `~/go/bin` on this Mac, and the owner put that folder in the `PATH` (read 2026-10-18). The five checks of the root passed on 2026-10-18.
- Python 3 on this Mac holds the Unicode 13.0 data in `unicodedata`. It gives the width class of a glyph (read 2026-10-16).

## The three most recent sessions

### 2026-10-18d: gate 2, the tests and the checks

**The owner finished the tutorial on tests** in `go-practice/`. Part 5 added a named error, `%w`, and a test with `errors.Is`. The mutation check changed `%w` to `%v`, and the test failed as it must.

**The owner wrote the two tests of `terminal.Error`.** Reviews found five small findings, and all are fixed or kept. The mutation check on `Unwrap()` failed the test, and the restore passed it. The five checks of the root passed.

**The owner set D-89.** A doc comment is needed only when the name does not state its whole job. The agent changed `AGENTS.md`, `docs/code-standards.md`, and the gate files.

**The agent made one mistake.** An unsaved file made a mutation check pass. The owner saw the result and asked, and the agent found the cause.

**The session paused** with this record. The next work is the raw mode.

### 2026-10-18c: gate 2, the tutorial on tests, Parts 3 and 4

**The owner finished Parts 3 and 4** in `go-practice/`. The test is now a table, with separate checks for the answer, the missing error, and the text of the error. The owner proved the checks with a mutation: a wrong letter in the message made the test fail.

**Six findings came from the reviews,** and all are fixed. The main ones: a long `||` condition that passed on wrong code, and a format string with the wrong number of verbs.

**The agent made one mistake.** It told the owner that `err` was not `nil` inside a check that tests `err == nil`. The owner caught it. The agent then gave the correct table of four cases.

**Part 5 is sent and not started.** The agent also gave the explanation and the example of a struct, on request.

**The session paused** with this record. `go-practice/main_test.go` was not committed at the pause.

### 2026-10-18b: gate 2, the error type and the tutorials

**The owner wrote `internal/terminal/error.go`** with the help of a tutorial and four reviews. The package is `terminal` (D-79). The fields are private (D-84), the cause is never nil (D-85), and the methods have pointer receivers (D-86). The review findings are all fixed.

**The agent made two skills** (D-83). The skill `tutorial` gives one part in each message, and it reviews the work of a part when the owner says "Review". The skill `explain-go` (D-82) answers questions about Go words for a reader who never programmed.

**Two tutorials ran.** Functions and methods are done. Writing tests is at Part 3, in `go-practice/` (D-87). The owner chose pointer receivers for the practice of pointers, and the agent noted the FAQ advice for a small read-only struct.

**The owner set two process rules** (D-87, D-88). The practice folder lives in the repo. The words `Resume gate <NN>` start a session, and no file holds a path of one machine.

## The archive

`docs/session-archive.md` holds the older session records, word for word (D-27). The records of 2026-10-16 and 2026-10-16b, both gate 0, moved there on 2026-10-18. The record of 2026-10-16c, gate 1, moved there on 2026-10-18 too. The record of 2026-10-18, gate 2, moved there on 2026-10-18 too.
