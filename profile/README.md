<div align="center">

<img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/banner.png" alt="kprompt — Talk to Your Cluster" width="100%" />

<br />

<img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/logo.png" alt="kprompt logo" width="96" />

# kprompt

### Talk to Your Cluster.

Open-source **AI Kubernetes CLI**: natural language compiles into a **reviewable plan**, then you approve before anything touches the cluster.

[![Release](https://img.shields.io/github/v/release/kprompt/kprompt?style=for-the-badge&color=2563eb&logo=github)](https://github.com/kprompt/kprompt/releases/latest)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-22c55e?style=for-the-badge)](https://github.com/kprompt/kprompt/blob/main/LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://github.com/kprompt/kprompt)
[![Website](https://img.shields.io/badge/Website-kprompt.ai-0ea5e9?style=for-the-badge)](https://kprompt.ai)

<br />

**Start here → [`github.com/kprompt/kprompt`](https://github.com/kprompt/kprompt)** · star · clone · try the demo

```bash
curl -fsSL https://kprompt.ai/install | bash
# or: brew install kprompt/tap/kprompt
```

[Website](https://kprompt.ai) · [Docs](https://kprompt.ai/docs) · [Discussions](https://github.com/kprompt/kprompt/discussions) · [Examples](https://github.com/kprompt/kprompt-examples)

</div>

---

## Observe agent demo (no API key)

Kind cluster, seven broken workloads, heuristic Observe agent — **zero LLM spend**, Autopilot **propose-only**:

<img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/kprompt-observe-demo.gif" alt="kprompt Observe agent on a broken kind cluster — incidents, health score, propose-only Autopilot" width="100%" />

```bash
git clone https://github.com/kprompt/kprompt-examples.git
cd kprompt-examples && make walkthrough
```

---

## What you can say

```bash
kprompt "scale api to 10"                     # plan first
kprompt "scale api to 10" --approve --wait    # apply after review
kprompt "explain why payment-api is crashing" # debug
kprompt "delete everything in the cluster"    # hard-denied
kprompt agent run -n payments --health --heuristic
```

Reads run immediately. Mutations show **Intent / Plan / Risk / Actions / Blast radius**, then ask `y/N` — or pass `--approve`.

---

## How it works

```text
Prompt → Intent → Plan → Safety → Approval → Executor → Kubernetes
```

| Step | What happens |
|------|----------------|
| **Prompt** | Describe the change in plain English |
| **Plan** | Typed actions, risk, blast radius |
| **Safety** | Wipe / delete-everything style prompts are hard-denied |
| **Apply** | Review, then execute against your local kubeconfig |

Same kubeconfig as `kubectl`. BYOK — your LLM keys stay local.

---

## What's in v0.5

| | Capability |
|---|------------|
| Plan → approve → apply CLI | Scale, deploy, rollback, Helm, explain, investigate |
| Observe agent | Namespace watch → Incident → gated Slack/webhook |
| Autopilot | **Propose-only** by default (never silent apply) |
| CI | `--output json` PlanResult for jq gates |
| Multi-context | Read fan-out + per-context mutate approval |

---

## Repositories

| Repo | Role |
|------|------|
| [**kprompt**](https://github.com/kprompt/kprompt) | CLI + Observe agent (start here) |
| [**kprompt-examples**](https://github.com/kprompt/kprompt-examples) | Kind demos — `make walkthrough` |
| [**kprompt-website**](https://github.com/kprompt/kprompt-website) | Site + `/install` |
| [**homebrew-tap**](https://github.com/kprompt/homebrew-tap) | `brew install kprompt/tap/kprompt` |

---

<div align="center">

[Website](https://kprompt.ai) · [CLI](https://github.com/kprompt/kprompt) · [Discussions](https://github.com/kprompt/kprompt/discussions) · [Good first issues](https://github.com/kprompt/kprompt/labels/good%20first%20issue) · [v0.5.0](https://github.com/kprompt/kprompt/releases/tag/v0.5.0)

**Star the CLI · PRs welcome · Talk to Your Cluster.**

</div>
