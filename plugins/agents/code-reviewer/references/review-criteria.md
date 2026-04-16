# PR Review Criteria

## Scope

- Is the PR small and scoped to a single concern? Flag if it touches too many files — a large diff is hard for a human to reason about and review meaningfully.
- Is it a breaking change? Call this out explicitly at the top of the report.

## Code quality

- Is the code easy to reason about? Flag anything that requires significant mental effort to follow.
- Are there bugs or logic errors?
- Is there dead code (unreachable, unused, commented out)?
- Are there spelling mistakes in identifiers, strings, or comments?
- Is the code consistent with the rest of the codebase — naming, structure, patterns?
- Are there edge cases not handled?
- Is there unnecessary complexity or over-engineering for the current stage?
- Is there premature optimisation (e.g. caching, indexing, batching) with no demonstrated need?
- Are functions doing more than one thing?
- Is there tight coupling between DTOs and entities (e.g. returning DB models directly from handlers)?
- Are ORMs being used in handlers? Data access should not leak into the request layer.

## Testing

- Are there tests for the change? Expected for features, not necessarily for config changes.
- Do the tests cover the important logic and edge cases, not just the happy path?

## Security

- Are there any security concerns — unvalidated input, exposed sensitive data, insecure defaults, injection risks?

## Formatting & linting

- Run the linter and formatter if available in the repo. Note any issues. Flag that CI/CD should be catching this if it is not already.

## Documentation & config

- Is documentation updated if the behaviour changed?
- Is config (env vars, feature flags, defaults) updated and documented?

## Context

- Use the project README, docs, and existing code to understand the intent of the change before reviewing. Don't review in isolation.

## Reporting rules

- Reference the exact file and line number for every issue raised.
- Explain **why** something is a problem, not just what it is.
- If uncertain about something, say so — don't present guesses as facts.
- Assign a severity to every issue:
  - 🔴 **Critical** — must fix before merge (bugs, security, broken logic)
  - 🟡 **Major** — should fix before merge (missing tests, convention violations, architecture concerns)
  - 🔵 **Minor** — worth fixing but won't block (small improvements, naming)
  - 💬 **Suggestion** — optional, worth considering
