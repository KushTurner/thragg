# Dependencies & Versions

## Never assume versions

Training data has a knowledge cutoff. Any version assumed from memory may be outdated. Before specifying a version for any tool, language, package, or action, look it up.

This applies to:
- Language versions (Go, Node, Python, Rust, etc.)
- GitHub Actions (e.g. `actions/checkout`, `docker/build-push-action`)
- Package dependencies
- Docker base images
- CLI tools

## How to find the latest version

Use WebSearch or WebFetch to check official sources before pinning a version:
- GitHub releases page for the tool
- Official docs or package registry (npm, pkg.go.dev, PyPI, crates.io, etc.)
- Docker Hub for base images

## Check the docs too

Once you have the latest version, check its official documentation before using its API or configuration. Don't assume syntax or options from an older version — they may have changed. Hallucinating outdated config is worse than admitting uncertainty.

## When in doubt, say so

If you cannot verify the latest version or docs, say so explicitly rather than guessing. Offer to search or ask the user to confirm.
