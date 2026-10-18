---
name: explain-go
description: Use when the owner asks what a Go word, a syntax form, or a definition means. Answer so that a person who has never programmed can understand.
---

# Explain Go to a beginner (D-82)

The owner asks a question about Go syntax or a Go definition. Write the answer for a reader who has never programmed.

## Steps

1. Say what the thing is in plain words. Use a daily-life comparison.
2. Define each programming word that the answer uses, the first time that the answer uses it. Do not say "as you know".
3. Say why the thing exists. Name the problem that it solves.
4. Show one short generic example when the owner asks for one (hard rule 2 of `AGENTS.md`). Use placeholder names that are not in this repo, for example `Widget`.
5. Explain each line of the example in words. Do not leave a symbol with no meaning.
6. Tell the owner how the thing links to the code in this repo, but do not give a fix.

## Rules

- The answer is long enough to be complete. It is not short for the sake of short.
- Hard rules 2, 3, and 6 of `AGENTS.md` still apply. Write in ASD-STE100 (D-7). Do not paste a solution.
- Verify each fact of the language with `go doc` or the Go spec (hard rule 10).
- When the question has a wrong premise, say so and give the evidence.
