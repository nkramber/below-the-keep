# Session archive

This file holds the session records that are older than the three records in `docs/SESSION_HANDOFF.md` (D-27). Each record moves here word for word. The newest record is at the top.

### 2026-10-16c: gate 1, the game design and the roadmap

**The owner chose the game** (D-36 to D-42). It is a classic roguelike with the name "Below the Keep", in a classic fantasy world with a light tone. A run has 5 floors, items, levels, and a win at the artifact.

**The owner set the screen, the content, and the map** (D-43 to D-53, D-67 to D-69). The game shows the whole floor, with 256 colors and a fixed glyph list. The floors are caves from a seeded map source, and the content is in JSON files.

**The owner set the bones** (D-54 to D-66, D-70). The entity model uses plain types, and each package has an error type. `main` handles each error and each panic. The limits are 40, 4, 3, and 400. Staticcheck is the fifth check, and the project has no CI.

**The owner accepted four docs:** `docs/game-design.md` (D-71), `docs/code-standards.md` (D-72), the roadmap with 13 code gates (D-73), and `README.md` (D-65).

**The process changed** (D-35). The start check now sets a gate to "active". The owner renamed the repo to `below-the-keep` (D-42). The view of the owner shows the text before a tool call only as a summary, so the hand-off now has a trap for it.

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
