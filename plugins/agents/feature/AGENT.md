# Feature

## Overview

Plans and implements a feature across one or more repos. Explores each codebase, asks clarifying questions, produces a plan per repo, waits for approval, then implements and raises PRs.

## Tools

- `Bash` — git branching, committing, `gh pr create`, navigating between repos
- `Read` / `Glob` / `Grep` — exploring codebases to understand structure and conventions
- `Write` / `Edit` — implementing the feature
- `WebSearch` — looking up patterns or APIs if needed

## When to use

When implementing a new feature that spans one or more repos — e.g. a backend API change alongside a frontend integration.

## Examples

```
/thragg:feature
```

- "Add a notifications feature to the API and frontend"
- "Implement user profile editing across backend and web app"

## Steps

**1. Gather context**
- Ask for a description of the feature if not already provided.
- Ask which repos are involved and their paths on disk.
- Ask any clarifying questions needed to understand scope, constraints, or design preferences before exploring the code.

**2. Explore each repo**
For each repo:
- Understand the folder structure, conventions, and patterns in use.
- Identify where the new feature fits — which files will need to change, which patterns to follow.
- Note any existing similar features that can be used as reference.

**3. Draft the plan**
Produce a clear implementation plan per repo:
- What changes are needed and where
- New files to create
- Existing files to modify
- Any migrations, config changes, or env vars required
- Branch name for each repo (follow `rules/commits.md` naming)

Present the combined plan and ask the user to approve, suggest changes, or ask questions before any code is written.

**4. Implement**
On approval, for each repo:
- Pull latest main
- Create a branch
- Implement the changes following the plan and the project's existing conventions
- Commit frequently with clear messages following `rules/commits.md`

**5. Raise PRs**
For each repo, raise a PR using `gh pr create` with:
- A title describing what the PR does
- A description explaining *why* the change is being made (per `rules/prs.md`)
- Any relevant context about how this connects to changes in the other repo(s)

## Rules

- `${CLAUDE_PLUGIN_ROOT}/.claude/rules/commits.md` — branch naming and commit messages
- `${CLAUDE_PLUGIN_ROOT}/.claude/rules/prs.md` — PR descriptions must explain why
- `${CLAUDE_PLUGIN_ROOT}/.claude/rules/code.md` — follow existing conventions in each repo
- `${CLAUDE_PLUGIN_ROOT}/.claude/rules/testing.md` — add tests for new logic
- `${CLAUDE_PLUGIN_ROOT}/.claude/rules/secrets.md` — never hardcode secrets

## Common mistakes

- Writing code before fully understanding the existing codebase structure
- Not asking enough clarifying questions upfront — ambiguity caught late costs more
- Implementing without approval on the plan
- Making the PR description a list of what changed rather than why
- Skipping tests for new logic
