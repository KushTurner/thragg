# Formatting & Linting

## Before every commit

Check how the project runs its formatter and linter, then run them:

- Look for scripts in `package.json` (e.g. `format`, `lint`) and use those
- Check for a `Makefile` with relevant targets
- Fall back to the language's built-in tool only if no project-level script exists:
  - **Go** — `gofmt -w .` and `go vet ./...`
  - **JavaScript / TypeScript** — Prettier and ESLint
  - **Python** — Ruff or Black + Flake8
  - **Rust** — `rustfmt` and Clippy

Always use the project's own tooling over running formatters directly. If a formatter config exists (`.prettierrc`, `pyproject.toml`, etc.), respect it.

## If no formatter is set up

Suggest adding one appropriate for the stack. Don't introduce one without asking.
