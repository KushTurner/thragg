# Code Conventions

## Design principles

- Prefer functional approaches and avoid shared mutable state where practical. Avoid unreadable higher-order function chains.
- Follow SOLID principles when writing OOP code.
- DRY — but only after you've seen the same pattern three times. Don't abstract early.
- YAGNI — don't build for hypothetical future requirements.
- KISS — always prefer the simpler solution. If two approaches work, take the simpler one.

## Functions

- Each function should do one thing.
- Avoid comments — code should be self-explanatory. Exception: language conventions that expect them (e.g. Go doc comments on exported symbols).

## File size

- Aim for well under 500 lines per file. No hard limit, but if a file is growing large it's usually a sign it needs splitting.

## Conventions

- Match the conventions already used in the repo — naming, casing, file structure, patterns. Read the existing code before writing new code.
- Folder structure is project-dependent. Don't impose a structure that doesn't fit.
