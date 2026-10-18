---
name: tutorial
description: Use when the owner asks for a tutorial on how to do something. Give a short plan and Part 1 in one message. Then give one part at a time, and check the work of each part before the next one.
---

# Give a tutorial one part at a time (D-83)

The owner asks for a tutorial. The owner asks questions about a part, and a question can change the work. A part that the agent gave earlier can then be out of date. So the agent gives one part in each message. The agent also checks the work of each part before it goes on.

## Steps

1. Send one first message. It holds the audience line, the title, the goal, a high-level plan, and Part 1. The plan is one line for each part, with no detail. Do not wait for the owner to say "start".
2. A part holds its goal, its steps, and a check that the owner can do. It ends with a line that tells the owner to ask a question, or to say "Review" when the work is done.
3. Wait. Send the next part only after the owner says "Review" or a similar word, and the review finds no mistake.
4. When the owner says "Review", review the work of the previous part first. Read the files that the part named. Check them against the check of the part, the rules of `AGENTS.md`, and `docs/code-standards.md`.
5. When the review finds a mistake, do not send the next part. Tell the owner the location (`file:line`), the problem, and the rule. Then ask one question that points to the problem. Give no fix.
6. Wait for the owner to correct the mistake. When the owner asks for more help, go one step up the hint ladder (D-24). Review the corrected work when the owner says "Review" again.
7. When the review finds no mistake, say so in one line. Then send the next part.
8. Write each part when the owner asks for it, from the current state of the work. Do not write it earlier.
9. When the owner asks a question about a part, answer it and stay in that part. When the answer changes a decision, record it as a decision (hard rule 8 of `AGENTS.md`).
10. When a change touches a later part, do not rewrite that part now. Say in one line which later parts the change touches. The next part then comes with the change already inside it.

## Rules

- One part in each message. The first message is the one exception: it holds the plan and Part 1.
- A part does not repeat the text of an earlier part. When a part needs a fact from an earlier part, it states the fact again in one line.
- Each part can stand alone, so the owner can read only the part that the work needs.
- Use the current names. When a rename happened, the new part uses the new name.
- The tutorial teaches the owner to do the work. The agent writes no production code (hard rule 1 of `AGENTS.md`).
- The review of a part follows "The review procedure" of `AGENTS.md`, but it covers only the work of that part. It does not write the finding table of the gate file. A finding that the owner keeps open at the end of the tutorial goes to the gate file.
- A code example needs a request of the owner (hard rule 2). Give the language feature, with placeholder names that are not in this repo.
- Hard rule 6 of `AGENTS.md` still applies. Write in ASD-STE100 (D-7).
- Verify each fact of the language with `go doc` or the Go spec (hard rule 10).
- For a question about the meaning of a Go word, use the skill `explain-go`.
- Name the audience in the first line of the first message, for example "Written for: the owner, who has not done this before."
