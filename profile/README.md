# kprompt

**Talk to Your Cluster.**

Open-source AI CLI to control Kubernetes with natural language — deploy, scale, rollback, explain, and inspect workloads from a single prompt.

```bash
curl -fsSL https://kprompt-website.vercel.app/install | bash

kprompt "list deployments"
kprompt "deploy redis"
kprompt "scale api to 10" --approve
kprompt "rollback payment-api" --approve
kprompt "explain why payment-api is crashing"
```

## Repositories

| Repo | What it is |
|------|------------|
| [**kprompt**](https://github.com/kprompt/kprompt) | CLI — plan → safety → apply against your cluster |
| [**kprompt-website**](https://github.com/kprompt/kprompt-website) | Product site + install endpoint |

## Links

- Website: [kprompt-website.vercel.app](https://kprompt-website.vercel.app)
- Latest release: [v0.2.0](https://github.com/kprompt/kprompt/releases/tag/v0.2.0)
- Install fallback: `curl -fsSL https://cdn.jsdelivr.net/gh/kprompt/kprompt@v0.2.0/install/install.sh | bash`
