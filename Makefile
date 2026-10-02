GO ?= go

# Test binaries run one package at a time by default so the suite does not
# build every package at once. Raise the limit for a faster run:
# make test TEST_P=4
TEST_P ?= 1

.PHONY: test
test:
	$(GO) test -p $(TEST_P) ./...

# Every folder under examples/ that `make open` can walk.
EXAMPLES := $(patsubst examples/%/,%,$(wildcard examples/*/))

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
		if ! ls "examples/$$dir"/*.go >/dev/null 2>&1; then \
			echo "--- skip $$dir (no Go files)"; \
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
