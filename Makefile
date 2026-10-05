GO ?= go

# Test binaries run one package at a time by default so the suite does not
# build every package at once. Raise the limit for a faster run:
# make test TEST_P=4
TEST_P ?= 1

.PHONY: test
test:
	$(GO) test -p $(TEST_P) ./... ./examples/...

# `go build ./...` links an executable for every example on each run, takes
# about a minute, and leaves the binaries in this directory. `go vet` compiles
# every package, including every example, without linking. Raise the limit for
# a faster run: make build BUILD_P=4
BUILD_P ?= 1

.PHONY: build
build:
	$(GO) vet -p $(BUILD_P) ./... ./examples/...

# Every folder under examples/ that `make open` can walk.
EXAMPLES := $(patsubst examples/%/,%,$(wildcard examples/*/))

# Overlays and helpers that `make open` never walks. `desktop-cat` is a
# transparent click-through overlay, not a closable window, so run it alone:
#   go run ./examples/desktop-cat
# The form and input demos below now live in `input-lab`; the old folders
# stay for their tests, so open skips them. Run one alone when needed:
#   go run ./examples/login
SKIP_OPEN := desktop-cat bind bind-hooks clipboard controls editing forms input login states

# The folder named after `open`, if one was given.
OPEN_FROM := $(word 2,$(MAKECMDGOALS))

# Open examples one window at a time, from the named folder through the last
# folder under examples/:
#   make open crash              crash, then each later example, to the end
#   make open crash ARGS=-web    serve each in a browser instead
# A bare `make open` starts from the first folder. The next example opens
# after the previous window closes. Folders without Go files are skipped.
.PHONY: open
open:
	@start=0; \
	for dir in $(sort $(EXAMPLES)); do \
		if [ -z "$(OPEN_FROM)" ] || [ "$$dir" = "$(OPEN_FROM)" ]; then start=1; fi; \
		if [ $$start -eq 0 ]; then continue; fi; \
		case " $(SKIP_OPEN) " in *" $$dir "*) \
			if [ "$$dir" != "$(OPEN_FROM)" ]; then \
				echo "--- skip $$dir (run manually: go run ./examples/$$dir)"; \
				continue; \
			fi;; \
		esac; \
		if ! ls "examples/$$dir"/*.go >/dev/null 2>&1; then \
			echo "--- skip $$dir (no Go files)"; \
			continue; \
		fi; \
		if [ "$$($(GO) list -f '{{.Name}}' ./examples/$$dir 2>/dev/null)" != "main" ]; then \
			echo "--- skip $$dir (not a command)"; \
			continue; \
		fi; \
		echo "==> go run ./examples/$$dir $(ARGS)"; \
		$(GO) run ./examples/$$dir $(ARGS) || echo "--- examples/$$dir exited $$?"; \
	done; \
	if [ $$start -eq 0 ]; then \
		echo "make open: no folder examples/$(OPEN_FROM)"; \
		echo "examples: $(sort $(EXAMPLES))"; \
		exit 1; \
	fi

# `make open crash` turns crash into a second make goal, so each example
# folder needs a target of its own. This one does nothing.
.PHONY: $(EXAMPLES)
$(EXAMPLES):
	@:

# Unknown names fail with a pointer to `make open`.
%:
	@echo "make: unknown target '$@'; run 'make open <example>'" >&2; exit 1
