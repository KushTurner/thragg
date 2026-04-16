# Database

## Data access

- No ORMs. Prefer code generation over raw SQL — e.g. sqlc for Go.
- Raw SQL is acceptable for one-off scripts or cases where generated queries genuinely don't fit, but it should be rare.

## Queries and indexes

- Write queries that can leverage indexes in the future — avoid patterns that would prevent index use (e.g. wrapping indexed columns in functions, leading wildcards in LIKE).
- Don't add indexes prematurely, but flag when a query would benefit from one as the data grows.

## Technology

- Use the same database technology in tests as in production. Don't substitute (e.g. SQLite in place of Postgres) — different engines have different constraints and behaviour.
