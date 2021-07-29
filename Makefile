GO := GOPRIVATE=github.com/isovalent CGO_ENABLED=0 go
INSTALL = $(QUIET)install
BINDIR ?= /usr/local/bin
TARGET := fake
TEST_TIMEOUT ?= 5s

GOLANGCILINT_WANT_VERSION = 1.41.1
GOLANGCILINT_VERSION = $(shell golangci-lint version 2>/dev/null)

.PHONY: all
all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET):
	$(GO) build -o $(TARGET) ./cmd/fake

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
	GOPRIVATE=github.com/isovalent CGO_ENABLED=1 go test -timeout=$(TEST_TIMEOUT) -race -cover $$($(GO) list ./...)

.PHONY: bench
bench:
	GOPRIVATE=github.com/isovalent CGO_ENABLED=1 go test -timeout=30s -bench=. $$($(GO) list ./...)

.PHONY: check
ifneq (,$(findstring $(GOLANGCILINT_WANT_VERSION),$(GOLANGCILINT_VERSION)))
check:
	golangci-lint run
else
check:
	docker run --rm -v `pwd`:/app -w /app docker.io/golangci/golangci-lint:v$(GOLANGCILINT_WANT_VERSION) golangci-lint run
endif
