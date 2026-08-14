# Hand-Edits to Generated SDK Code

This file catalogs commits that modified files inside `v24/api/keyfactor/v{1,2}/` (and, in future, other version directories) after their initial generation. These files carry generator-output "DO NOT EDIT" headers, but the project's `.openapi-generator-ignore` files are empty — no protection mechanism is in place. **Without the right templates and swagger patches, naive regeneration would silently drop every hand-edit listed below.**

## Conventions

For each file, hand-edits are listed in commit order (oldest first). Each entry notes:

- **Commit SHA + subject** — recover the full diff with `git show <sha> -- <path>`.
- **What it changed** — brief description.
- **Reproduced by upstream swagger?** — Yes if the swagger definition already implies the same shape; No if it does not (i.e. this is pure Go logic with no swagger counterpart).
- **Regression test pins this?** — Yes if a `*_test.go` test would fail without the hand-edit.
- **Action on regen** — `preserve` (must be re-applied post-regen), `verify` (re-check whether reproduced), `obsolete` (intentional removal), `docs-only` (no behavioral impact).

---

## v24 (out of scope for any current regen — no v24 swagger has been supplied)

### `v24/api/keyfactor/v1/client.go` + `v24/api/keyfactor/v2/client.go`

1. **`2a6c5b4`** — *fix(v24): plumb Server.ClientTimeout into rebuilt auth config* — inside `buildHttpClientV2()`, adds `HttpClientTimeout: cfg.ClientTimeout` to the `baseConfig := auth_providers.CommandAuthConfig{...}` struct literal in both files. Without it, `Server.ClientTimeout` (added upstream by `keyfactor-auth-client-go` to fix [issue #51](https://github.com/Keyfactor/keyfactor-auth-client-go/issues/51)) was silently dropped when this SDK rebuilt its own `CommandAuthConfig`, so every caller — including the Terraform provider's `request_timeout` setting — fell back to `auth_providers.DefaultClientTimeout` (60s) regardless of what was configured. This surfaced as `net/http: timeout awaiting response headers` on long-running calls such as PFX enrollment.
   - Reproduced by upstream swagger: **No** — pure Go logic, no swagger counterpart.
   - Pinned by test: **Yes** — `TestBuildHttpClientV2_ClientTimeoutPropagation` in `v24/api/keyfactor/v1/client_test.go` and `v24/api/keyfactor/v2/client_test.go` calls `buildHttpClientV2()` against a fake Command server and asserts the resulting `CommandAuthConfigBasic.HttpClientTimeout` and derived `BuildTransport().ResponseHeaderTimeout` reflect the configured value.
   - Action: **preserve**.
