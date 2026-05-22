# Security Policy

## Reporting a Vulnerability
If you discover a security issue, please **do not** open a public issue. Instead, email us at **security@orchestrion.dev** with a detailed description. We will acknowledge receipt within 48 hours and aim to resolve the issue promptly.

## Supported Versions
We support the latest major release and the previous one. Older versions receive no security updates.

## Security Features
- **Secret detection** – `detect-secrets` runs on every commit via a pre‑commit hook.
- **Container scanning** – Trivy scans built Docker images.
- **Static analysis** – CodeQL runs on each PR.
- **SBOM** – `syft` generates an SPDX SBOM for every release.
- **OpenSSF Scorecard** – Automated scorecard badge in the repository.

## Responsible Disclosure
We follow a 90‑day disclosure timeline. After the fix is released, feel free to publish details.

---
*This repository follows the OpenSSF best‑practice checklist.*