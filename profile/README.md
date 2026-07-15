# kprompt

**Talk to Your Cluster.**

Open-source AI CLI to control Kubernetes with natural language. Describe what you want — deploy, scale, rollback, inspect, or explain — and kprompt turns it into a plan you can review before anything touches the cluster.

```bash
curl -fsSL https://kprompt-website.vercel.app/install | bash
```

## What you can say

```bash
kprompt "list deployments"
kprompt "show pods" -n default
kprompt "deploy redis"
kprompt "deploy nginx" --approve
kprompt "scale api to 10"
kprompt "scale api to 10" --approve
kprompt "rollback payment-api" --approve
kprompt "explain why payment-api is crashing"
```

Reads run immediately. Mutations show a plan first, then ask `y/N` on a TTY — or pass `--approve` to apply without the prompt.

## How it works

**Prompt → Intent → Plan → Safety → Approval → Executor → Kubernetes**

1. **Prompt** — describe the change in plain English  
2. **Plan** — intent is parsed into concrete cluster actions  
3. **Safety** — destructive prompts (wipe cluster, delete everything, …) are hard-denied  
4. **Apply** — review, then execute against your live kubeconfig context  

Same kubeconfig as `kubectl`. No separate agent or control plane to install.

## What’s in v0.2

| Capability | Status |
|------------|--------|
| Deploy (e.g. redis / nginx shortcuts) | Plan → safety → apply |
| Scale | Plan → safety → apply |
| Rollback | Plan → safety → apply |
| Get / list (pods, deployments, services) | Read-only |
| Explain-lite (status + events) | Read-only |
| `kprompt config` show / set | Defaults in `~/.kprompt/config.yaml` |
| Multi-LLM providers | OpenAI, Anthropic, Gemini, Groq, Ollama, and more |
| Interactive approval | TTY `y/N`, or `--approve` |

MIT licensed. Kind E2E coverage under `go test -tags=e2e`.

## Quick start

1. Install the CLI (binary into `~/.local/bin`):

```bash
curl -fsSL https://kprompt-website.vercel.app/install | bash
# fallback: curl -fsSL https://cdn.jsdelivr.net/gh/kprompt/kprompt@v0.2.0/install/install.sh | bash
```

2. Point kubeconfig at a cluster (`~/.kube/config` or `KUBECONFIG`).

3. Set an LLM API key:

```bash
export KPROMPT_OPENAI_API_KEY=sk-...          # --provider openai (default)
export KPROMPT_ANTHROPIC_API_KEY=sk-ant-...   # --provider anthropic
export KPROMPT_GEMINI_API_KEY=...             # --provider gemini
export KPROMPT_GROQ_API_KEY=...               # --provider groq
# local: kprompt --provider ollama --model llama3.2 "..."
```

4. Optional defaults (no secrets written to disk):

```bash
kprompt config set provider gemini
kprompt config set model gemini-2.0-flash
kprompt config set namespace default
kprompt config
```

5. Run a prompt:

```bash
kprompt "list deployments"
kprompt "deploy redis"
```

## Why kprompt

| Traditional Kubernetes | With kprompt |
|------------------------|--------------|
| Write YAML and remember kubectl flags | Describe what you want |
| Copy manifests and hunt docs | AI plans the change |
| Manual debugging loops | Explain from live status + events |
| Apply first, hope later | Review plan, then approve |

## Repositories

| Repo | Role |
|------|------|
| [**kprompt**](https://github.com/kprompt/kprompt) | CLI product — Go, MIT |
| [**kprompt-website**](https://github.com/kprompt/kprompt-website) | Product site + `/install` endpoint |

## Links

- Website & usage guide: [kprompt-website.vercel.app](https://kprompt-website.vercel.app)
- CLI docs: [github.com/kprompt/kprompt](https://github.com/kprompt/kprompt)
- Providers: [docs/providers.md](https://github.com/kprompt/kprompt/blob/main/docs/providers.md)
- Latest release: [v0.2.0](https://github.com/kprompt/kprompt/releases/tag/v0.2.0)
- Issues & contributions welcome on the [CLI repo](https://github.com/kprompt/kprompt/issues)
