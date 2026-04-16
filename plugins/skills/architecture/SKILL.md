---
description: Review or plan software architecture following project principles
---

Help the user think through software architecture decisions for their project. Apply the following principles throughout.

## Principles

**Start simple, expand later**
- Don't introduce infrastructure (caching, Redis, queues, CDNs, etc.) until there is a clear, demonstrated need. No users yet means no need for scale.
- Solve the problem in front of you, not the hypothetical future one.

**Prefer stateless over stateful**
- Favour HTTP over persistent connections. Don't introduce WebSockets if HTTP polling or SSE would do the job.
- Avoid shared mutable state at the system level where possible.

**Avoid framework lock-in**
- Frameworks are fine when they add genuine value, but avoid building everything around one. The code should be understandable without deep framework knowledge.

**Complexity budget**
- Every added component (service, queue, cache, external dependency) has a maintenance cost. Only introduce it when the problem genuinely requires it.

## What to do

1. Ask the user what they are building or what decision they are trying to make, if not already clear from context.
2. Identify any proposed or existing components that are premature given the current stage of the project.
3. Suggest the simplest architecture that solves the problem.
4. Flag any areas where complexity may be needed in the future, so the current design leaves room for it without over-engineering now.
5. If reviewing existing architecture, point out what is working well and what could be simplified.
