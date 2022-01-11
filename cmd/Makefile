include ../Makefile.defs

.PHONY: all
all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET):
	CGO_ENABLED=0 $(GO) $(GO_BUILD_FLAGS) build -o $(TARGET)

.PHONY: install
install: $(TARGET)
	$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(INSTALL) -m 0755 ./$(TARGET) $(DESTDIR)$(BINDIR)

.PHONY: clean
clean:
	rm -f $(TARGET)
	rm -f cpu.prof mem.prof

.PHONY: test
test:
	$(GO) test $(GO_TEST_FLAGS) ./...

.PHONY: bench
bench:
	$(GO) test $(GO_BENCH_FLAGS) $$($(GO) list ./...)

.PHONY: check
ifneq (,$(findstring $(GOLANGCILINT_WANT_VERSION),$(GOLANGCILINT_VERSION)))
check:
	golangci-lint run --config=../.golangci.yml
else
check:
	# TODO: relative replace directives don't play well with docker mounts
	echo "Running CI checks via docker is not supported for submodules" && exit 1
endif
