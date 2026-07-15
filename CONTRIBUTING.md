# Contributing to kprompt

Thanks for helping make **Talk to Your Cluster** better.

## Where to contribute

| Area | Repo |
|------|------|
| CLI (Go) | [`kprompt/kprompt`](https://github.com/kprompt/kprompt) |
| Website / install endpoint | [`kprompt/kprompt-website`](https://github.com/kprompt/kprompt-website) |
| Org profile & defaults | [`kprompt/.github`](https://github.com/kprompt/.github) |

Architecture planning is private; product discussions belong on public issues.

## Before you start

1. Search [existing issues](https://github.com/kprompt/kprompt/issues) for duplicates.
2. For bugs / features, open an issue with the templates (or discuss first for large changes).
3. Keep PRs focused — one concern per PR.

## Development (CLI)

```bash
git clone https://github.com/kprompt/kprompt.git
cd kprompt
go test ./...
go build -o bin/kprompt ./cmd/kprompt
./bin/kprompt version
```

Optional kind E2E:

```bash
go test -tags=e2e ./test/e2e/
```

## Pull requests

- Describe **why** the change matters, not only what changed
- Add / update tests when behavior changes
- Do not commit secrets, API keys, or kubeconfigs
- Match existing Go / TS style in the target repo
- Link related issues

## Security

Do not report vulnerabilities in public issues. See [SECURITY.md](./SECURITY.md).

## Code of Conduct

Participation is governed by our [Code of Conduct](./CODE_OF_CONDUCT.md).
