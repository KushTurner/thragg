# Code Reviewer

## Overview

Autonomously reviews a pull request against project rules and coding principles. Produces a structured report with severity-tagged issues directly in Claude — no GitHub context switching required.

## When to use

When you want a thorough review of a PR before merging. Works best on small, focused PRs — flag if the PR is too large to review meaningfully.

## Usage

```
/thragg:code-reviewer <PR number>
```

Or with a repo:
```
/thragg:code-reviewer <owner>/<repo>#<PR number>
```

## Steps

1. Fetch the PR diff, title, description, and changed files using `gh pr view` and `gh pr diff`.
2. Read the project README and any relevant docs to understand context.
3. Read the following files from the plugin root — these are your review guides:
   - `${CLAUDE_PLUGIN_ROOT}/plugins/agents/code-reviewer/references/review-criteria.md`
   - `${CLAUDE_PLUGIN_ROOT}/rules/code.md`
   - `${CLAUDE_PLUGIN_ROOT}/rules/testing.md`
   - `${CLAUDE_PLUGIN_ROOT}/rules/database.md`
   - `${CLAUDE_PLUGIN_ROOT}/rules/formatting.md`
4. Run the linter and formatter if available in the repo.
5. Produce a structured report in the following format:

---

### PR: <title>

**Summary**
One paragraph — what this PR does and your overall verdict.

⚠️ Breaking change: yes / no

---

**Issues**

Group by severity. For each issue:
- Severity tag (🔴 Critical / 🟡 Major / 🔵 Minor / 💬 Suggestion)
- File and line number
- What the issue is
- Why it is a problem

---

**Positives**
What was done well.

---

**Checklist**
A checkbox list of every issue raised, for tracking:
- [ ] 🔴 `path/to/file.ts:42` — description
- [ ] 🟡 `path/to/file.ts:87` — description

---

6. Ask the user which issues they want to address and help them work through them one by one.

## Common mistakes

- Reviewing without reading the README or existing code for context first
- Posting to GitHub — keep everything in Claude
- Flagging style issues that the linter/formatter should catch instead
- Stating opinions as facts — if uncertain, say so
