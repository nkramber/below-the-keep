# Open questions

This file holds each question with no answer yet. When the owner answers a question, write the answer as a row in `docs/decisions.md`. Then move the question to "Answered" below, with the `D-` id.

Ask a question when the work reaches the point that needs the answer. Ask in batches of max four questions (hard rule 8 of `AGENTS.md`).

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-17 | Which Go version goes in the `go` line of `go.mod`? The toolchain on this Mac is 1.27.2. | It sets the language features that the code can use. | Gate 2 (D-73). |
| OQ-20 | Is a `Makefile` or a shell script tool config, which the agent can edit, or code of the owner? | D-11 does not name these files, so they belong to the owner now. | The first gate that adds such a file. |
| OQ-21 | Which license does the public repo use? With no license, the default copyright applies, and other people have no right to reuse the code. | The repo is public (D-29). | Gate 2, before it merges (D-73). |
| OQ-24 | Which exit code does the game use for an error, and which for a panic? D-55 says only "a code that is not 0". | A manual check of gate 2 reads the code with `echo $?`. | Gate 2. |

## Asked, no answer yet

No question waits now.

## Answered

| # | Question | Answer |
|---|---|---|
| OQ-1 | What external libraries, if any, must we have? | D-21: the standard library alone. |
| OQ-2 | What is the game style? What tone, and what world? | D-36: a classic roguelike in a classic fantasy world, with a light tone and dry humor. |
| OQ-3 | What is the core loop of the gameplay? How does a run end? | D-37: explore, fight, use items, and go down. 5 floors, a win at the artifact, and nothing carries over. D-40: levels, experience points, and items make the hero stronger. |
| OQ-4 | What is the combat model? Turn order, party size, the stats, and the actions. | D-38: one action each, a hit roll and a damage roll, melee and items, and monsters that wander and then chase. D-39: five stats, regen and potions, no hunger, and items on the floor and from drops. D-41: strength replaces damage. |
| OQ-5 | What does the screen show? | D-45: no limit of 80 by 24. D-46: the map on the left, a side panel on the right, and the messages at the bottom. |
| OQ-6 | Which glyphs can the game show? | D-43: ASCII, box-drawing characters, and Unicode symbols. D-48: a fixed list of glyphs that are one column wide. |
| OQ-7 | Which colors? | D-44: the 256 colors. D-49: the game paints the background. |
| OQ-8 | Can the player save and load a game? Which format? | D-50: no save. A run ends in one sitting. |
| OQ-9 | Where does the game content live? Which paths? | D-51: JSON files that `embed` loads. OQ-13 sets the path. |
| OQ-10 | Does the game use a seeded random source? | D-52 and D-53: one seed makes a map source and a play source. A flag sets the seed, and the end screen shows it. |
| OQ-11 | Which entity model? | D-54: plain concrete types. `Actor` and `Item` hold a kind from JSON and their own state. |
| OQ-12 | What is the error policy? | D-55: an error type for each package, rule outcomes as results, and `main` handles each error and each panic. |
| OQ-13 | Which numbers apply to the code? What are the rules of the package layout? | D-57: 40 lines, 4 parameters, 3 levels, and 400 lines in a file. `main.go` at the root, and the packages in `internal/`. D-58: two details. |
| OQ-14 | What is the test policy? | D-59: tests for the logic, table-driven tests for many cases, a manual checklist for the terminal, and no coverage target. |
| OQ-15 | Do we use a dev tool outside the Go toolchain? | D-66: staticcheck, installed with `go install ...@latest`, outside `go.mod`. It is the fifth check. |
| OQ-16 | Does a GitHub Actions workflow run the checks on each pull request? | D-64: no CI. The owner runs the checks on the Mac before each commit. |
| OQ-18 | How does the game write logs? | D-60, D-62, and D-63: always on, one file for each run in `~/Library/Caches/below-the-keep/`, max 10 MB in total, text at the info level, and a `-debug` flag. |
| OQ-19 | What is the minimum size of the terminal, and what does the game do on a resize? | D-47: the whole map is visible, with a fixed size. A window that is too small pauses the game. |
| OQ-22 | Who writes the `README.md`, and when? | D-65: the agent, in gate 1. `README.md` joins the docs list of hard rule 1. |
| OQ-23 | What is the name of the game? | D-42: "Below the Keep". The repo, the folder, and the module path use `below-the-keep`. |
