# thragg

Sync shared Claude agent config (rules, `CLAUDE.md`, `settings.json`) across multiple repos from a single source of truth.

## Install

```bash
./install.sh
```

This builds the binary and symlinks it to `/usr/local/bin/thragg`.

## Usage

### 1. Initialize

Run this once inside your thragg repo to configure which repos it manages:

```bash
thragg init
```

Creates a `.thragg.json` in the current directory with your project name and target repo paths. This file is local to your machine and not committed.

### 2. Apply

Push your `CLAUDE.md`, `.claude/settings.json`, and `rules/` into every configured repo:

```bash
thragg apply
```

### 3. List

Show the repos currently configured:

```bash
thragg list
```

## What gets copied

| Source | Destination |
|---|---|
| `CLAUDE.md` | `<repo>/CLAUDE.md` |
| `.claude/settings.json` | `<repo>/.claude/settings.json` |
| `rules/` | `<repo>/rules/` |
