# Commit Rules

- Never commit directly to `main`. Always branch first.
- Before starting any new feature branch, pull the latest `main` from remote first.
- Always squash and merge when merging PRs.
- Never add Claude co-author attribution ("Co-Authored-By: Claude ...") to commit messages.
- Commit frequently — small, focused commits make it easy to backtrack if something goes wrong.

## Message format

```
<type>: <message>
```

Message should be past tense and describe what changed, e.g. `feat: Added user authentication`.

**Types:**
- `feat` — new feature
- `fix` — bug fix
- `refactor` — code change that isn't a fix or feature
- `test` — adding or updating tests
- `docs` — documentation only
- `ci` — CI/CD pipeline changes
- `config` — configuration changes
- `chore` — maintenance, dependency updates, etc.
