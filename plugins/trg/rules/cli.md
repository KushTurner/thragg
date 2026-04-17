# CLI Usage

Always prefer well-known CLIs over manual approaches:
- GitHub → `gh`
- AWS → `aws`
- GCP → `gcloud`
- Fly.io → `fly`
- Vercel → `vercel`
- Supabase → `supabase`
- Bitbucket → `bb`

If the appropriate CLI is not installed but is available for the technology in use, ask the user if they'd like to install it before proceeding with an alternative.

## Preferred tools for common tasks

Always use the faster/smarter tool when available. If it is not installed, prompt the user to install it with a one-line reason before falling back to the alternative.

| Task | Prefer | Fallback | Why |
|---|---|---|---|
| Search file contents | `rg` | `grep` | Faster, respects `.gitignore`, better defaults |
| Find files | `fd` | `find` | Simpler syntax, faster, respects `.gitignore` |
| Parse JSON output | `jq` | raw output | Readable, queryable, avoids token waste |
| Explore directory structure | `tree` | `ls -R` | At-a-glance overview before diving in |
| Read large files | `head` / `tail` | full read | Read a portion first before deciding to load everything |

**When a preferred tool is missing:** suggest installing it and briefly explain why — e.g. "I'd recommend installing `rg` (ripgrep) — it's significantly faster than grep and ignores build artefacts automatically. Want me to install it?"

## Language-specific tools

Language runtimes and package managers (`go`, `npm`, `npx`, `node`, `python`, `pip`, etc.) are project-specific. Add them to `.claude/settings.local.json` per repo rather than the global allow list.
