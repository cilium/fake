GO := GOPRIVATE=github.com/isovalent CGO_ENABLED=0 go
INSTALL = $(QUIET)install
BINDIR ?= /usr/local/bin
TARGET := fake
TEST_TIMEOUT ?= 5s

all: $(TARGET)

$(TARGET):
	$(GO) build -o $(TARGET) ./cmd/fake

install: $(TARGET)
	$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(INSTALL) -m 0755 ./$(TARGET) $(DESTDIR)$(BINDIR)

clean:
	rm -f $(TARGET)
	rm -f cpu.prof mem.prof

test:
	GOPRIVATE=github.com/isovalent CGO_ENABLED=1 go test -timeout=$(TEST_TIMEOUT) -race -cover $$($(GO) list ./...)

bench:
	GOPRIVATE=github.com/isovalent CGO_ENABLED=1 go test -timeout=30s -bench=. $$($(GO) list ./...)

vet:
	$(GO) vet $$($(GO) list ./...)

.PHONY: all $(TARGET) install clean test bench vet
