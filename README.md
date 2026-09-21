# modtest

A minimal Go module that demonstrates semantic-version-driven release automation.

The current version is baked into the binary at compile time via `go:embed` and is
automatically bumped on a schedule by a GitHub Actions workflow.

## Features

- **Embedded version** — `VERSION.txt` is embedded into the module at build time,
  so every binary knows exactly which version it is.
- **Auto version bumping** — a GitHub Actions workflow bumps the patch version every
  day (11:00 Beijing time), tags it, and pushes the release.
- **Tiny version tool** — `cmd/bump` reads `VERSION.txt` and prints the next patch
  version, using the official `golang.org/x/mod/semver` package.

## Installation

Install the command-line tool:

```bash
go install github.com/qq1u/modtest/cmd/modtest@latest
```

## Usage

### CLI

```bash
modtest
# 2026/09/21 22:40:20 Version: v0.0.1
```

### As a library

```go
import "github.com/qq1u/modtest"

fmt.Println("current version:", modtest.Version)
```

## How versioning works

- `VERSION.txt` holds the current version (e.g. `v0.0.1`).
- `version.go` embeds it with `//go:embed VERSION.txt` and exposes it as the exported
  `modtest.Version`.
- `cmd/bump` parses the file, increments the patch number and prints the next version:

```bash
go run ./cmd/bump          # reads VERSION.txt, prints e.g. v0.0.2
```

## Automatic release pipeline

The workflow in [`.github/workflows/bump.yml`](.github/workflows/bump.yml) runs:

- **Every day at 03:00 UTC (11:00 Beijing time)**, or
- **Manually** via the *Run workflow* button (Actions → Bump → Run workflow).

Each run:

1. Computes the next version with `cmd/bump`.
2. Checks the remote tags via `git ls-remote` and skips if the version was already published.
3. Builds the whole module (`go build ./...`) — a failing build never gets released.
4. Writes the new version to `VERSION.txt`, commits, tags (`vX.Y.Z`) and pushes.

Published versions are immediately usable with Go tooling:

```bash
go get github.com/qq1u/modtest@latest
go list -m -versions github.com/qq1u/modtest
go install github.com/qq1u/modtest/cmd/modtest@latest
```

Note: after a tag is pushed, `proxy.golang.org` takes a few minutes to index it.

## License

MIT — see [LICENSE](LICENSE).