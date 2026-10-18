@AGENTS.md

# Skills (D-82)

- `explain-go` (`.claude/skills/explain-go/SKILL.md`): when the owner asks a question about Go syntax or a Go definition, answer so that a person who has never programmed can understand. Invoke the skill for each such question.
- `tutorial` (`.claude/skills/tutorial/SKILL.md`): when the owner asks for a tutorial on how to do something, send a short plan and Part 1 in the first message. Then send one part in each message. When the owner says "Review", review the work of the previous part first. Send the next part only when the review finds no mistake. A change in one part then never makes another part out of date (D-83). Invoke the skill for each such request.
