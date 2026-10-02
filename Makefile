GO ?= go

# Test binaries run one package at a time by default so the suite does not
# build every package at once. Raise the limit for a faster run:
# make test TEST_P=4
TEST_P ?= 1

.PHONY: test
test:
	$(GO) test -p $(TEST_P) ./...
