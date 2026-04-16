---
description: Review or plan software architecture following project principles
---

# Architecture

## Overview

Helps review or plan software architecture decisions. Applies a set of guiding principles focused on simplicity, statelessness, and avoiding premature complexity.

## When to use

- Starting a new project and deciding on the initial stack or structure
- Planning a new feature that might introduce infrastructure (queues, caching, WebSockets, etc.)
- Reviewing existing architecture for over-engineering or unnecessary complexity

## Examples

```
/thragg:architecture
```

- "I'm building a notifications feature, should I use WebSockets or polling?"
- "We're getting slow queries, should I add Redis?"
- "Review the architecture of this service"

## Steps

1. Ask what the user is building or deciding, if not already clear from context.
2. Identify any proposed or existing components that are premature given the current stage.
3. Suggest the simplest architecture that solves the problem.
4. Flag areas where complexity may be needed later, so the current design leaves room without over-engineering now.
5. If reviewing existing architecture, call out what is working well and what could be simplified.

## Principles to apply

**Start simple, expand later**
Don't introduce infrastructure (caching, Redis, queues, CDNs, etc.) until there is a clear, demonstrated need. No users yet means no need for scale.

**Prefer stateless over stateful**
Favour HTTP over persistent connections. Don't reach for WebSockets if HTTP polling or SSE would do the job. Avoid shared mutable state at the system level.

**Avoid framework lock-in**
Frameworks are fine when they add genuine value. Don't build everything around one — code should be understandable without deep framework knowledge.

**Complexity budget**
Every added component has a maintenance cost. Only introduce it when the problem genuinely requires it.

## Common mistakes

- Adding caching before there is a performance problem
- Reaching for WebSockets when the use case is one-directional or infrequent
- Introducing a message queue for a single consumer with no latency requirement
- Designing for millions of users before having any
