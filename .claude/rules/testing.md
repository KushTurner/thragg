# Testing

## Philosophy

- Test real logic, not mocks. Mocks should be a last resort, not a default.
- Use the testing pyramid as a guide — most tests are unit, fewer are integration, fewer still are e2e.
- Prefer real dependencies over fakes where practical. Use tools like Testcontainers to run real databases and services in tests. Never substitute a different technology (e.g. SQLite in place of Postgres) — different constraints mean different behaviour.
- TDD is encouraged to drive structure, but not required.

## Writing tests

- Follow Arrange / Act / Assert.
- One clear assertion per test, or a small focused group. Not ten different assertions per test.
- Tests must be deterministic — no randomness, no time-dependent behaviour, no reliance on execution order.
- Tests must never share state. Each test sets up and tears down its own state.
- Use testing libraries to reduce boilerplate setup where it genuinely helps.

## Scope

- Avoid unnecessary tests. Don't test framework behaviour, trivial getters, or things that can't fail.
- Tests should not constrain you from introducing new code. If a test breaks because of unrelated new code, the test is too tightly coupled.
