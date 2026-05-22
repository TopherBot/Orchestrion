# Orchestrion 🚀

**Orchestrion** is a lightweight, extensible CI/CD orchestration platform you can run on‑prem or in the cloud.

## Key Features
- **Pipeline‑as‑code** – pipelines are simple YAML files (`orchestrion.yml`).
- **Multi‑language support** – runs Docker containers, supports Node, Python, Go, Rust, etc.
- **Built‑in DevSecOps** – integrates `pre‑commit`, `detect‑secrets`, `Trivy`, `CodeQL`, SBOM generation.
- **Scalable workers** – dynamic Docker‑in‑Docker workers with resource limits.
- **API‑first** – REST + OpenAPI spec for automation and UI integration.
- **Zero‑trust auth** – JWT with RBAC, OIDC support.
- **Observability** – Prometheus metrics, Grafana dashboards, GitHub Actions‑style UI.

## Quick Start (Docker)
```bash
# Clone the repo
git clone https://github.com/yourorg/orchestrion.git && cd orchestrion

# Build and start the service (Docker Compose)
docker compose up --build -d

# The API is available at http://localhost:8080
```

## Project Structure
```
├─ .github/                # CI/CD for this repo
│   └─ workflows/ci.yml    # Lint, test, security, build, Docker image
├─ cmd/orchestrion/        # Main binary entrypoint
├─ internal/               # Core engine (pipeline parser, executor, auth)
├─ api/                    # OpenAPI spec (v1.yaml)
├─ examples/               # Sample pipelines
├─ Dockerfile
├─ go.mod & go.sum
├─ .pre-commit-config.yaml
├─ .gitignore
├─ LICENSE
├─ SECURITY.md
└─ README.md
```