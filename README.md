# School Management System

Quickstart (local development)

1. Copy the example env and optionally edit values:

```powershell
cp .env.example .env
```

2. Run locally (uses SQLite by default):

```powershell
Set-Location -Path "c:\Users\dell\school-management-system"
$env:DB_DRIVER = "sqlite"
$env:DB_PATH = "school.db"
go run ./cmd/server
```

3. Or use the helper script (PowerShell):

```powershell
./scripts/run_local.ps1
```

Notes
- The default local DB is SQLite (`school.db`) for quick startup.
- To use Postgres, set `DB_DRIVER=postgres` and configure DB env vars in `.env`.
- The server exposes a health endpoint at `http://localhost:8080/health` and API routes under `/api`.

Developer productivity
----------------------

- Pre-commit hooks: a sample hook is available at `.githooks/pre-commit` (PowerShell). To enable hooks locally run:

```powershell
git config core.hooksPath .githooks
```

- Linting: `.golangci.yml` contains default linter settings. Install with:

```powershell
./scripts/install_tools.ps1
# or on mac/linux
./scripts/install_tools.sh
```

- Makefile: common commands are available: `make build`, `make run`, `make dev`, `make test`, `make lint`, `make format`.

- Removing tracked local files: run `.	emplates\remove_tracked_local.ps1` (or follow README section) to stop tracking `school.db`, `logs/`, and `.env`.

These tools will help keep code formatted, linted, and decrease friction when running locally.

