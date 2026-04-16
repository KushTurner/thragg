# thragg

Sync shared Claude agent config (rules, `CLAUDE.md`, `settings.json`) across multiple repos from a single source of truth.

## Install

```bash
./install.sh
```

This builds the binary and symlinks it to `/usr/local/bin/thragg`.

## Usage

```bash
thragg apply ~/projects/repo1 ~/projects/repo2
```

Copies `CLAUDE.md`, `.claude/settings.json`, and `rules/` into each repo.

## Repo-specific settings

To add allow/deny rules on top of the global baseline without them being overwritten, put them in `.claude/settings.local.json` in the target repo. Claude Code merges this automatically.
