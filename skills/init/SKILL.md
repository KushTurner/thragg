---
description: Initialise a repo with shared Claude rules and settings
---

Set up this repo with the shared thragg rules. The plugin root is at `${CLAUDE_PLUGIN_ROOT}`.

Follow these steps exactly:

## 1. Copy rules

Copy every file from `${CLAUDE_PLUGIN_ROOT}/rules/` into `./rules/` in the current repo, creating the directory if needed. Overwrite existing files — thragg owns these.

## 2. Update CLAUDE.md

Check if `./CLAUDE.md` exists in the current repo.

- If it does **not** exist, create it with the following content:
  ```
  # CLAUDE.md

  @rules/commits.md
  @rules/prs.md
  @rules/cli.md
  @rules/secrets.md
  ```
- If it **does** exist, check whether it already imports the thragg rules (look for `@rules/commits.md`). If not, append the following block to the end of the file:
  ```

  # Shared Rules (thragg)

  @rules/commits.md
  @rules/prs.md
  @rules/cli.md
  @rules/secrets.md
  ```

## 3. Merge settings.json

Read the permissions block from `${CLAUDE_PLUGIN_ROOT}/.claude/settings.json`.

Check if `./.claude/settings.json` exists in the current repo.

- If it does **not** exist, create `.claude/` and write the permissions block as the full file.
- If it **does** exist, merge: add any allow/deny entries from the plugin that are not already present. Preserve all other keys (e.g. `mcpServers`) exactly as they are.

## 4. Confirm

Tell the user what was created or updated.
