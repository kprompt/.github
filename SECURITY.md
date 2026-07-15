# Security Policy

## Supported versions

Security fixes are applied to the **latest release** of [`kprompt/kprompt`](https://github.com/kprompt/kprompt/releases). Older tags are not patched.

## Reporting a vulnerability

**Do not** open a public issue for security problems.

Please use [GitHub Private Vulnerability Reporting](https://github.com/kprompt/kprompt/security/advisories/new) on the CLI repository.

Include:

- Affected version (`kprompt version`) and install method
- Description of the issue and impact
- Steps to reproduce (PoC if available)
- Whether the issue relates to the CLI, install script, or website `/install` endpoint

We aim to acknowledge reports within **7 days** and share a remediation plan when confirmed.

## Scope

In scope:

- `kprompt` CLI binary and source
- Install script (`install/install.sh` and the website install route)
- Credential / config handling (`~/.kprompt`, kubeconfig usage, LLM API keys)

Out of scope:

- Compromised third-party LLM providers or API keys you leave in env/shell history
- Misconfigured Kubernetes RBAC or cluster-admin kubeconfigs
- Issues that only affect intentionally dangerous `--approve` misuse after a clear plan was shown

## Safe use notes

- Prefer plan-only runs before `--approve`
- Never put API keys in `~/.kprompt/config.yaml` (env vars only)
- Treat wipe / broad-delete prompts as denied by design; report if one is still accepted
