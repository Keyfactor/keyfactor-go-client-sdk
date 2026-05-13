# keyfactor-go-client-sdk Makefile
#
# Regenerates Go SDK packages from OpenAPI swagger inputs.
#
# Generation is NEVER in-place. The workflow is:
#   1. make regen-prepare   — snapshot current code into .regen-baseline/
#   2. make regen-stage     — generate fresh code into .regen-staging/ (NEVER overwrites live)
#   3. make regen-diff      — structured diff + hand-edit cross-check against HAND_EDITS.md
#   4. (human review of .regen-diff.txt)
#   5. touch .regen-approved
#   6. make regen-apply     — swap staging into live; run `go test`
#   7. make regen-clean     — remove baseline/staging/diff/sentinel artifacts
#
# The bare `generate` target is a tripwire that errors out by design — preserves
# hand-edits cataloged in HAND_EDITS.md from being silently clobbered.

# === Configuration ===
COMMAND_VERSION           ?= v25
SWAGGER_V1                ?= swagger/Keyfactor-Command-$(COMMAND_VERSION)-v1.swagger.json
SWAGGER_V2                ?= swagger/Keyfactor-Command-$(COMMAND_VERSION)-v2.swagger.json
OPENAPI_GENERATOR_VERSION ?= 6.3.0

# === Derived paths ===
V1_LIVE     = $(COMMAND_VERSION)/api/keyfactor/v1
V2_LIVE     = $(COMMAND_VERSION)/api/keyfactor/v2
V1_STAGE    = .regen-staging/$(V1_LIVE)
V2_STAGE    = .regen-staging/$(V2_LIVE)
V1_BASELINE = .regen-baseline/$(V1_LIVE)
V2_BASELINE = .regen-baseline/$(V2_LIVE)

GENERATOR_JAR = openapi-generator-cli.jar

.PHONY: help regen-prepare regen-stage regen-diff regen-apply regen-clean \
        generate fmt deps-check

help:
	@echo "Targets:"
	@echo "  regen-prepare  — snapshot $(V1_LIVE) + $(V2_LIVE) into .regen-baseline/"
	@echo "  regen-stage    — generate into .regen-staging/ (NEVER touches live code)"
	@echo "  regen-diff     — structured diff + hand-edit cross-check; writes .regen-diff.txt"
	@echo "  regen-apply    — swap staging -> live (requires .regen-approved sentinel + passes go test)"
	@echo "  regen-clean    — remove .regen-baseline, .regen-staging, .regen-diff.txt, sentinel"
	@echo ""
	@echo "Config (override on CLI):"
	@echo "  COMMAND_VERSION=$(COMMAND_VERSION)"
	@echo "  SWAGGER_V1=$(SWAGGER_V1)"
	@echo "  SWAGGER_V2=$(SWAGGER_V2)"
	@echo "  OPENAPI_GENERATOR_VERSION=$(OPENAPI_GENERATOR_VERSION)"

# ---------- regen workflow ----------

regen-prepare: deps-check
	@echo ">>> Snapshotting $(V1_LIVE) and $(V2_LIVE) -> .regen-baseline/"
	@test -d $(V1_LIVE) || (echo "ERROR: $(V1_LIVE) does not exist"; exit 1)
	@test -d $(V2_LIVE) || (echo "ERROR: $(V2_LIVE) does not exist"; exit 1)
	@rm -rf .regen-baseline
	@mkdir -p $(dir $(V1_BASELINE)) $(dir $(V2_BASELINE))
	@cp -R $(V1_LIVE) $(V1_BASELINE)
	@cp -R $(V2_LIVE) $(V2_BASELINE)
	@echo "    baseline captured."

regen-stage: regen-prepare $(GENERATOR_JAR)
	@test -f $(SWAGGER_V1) || (echo "ERROR: $(SWAGGER_V1) not found"; exit 1)
	@test -f $(SWAGGER_V2) || (echo "ERROR: $(SWAGGER_V2) not found"; exit 1)
	@rm -rf .regen-staging
	@COMMAND_VERSION=$(COMMAND_VERSION) \
	 SWAGGER_V1=$(SWAGGER_V1) \
	 SWAGGER_V2=$(SWAGGER_V2) \
	 OUTPUT_DIR=.regen-staging \
	 GENERATOR_JAR=$(GENERATOR_JAR) \
	 scripts/regen.sh
	@echo ">>> Copying preserved files (regression tests etc.) from baseline into staging"
	@if [ -x scripts/check-hand-edits.sh ]; then \
		scripts/check-hand-edits.sh preserve $(V1_BASELINE) $(V2_BASELINE) $(V1_STAGE) $(V2_STAGE); \
	else \
		echo "    WARN: scripts/check-hand-edits.sh not executable; skipping preserve step"; \
	fi
	@echo ">>> Staging complete: $(V1_STAGE), $(V2_STAGE)"

