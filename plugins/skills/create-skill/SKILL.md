---
description: Create a new thragg skill following the standard format
---

# Create Skill

## Overview

Guides the creation of a new skill in the thragg plugin. Ensures the skill follows the standard structure, is placed in the right location, and is well-scoped.

## Tools

- `Write` — creating the new SKILL.md and any reference files
- `Read` — checking existing skills for consistency

## When to use

When you want to add a new slash command to thragg — for example a code review skill, a PR description skill, or a debugging guide.

## Examples

```
/thragg:create-skill
```

- "Create a skill for reviewing pull requests"
- "I want a skill that helps write commit messages"

## Steps

1. Ask the user what the skill should do and when it should be invoked, if not already clear.
2. Confirm the skill name — it must be kebab-cased (e.g. `review-pr`, `write-commit`).
3. Create the folder at `plugins/skills/<skill-name>/`.
4. Write `plugins/skills/<skill-name>/SKILL.md` using the structure below.
5. If the skill benefits from reference code or examples, create a `plugins/skills/<skill-name>/references/` folder and add the relevant files there. Reference them in the skill.
6. Ask the user if the new skill should be committed.

## SKILL.md structure

```markdown
---
description: one-line description shown in the plugin manager
---

# Title

## Overview
What this skill does in 2-3 sentences.

## Tools
What tools this skill uses (WebSearch, WebFetch, Bash, Read, Write, Edit, gh, etc.)

## When to use
Specific situations or triggers that make this skill relevant.

## Examples
Concrete examples of how to invoke or use the skill.

## Principles (optional)
Key principles that guide decisions made during this skill.

## Steps
Numbered steps Claude follows when running this skill.

## Rules (optional)
Links to rule files this skill reads and applies:
- `${CLAUDE_PLUGIN_ROOT}/rules/example.md` — short description

## References (optional)
Links to reference files in the references/ folder:
- `references/example.ts` — short description

## Common mistakes
What to avoid when using or applying this skill.
```

## Common mistakes

- Scoping a skill too broadly — one skill should do one thing
- Skipping the `description` frontmatter — it shows in the plugin manager
- Using PascalCase or snake_case for the folder name instead of kebab-case
- Adding a references folder when the skill is simple enough not to need one
