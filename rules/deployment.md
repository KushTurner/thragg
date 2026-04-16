# Deployment

- Software must be deployable via Docker. A `Dockerfile` (and `docker-compose.yml` where relevant) should be included.
- Local development should be runnable without complex setup — a single command ideally (e.g. `docker compose up`, `make dev`).
- Don't assume a specific cloud provider unless the project already has one established.
