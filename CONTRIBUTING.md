# Contributing to kprompt

Thanks for helping make **Talk to Your Cluster** better.

## Where to contribute

| Area | Repo |
|------|------|
| CLI + Observe agent (Go) | [`kprompt/kprompt`](https://github.com/kprompt/kprompt) — start with [CONTRIBUTING.md](https://github.com/kprompt/kprompt/blob/main/CONTRIBUTING.md) |
| Kind demos | [`kprompt/kprompt-examples`](https://github.com/kprompt/kprompt-examples) |
| Website / install | [`kprompt/kprompt-website`](https://github.com/kprompt/kprompt-website) |
| Org profile & defaults | this repo — [`kprompt/.github`](https://github.com/kprompt/.github) |
| Homebrew | [`kprompt/homebrew-tap`](https://github.com/kprompt/homebrew-tap) |

## First contribution path

1. Open [`good first issue`](https://github.com/kprompt/kprompt/labels/good%20first%20issue) on the CLI repo.
2. Comment `I'd like to take this`.
3. Follow the CLI [CONTRIBUTING.md](https://github.com/kprompt/kprompt/blob/main/CONTRIBUTING.md) checks.
4. Ship a small focused PR.

Questions / demos: [Discussions](https://github.com/kprompt/kprompt/discussions).

## Before you start

1. Search existing issues for duplicates.
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

Do not report vulnerabilities in public issues. See [SECURITY.md](./SECURITY.md) if present.

## Code of Conduct

Participation is governed by our [Code of Conduct](./CODE_OF_CONDUCT.md) when present. Be respectful either way.

## License

Contributions are accepted under the [Apache License 2.0](https://github.com/kprompt/kprompt/blob/main/LICENSE).
