# go-template

Starter for a Go service in the dDan4a layout: one entity (`note`) goes through every layer, so a new project starts from
working code, tests and CI instead of empty folders.

```bash
go run golang.org/x/tools/cmd/gonew@latest github.com/dDan4a/go-template github.com/dDan4a/<repo>
```

`gonew` copies the latest version and rewrites the module path in `go.mod` and every import, the generated mocks included.
Then rename or delete `note`, rewrite this README and keep what the project needs.

## Layout

| Path | What |
| --- | --- |
| `main.go`, `internal/cmd/` | cobra root; `internal/cmd/server` builds the dependencies and serves HTTP |
| `internal/config`, `internal/version` | config from env (`DATABASE_URL`, `HTTP_ADDR`), commit and build time from ldflags |
| `internal/domain` | entities, no transport or storage |
| `internal/repository/face.go` | storage interfaces and shared errors, `//go:generate minimock` |
| `internal/repository/note/pg/` | Postgres implementation: `repository.go` with `New`, one file per method |
| `internal/service/note/` | use cases, one file each, unit tests on minimock next to them |
| `internal/controller/http/` | net/http handlers, one file each; `converter/` maps JSON to domain |
| `internal/tests/` | shared test helpers, `mocks/` (generated, committed), `e2e/` (build tag `e2e`) |
| `migrations/schema.sql` | applied on start with `CREATE ... IF NOT EXISTS`; switch to goose at the first `ALTER` |

## Run

```bash
docker compose up -d   # Postgres on 5432, POSTGRES_PORT to change it
make run               # :8080, HTTP_ADDR to change it
make test e2e lint     # e2e needs the Postgres above
make generate          # mocks after an interface change
```

CI (`.github/workflows/ci.yml`) calls the reusable workflows of [dDan4a/ci](https://github.com/dDan4a/ci): lint, unit
tests, e2e against a Postgres service, govulncheck, `go mod tidy`, hadolint. `.golangci.yml` is its reference config.
