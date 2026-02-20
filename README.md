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

