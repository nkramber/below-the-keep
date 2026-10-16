# Session hand-off

`AGENTS.md` is the entry point, and it sends you here. This file holds the resume point and the procedures of each session. It also holds the next steps, the facts that expire, and the three most recent session records (D-27).

Update this file inside the pull request of the gate, before the owner merges it. Never write a merge commit, a pull request number, or a check result that you did not read.

## RESUME HERE (2026-10-16, gate 1 next)

**Gate 0 merged.** Pull request #1 and the follow-up of D-31 carry it. No code exists.

**Gate 1 comes next: the game design and the roadmap** (D-17). Gate file: `docs/gates/gate-01-design.md`.

- Branch: `gate-01-design`. The owner makes it from `main` (D-12).
- Open questions: OQ-2 to OQ-16, OQ-18, OQ-19, OQ-22, and OQ-23.

**The check before the work.** The owner runs `/permissions` and looks for the 76 deny rules of the guard (D-19).

**Next action.** Do the start check of section 1. Then ask OQ-2, the game style.

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
8. When the branch does not agree with the prompt, tell the owner, and stop.
9. Read the open findings and the open questions of the gate file.
10. Tell the owner the state in max five lines. Then do the first action of the prompt.

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

### 4. The gate end

The completion check passes when each item below is true:

- Each success criterion of the gate file holds, or an owner decision moved it.
- Each finding has the status fixed, kept, or moved.
- A code gate: the owner ran the four checks of hard rule 5, and the output shows no failure.
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

- The guard loads at the start of a session. The session that made `.claude/settings.json` possibly ran without its rules (unverified).
- Never make a `.go` file, also not as a probe of the guard.
- The guard blocks `git branch`. Use `git status -sb` to read the branch.
- The guard blocks `git -C`. Start a command with an absolute `cd` to the repo.
- A merge before the status mark leaves `main` stale. Pull request #1 did this, and the follow-up of D-31 fixed it. Steps 1 and 2 of section 4 come before the pull request opens.

## Next steps, in order

1. **Gate 1: the game design and the roadmap** (D-17). Gate file: `docs/gates/gate-01-design.md`. First action: ask OQ-2, the game style.

## Facts that expire

- The Go toolchain on this Mac: `go1.27.2 darwin/arm64` in `/usr/local/go` (read 2026-10-16).
- Go 1.27.2 `syscall` on darwin/arm64 holds `TIOCGETA`, `TIOCSETA`, `TIOCGWINSZ`, `SIGWINCH`, and the `Termios` type. It holds no type for the window size (read 2026-10-16 in the Go source).
- `golang.org/x/term` v0.47.0 of 2026-10-09 needs `golang.org/x/sys` v0.49.0 (read 2026-10-16). The repo does not use it (D-21).
- Claude Code 2.1.288 is on this Mac (read 2026-10-16). It reads `AGENTS.md` at the start of a session when no `CLAUDE.md` exists (D-28).
- The Edit deny rules cover the Edit, Write, and NotebookEdit tools, a Bash redirect, and the known file commands. They do not cover a script that writes a file (Claude Code docs, read 2026-10-16).
- The GitHub repo `nkramber/terminal-rpg` is public. On 2026-10-16, `main` held `9969b29`, the empty root commit, and `e275872`, the squash commit of #1.

## The three most recent sessions

### 2026-10-16b: gate 0, the follow-up of the status

**Pull request #1 merged as `e275872`** before the status mark. Three parts stayed stale on `main`: the roadmap row of gate 0, the gate file, and "RESUME HERE".

**The owner chose a follow-up pull request** (D-31) and a new rule for the gate end (D-32). The owner also approved D-22 to D-29 (D-30).

**This pull request obeys D-32.** It marks gate 0 as merged before it opens, and "RESUME HERE" names gate 1.

**The owner writes each commit message and each pull request text** (D-33). The pull request template is gone (D-34).

### 2026-10-16: gate 0, the process and the session rules

**The owner set the frame** (D-1 to D-10). The owner gave the goals in order, the language, and the number one rule. The owner also set ASD-STE100 in chat, one gate in each pull request, and no Gitar and no second-provider review.

**The agent read decktome.** It read `AGENTS.md`, `CLAUDE.md`, `docs/SESSION-HANDOFF.md`, the `one-pr-one-session` skill, the two question registers, `docs/decisions.md`, the pull request template, and the `ste-writing` skill.

**This repo keeps these parts of decktome:** the start check, the clean session, the resume section, the facts that expire, the three session records, the decision register, and the transitional prompt. This repo calls that prompt the start prompt.

**This repo changes these parts:** the owner writes the code and does each git step. The agent reviews in place of Gitar and Codex. A gate can use many sessions, and a pause ends with a continue prompt. No skill, hook, or `make` target enforces the rules. The guard and the agent do.

**The owner answered 13 questions** (D-8, D-11 to D-21). The agent proposed eight rules for sign-off at the merge (D-22 to D-29).

## The archive

`docs/session-archive.md` does not exist yet. The fourth session record makes it (D-27).
