# Contributing

Thanks for contributing to sshctl.

## Development

```bash
go test ./...
go build -o bin/sshctl .
```

Cross-build:

```bash
make dist
# or on Windows:
powershell -File scripts/build.ps1
```

## Pull requests

- Keep changes focused; avoid unrelated refactors.
- Add/adjust tests for config, search, and crypto helpers when behavior changes.
- Do not commit `~/.sshctl/servers.json`, binaries under `bin/`/`dist/`, or real credentials.
- Run `go test ./...` before opening a PR.

## Code style

Follow standard Go formatting (`gofmt` / `go fmt ./...`).
