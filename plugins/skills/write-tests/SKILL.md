---
description: Write tests for existing code following project testing principles
---

# Write Tests

## Overview

Writes tests for existing code following the project's testing rules. Focuses on real logic over mocks, arrange/act/assert structure, isolated state, and deterministic outcomes. Tests should give you confidence, not just coverage numbers.

## When to use

- Adding tests to existing untested code
- A feature is complete and needs tests written before merging
- Improving weak or overly-mocked tests

## Examples

```
/thragg:write-tests
```

- "Write tests for the payment service"
- "This handler has no tests, add them"
- "These tests are all mocks, rewrite them with Testcontainers"

## Steps

1. Read `${CLAUDE_PLUGIN_ROOT}/rules/testing.md` — this governs all test decisions.
2. Read the code to be tested. Understand what it does before writing anything.
3. Identify what is worth testing — real logic, edge cases, error paths. Skip trivial getters, framework behaviour, and things that cannot fail.
4. Check what testing tools and libraries are already used in the repo and follow them. Don't introduce a new testing framework.
5. Check if the project uses Testcontainers or has real service dependencies — use them if available rather than mocking.
6. Write tests following arrange / act / assert. Each test: one clear outcome, no shared state, deterministic.
7. Ask the user to review before finalising.

## Rules to apply

- `${CLAUDE_PLUGIN_ROOT}/rules/testing.md`
- `${CLAUDE_PLUGIN_ROOT}/rules/code.md` — tests are code too

## Common mistakes

- Mocking the database when Testcontainers or a real DB is available
- Writing tests that share state between test cases
- One test asserting ten different things — split into focused tests
- Testing implementation details instead of behaviour
- Writing tests that pass today but break when unrelated code changes
- Adding a new test library when one is already in use in the repo
