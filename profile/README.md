<div align="center">

<img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/banner.png" alt="kprompt — Talk to Your Cluster" width="100%" />

<br />

<img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/logo.png" alt="kprompt logo" width="96" />

# kprompt

### 💬 Talk to Your Cluster.

Open-source AI CLI to control **Kubernetes** with natural language.  
Describe what you want — deploy, scale, rollback, inspect, or explain — and kprompt turns it into a plan you can review before anything touches the cluster.

[![Release](https://img.shields.io/github/v/release/kprompt/kprompt?style=for-the-badge&color=2563eb&logo=github)](https://github.com/kprompt/kprompt/releases)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-22c55e?style=for-the-badge)](https://github.com/kprompt/kprompt/blob/main/LICENSE)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://github.com/kprompt/kprompt)
[![Website](https://img.shields.io/badge/Website-kprompt.ai-0ea5e9?style=for-the-badge)](https://kprompt.ai)

<br />

```bash
curl -fsSL https://kprompt.ai/install | bash
```

[🌐 Website](https://kprompt.ai) · [📦 CLI](https://github.com/kprompt/kprompt) · [🚀 Releases](https://github.com/kprompt/kprompt/releases) · [📖 Usage](https://kprompt.ai/#usage)

</div>

---

## ✨ What you can say

```bash
kprompt "list deployments"                    # 👀 inspect
kprompt "show pods" -n default                # 📋 read-only
kprompt "deploy redis"                        # 🧩 plan first
kprompt "deploy nginx" --approve              # ✅ apply
kprompt "scale api to 10"                     # 📈 plan
kprompt "scale api to 10" --approve           # ⚡ apply
kprompt "rollback payment-api" --approve      # ⏪ undo
kprompt "explain why payment-api is crashing" # 🔍 debug
```

> 💡 **Reads** run immediately. **Mutations** show a plan first, then ask `y/N` on a TTY — or pass `--approve`.

---

## ⚙️ How it works

```text
🗣️ Prompt  →  🧠 Intent  →  📝 Plan  →  🛡️ Safety  →  ✅ Approval  →  ⚙️ Executor  →  ⎈ Kubernetes
```

| Step | What happens |
|------|----------------|
| 1️⃣ **Prompt** | Describe the change in plain English |
| 2️⃣ **Plan** | Intent is parsed into concrete cluster actions |
| 3️⃣ **Safety** | Wipe / delete-everything style prompts are **hard-denied** |
| 4️⃣ **Apply** | Review, then execute against your live kubeconfig |

Same kubeconfig as `kubectl`. No separate agent or control plane to install. 🎯

---

## 🚀 What’s in v0.2

| | Capability | Status |
|---|------------|--------|
| 🧩 | Deploy (redis / nginx shortcuts) | Plan → safety → apply |
| 📈 | Scale | Plan → safety → apply |
| ⏪ | Rollback | Plan → safety → apply |
| 👀 | Get / list (pods, deployments, services) | Read-only |
| 🔍 | Explain-lite (status + events) | Read-only |
| ⚙️ | `kprompt config` show / set | `~/.kprompt/config.yaml` |
| 🤖 | Multi-LLM providers | OpenAI, Anthropic, Gemini, Groq, Ollama… |
| ✅ | Interactive approval | TTY `y/N`, or `--approve` |

Apache-2.0 licensed · Kind E2E under `go test -tags=e2e`

---

## 🏁 Quick start

### 1️⃣ Install

```bash
curl -fsSL https://kprompt.ai/install | bash
# fallback 🛟
curl -fsSL https://cdn.jsdelivr.net/gh/kprompt/kprompt@v0.2.0/install/install.sh | bash
```

### 2️⃣ Connect your cluster ☸️

Point kubeconfig at a cluster (`~/.kube/config` or `KUBECONFIG`).

### 3️⃣ Set an LLM API key 🔑

```bash
export KPROMPT_OPENAI_API_KEY=sk-...          # openai (default)
export KPROMPT_ANTHROPIC_API_KEY=sk-ant-...   # anthropic
export KPROMPT_GEMINI_API_KEY=...             # gemini
export KPROMPT_GROQ_API_KEY=...               # groq
# local 🦙: kprompt --provider ollama --model llama3.2 "..."
```

### 4️⃣ Optional defaults 💾

```bash
kprompt config set provider gemini
kprompt config set model gemini-2.0-flash
kprompt config set namespace default
kprompt config
```

> 🔐 API keys stay in your environment — never written to `config.yaml`.

### 5️⃣ Run a prompt 🎉

```bash
kprompt "list deployments"
kprompt "deploy redis"
```

---

## 💡 Why kprompt

| ❌ Traditional Kubernetes | ✅ With kprompt |
|---------------------------|-----------------|
| Write YAML & memorize kubectl flags | Describe what you want |
| Copy manifests & hunt docs | AI plans the change |
| Manual debugging loops | Explain from live status + events |
| Apply first, hope later | Review plan, then approve |

---

## 📚 Repositories

| | Repo | Role |
|---|------|------|
| 🖥️ | [**kprompt**](https://github.com/kprompt/kprompt) | CLI product — Go, Apache-2.0 |
| 🌐 | [**kprompt-website**](https://github.com/kprompt/kprompt-website) | Product site + `/install` endpoint |

<p align="center">
  <img src="https://raw.githubusercontent.com/kprompt/.github/main/profile/assets/og.png" alt="kprompt logo mark" width="160" />
</p>

---

<div align="center">

### 🔗 Links

[Website](https://kprompt.ai) · [CLI docs](https://github.com/kprompt/kprompt) · [Providers](https://github.com/kprompt/kprompt/blob/main/docs/providers.md) · [v0.2.0](https://github.com/kprompt/kprompt/releases/tag/v0.2.0) · [Issues](https://github.com/kprompt/kprompt/issues)

**⭐ Star the CLI · 🛠️ PRs welcome · 💬 Talk to Your Cluster.**

</div>
