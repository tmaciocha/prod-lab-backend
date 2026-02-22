# prod-lab-backend

Minimal Go Http API for Kubernetes lab.

# Feaures
* REST endpoint /health
* Returns JSON status
* Built using standard Go net/http
* Containerized with multi-stage Docker build

# Run locally
`go run main.go`

# Build container
`podman build -t prod-lab-backend:0.1`

# Run container
`podman run --rm -p 8080:8080 prod-lab-backend:0.1`

# Test:
`curl http://localhost:8080/health`

---

# CI
This repository uses GitHub Actions to automatically build and publish the backend container image.

**What happens on every push to `main`?**
1. Github checks out the repository.
2. A Docker image is built using the Dockerfile.
3. The image is tagged with the commit SHA.
4. The image is pushed to GitHub Container Registry (GHCR)

Example image format:
`ghcr.io/<github-username>/prod-lab-backend:<commit-sha>`

This ensures:
* every commit produces a unique immutable image.
* No manual `podman build` is required. (I use podan not docker because of Fedora OS)
* The image can be pulled from anywhere (Kubernetes, VPS, local machine)

**Workflow location**
`.github/workflows/backend.yml`

**Trigger**
`push to main branch`

**Registry**
Images are stored in:
`ghcr.io`

You can view published images in:
Github -> Repository -> Packages



This CI pipeline ensures that container images are built in clean Ubuntu environment provided by GitHub Actions. 
This guarantees that builds are reproducible and independent from the developer's local machine.