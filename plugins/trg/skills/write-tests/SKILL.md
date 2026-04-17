---
description: Write tests for existing code following project testing principles
---

# Write Tests

## Overview

Writes tests for existing code following the project's testing rules. Focuses on real logic over mocks, arrange/act/assert structure, isolated state, and deterministic outcomes. Tests should give you confidence, not just coverage numbers.

## Tools

- `Read` — reading the code to be tested and existing test patterns in the repo
- `Write` / `Edit` — writing new tests
- `Bash` — running tests to verify they pass

## When to use

- Adding tests to existing untested code
- A feature is complete and needs tests written before merging
- Improving weak or overly-mocked tests

## Examples

```
/trg:write-tests
```

- "Write tests for the payment service"
- "This handler has no tests, add them"
- "These tests are all mocks, rewrite them with Testcontainers"

## Steps

1. Read the code to be tested. Understand what it does before writing anything.
2. Check what testing tools and libraries are already used in the repo — follow them, don't introduce new ones.
3. Check if the project uses Testcontainers or has real service dependencies — use them rather than mocking.
4. Identify what is worth testing: real logic, edge cases, error paths. Skip trivial getters, framework behaviour, and things that cannot fail.
5. Write tests following arrange / act / assert. Each test: one clear outcome, no shared state, deterministic.
6. Run the tests and confirm they pass.
7. Ask the user to review before finalising.

## Rules

- `${CLAUDE_PLUGIN_ROOT}/plugins/trg/rules/testing.md` — governs all test decisions
- `${CLAUDE_PLUGIN_ROOT}/plugins/trg/rules/code.md` — tests are code too

## References

- `references/go_test.go` — Go test patterns with Testcontainers and Testify
- `references/example.test.ts` — TypeScript test patterns with Vitest and Testcontainers

## Common mistakes

- Mocking the database when Testcontainers or a real DB is available
- Writing tests that share state between test cases
- One test asserting ten different things — split into focused tests
- Testing implementation details instead of behaviour
- Writing tests that pass today but break when unrelated code changes
- Adding a new test library when one is already in use in the repo
