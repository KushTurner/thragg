# thragg

A Claude Code plugin that syncs shared agent rules and settings into any repo.

## Install

Add the marketplace and install the plugin:

```
/plugin marketplace add KushTurner/thragg
```

## Usage

Run this in any repo you want to set up:

```
/thragg:init
```

This will:
- Copy shared `rules/` into the repo
- Create or append to `CLAUDE.md` with the shared rule imports
- Merge the shared `permissions` block into `.claude/settings.json`, preserving any existing config (e.g. MCP servers)

## Repo-specific settings

To add allow/deny rules on top of the shared baseline without them being overwritten, put them in `.claude/settings.local.json`. Claude Code merges this automatically.
