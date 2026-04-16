# Deployment

- Software must be deployable via Docker. A `Dockerfile` (and `docker-compose.yml` where relevant) should be included.
- Local development should be runnable without complex setup — a single command ideally (e.g. `docker compose up`, `make dev`).
- Don't assume a specific cloud provider unless the project already has one established.

## Dockerfiles

- Keep Dockerfiles simple and readable.
- Don't over-optimise by hunting for the smallest possible base image — use the standard, well-known image for the language or runtime. Optimise later if there is a real reason to.
- Prefer images that include common utilities out of the box (e.g. ca-certs) over stripped-down ones that require manual setup.
- Multi-stage builds are a good pattern: build the binary or artifact in one stage, copy only what is needed into the final image. This keeps the final image clean without sacrificing readability.
