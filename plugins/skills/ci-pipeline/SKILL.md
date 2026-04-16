---
description: Set up or review a GitHub Actions CI/CD pipeline
---

# CI Pipeline

## Overview

Sets up or reviews a GitHub Actions pipeline following project principles. Keeps pipelines simple, fast, and focused on a quick feedback loop.

## When to use

- Setting up CI/CD for a new repo
- Reviewing or improving an existing pipeline
- Adding a new workflow (e.g. a release or deploy workflow)

## Examples

```
/thragg:ci-pipeline
```

- "Set up a CI pipeline for this Go project"
- "Review our existing GitHub Actions setup"
- "Add a workflow for deploying to production on merge to main"

## Principles

**Keep it simple and fast**
Pipelines should give fast feedback. Avoid unnecessary steps, long waits, or complex conditional logic.

**Separate workflows by trigger**
One workflow per trigger — don't use `if` conditions to switch behaviour based on branch or tag inside a single workflow. A PR workflow and a main branch workflow are separate files.

**Bare minimum jobs**
Every pipeline must have at minimum: build, lint, test. Nothing ships without these passing.

**Reusable actions**
Common steps shared across workflows belong in `.github/actions/<action-name>/action.yml` as composite actions. Reference them across workflows rather than duplicating steps.

**Leverage the marketplace**
Use well-maintained marketplace actions over writing shell steps from scratch — e.g. `docker/build-push-action`, `actions/cache`, `actions/setup-go`, `actions/setup-node`.

**Cache aggressively**
Cache dependencies and build layers to keep pipelines fast. Docker layer caching in particular can cut build times significantly.

## Steps

1. Ask what the project does and what stack it uses, if not already clear.
2. Ask what workflows are needed (PR, main branch, release, deploy, etc.).
3. Check if `.github/workflows/` already exists and review any existing workflows.
4. Generate the appropriate workflow files using the references in this skill as a starting point.
5. Suggest caching strategies appropriate for the stack.
6. Flag any existing workflows that are overly complex or could be simplified.

## References

- `references/pr.yml` — PR workflow template (build, lint, test)
- `references/main.yml` — main branch workflow template (build, push)
- `references/composite-action.yml` — example reusable composite action

## Common mistakes

- Putting PR and main branch logic in the same workflow with `if` conditions
- Not caching dependencies or Docker layers
- Writing long shell scripts inline instead of using marketplace actions
- Adding steps that aren't necessary at this stage (e.g. security scanning before you have users)
