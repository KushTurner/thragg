---
description: Review or write error handling following best practices
---

# Error Handling

## Overview

Reviews or writes error handling code following established best practices. Errors should be explicit, carry context as they propagate, and be handled at exactly one layer — never swallowed, never duplicated.

## When to use

- Writing new error handling in any language
- Reviewing existing error handling for anti-patterns
- Debugging issues caused by lost error context

## Examples

```
/thragg:error-handling
```

- "Review the error handling in this file"
- "Help me handle errors correctly in this service"

## Principles

**Never ignore errors**
Every error must be checked and either handled or propagated. No silent discards.

**Add context as errors propagate**
Each layer should wrap the error with enough context to understand where and why it failed — without duplicating information already in the error.

**Handle at one layer**
Either log the error or return it — never both. Logging and returning causes the same error to appear multiple times. Handle once, at the highest appropriate level.

**Never panic for expected errors**
Panic is for truly unexpected states. Return errors for anything a caller could reasonably handle or recover from.

## Language-specific references

- `references/go.go` — Go error handling patterns and anti-patterns

## Common mistakes

- Using `%v` or `%s` instead of `%w` in Go's `fmt.Errorf` — breaks the error chain, `errors.Is` and `errors.As` stop working
- Logging the error and returning it — same error appears multiple times in logs
- Ignoring errors with `_` or no check at all
- Calling `log.Fatal` outside of `main` — skips deferred functions and exits immediately
- Comparing errors with `==` instead of `errors.Is` — breaks with wrapped errors
- Panicking in library or application code for recoverable errors
