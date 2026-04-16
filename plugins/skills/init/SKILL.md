---
description: Initialise a repo with shared Claude rules and settings
---

# Init

## Overview

Sets up the current repo with shared thragg rules and settings. Copies rule files into `rules/`, creates or appends to `CLAUDE.md`, and merges the shared permissions block into `.claude/settings.json` without overwriting existing config.

The plugin root is at `${CLAUDE_PLUGIN_ROOT}`.

## When to use

When setting up a repo with thragg for the first time, or when re-syncing a repo after thragg rules have been updated.

## Examples

```
/thragg:init
```

Run this in the root of the repo you want to set up.

## Steps

**1. Copy rules**
Copy every file from `${CLAUDE_PLUGIN_ROOT}/rules/` into `./rules/`, creating the directory if needed. Overwrite existing files — thragg owns these.

**2. Update CLAUDE.md**
Check if `./CLAUDE.md` exists.
- If it does **not** exist, create it:
  ```
  # CLAUDE.md

  @rules/commits.md
  @rules/prs.md
  @rules/cli.md
  @rules/secrets.md
  @rules/code.md
  @rules/testing.md
  @rules/formatting.md
  @rules/database.md
  @rules/deployment.md
  ```
- If it **does** exist, check whether it already imports thragg rules (look for `@rules/commits.md`). If not, append:
  ```

  # Shared Rules (thragg)

  @rules/commits.md
  @rules/prs.md
  @rules/cli.md
  @rules/secrets.md
  @rules/code.md
  @rules/testing.md
  @rules/formatting.md
  @rules/database.md
  @rules/deployment.md
  ```

**3. Merge settings.json**
Read the permissions block from `${CLAUDE_PLUGIN_ROOT}/.claude/settings.json`. Check if `./.claude/settings.json` exists.
- If it does **not** exist, create `.claude/` and write the permissions block as the full file.
- If it **does** exist, merge: add any allow/deny entries not already present. Preserve all other keys (e.g. `mcpServers`) exactly as they are.

**4. Confirm**
Tell the user what was created or updated.

## Common mistakes

- Running this inside the thragg repo itself.
- Running from a subdirectory instead of the repo root.