regen-diff: regen-stage
	@echo ">>> Files in baseline missing from staging (CRITICAL):"
	@diff -rq $(V1_BASELINE) $(V1_STAGE) 2>/dev/null | grep "Only in $(V1_BASELINE)" || echo "    (none for v1)"
	@diff -rq $(V2_BASELINE) $(V2_STAGE) 2>/dev/null | grep "Only in $(V2_BASELINE)" || echo "    (none for v2)"
	@echo ""
	@echo ">>> Hand-edit cross-check (scripts/check-hand-edits.sh):"
	@if [ -x scripts/check-hand-edits.sh ]; then \
		scripts/check-hand-edits.sh check $(V1_STAGE) $(V2_STAGE); \
	else \
		echo "    WARN: scripts/check-hand-edits.sh not executable; skipping"; \
	fi
	@echo ""
	@echo ">>> Writing full unified diff -> .regen-diff.txt"
	@: > .regen-diff.txt
	@diff -ruN $(V1_BASELINE) $(V1_STAGE) >> .regen-diff.txt 2>/dev/null || true
	@diff -ruN $(V2_BASELINE) $(V2_STAGE) >> .regen-diff.txt 2>/dev/null || true
	@echo "    .regen-diff.txt: $$(wc -l < .regen-diff.txt) lines"
	@echo ""
	@echo "NEXT: review .regen-diff.txt manually, then 'touch .regen-approved' to allow apply."

regen-apply:
	@test -d .regen-staging || (echo "ERROR: nothing staged; run 'make regen-stage' first"; exit 1)
	@test -f .regen-approved || (echo "ERROR: .regen-approved sentinel missing; review .regen-diff.txt then 'touch .regen-approved'"; exit 1)
	@echo ">>> Swapping staging -> live"
	@rm -rf $(V1_LIVE) $(V2_LIVE)
	@mv $(V1_STAGE) $(V1_LIVE)
	@mv $(V2_STAGE) $(V2_LIVE)
	@rm -f .regen-approved
	@echo ">>> Running go test ./$(COMMAND_VERSION)/..."
	@cd $(COMMAND_VERSION) && go test ./... || (echo "ERROR: tests fail post-apply; restore from .regen-baseline/ if needed"; exit 1)
	@echo ">>> Apply complete. .regen-baseline/ retained for inspection."

regen-clean:
	@rm -rf .regen-baseline .regen-staging .regen-diff.txt .regen-approved
	@echo "    cleaned."

# Tripwire: bare `generate` would skip all safety
generate:
	@echo "ERROR: in-place generation is disabled."
	@echo "Use the staged workflow:"
	@echo "  make regen-prepare && make regen-stage && make regen-diff"
	@echo "  (human review of .regen-diff.txt)"
	@echo "  touch .regen-approved && make regen-apply"
	@exit 1

# ---------- supporting targets ----------

$(GENERATOR_JAR):
	@echo ">>> Downloading openapi-generator-cli $(OPENAPI_GENERATOR_VERSION)"
	@wget -q https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar -O $(GENERATOR_JAR)

deps-check:
	@command -v java >/dev/null 2>&1 || (echo "ERROR: 'java' not on PATH"; exit 1)
	@command -v wget >/dev/null 2>&1 || (echo "ERROR: 'wget' not on PATH"; exit 1)
	@command -v go >/dev/null 2>&1 || (echo "ERROR: 'go' not on PATH"; exit 1)
	@command -v python3 >/dev/null 2>&1 || (echo "ERROR: 'python3' not on PATH (needed by scripts/regen.sh)"; exit 1)
	@command -v jq >/dev/null 2>&1 || (echo "ERROR: 'jq' not on PATH (needed by scripts/regen.sh for swagger patches)"; exit 1)

fmt:
	@gofmt -w $(COMMAND_VERSION)/

# NOTE: generator invocation, swagger patches, namespace strip, and postprocess
# template substitution all live in scripts/regen.sh (adapted from
# Keyfactor/API-definitions @ 04b7b2c:go/command/openapi-generate.sh).
# The Makefile orchestrates the safety harness around it.
