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

check: gofmt ineffassign lint staticcheck vet

gofmt:
	@source="$$(find . -type f -name '*.go' -not -path './vendor/*')"; \
	unformatted="$$(gofmt -l $$source)"; \
	if [ -n "$$unformatted" ]; then echo "unformatted source code:" && echo "$$unformatted" && exit 1; fi

ineffassign:
	$(GO) run ./vendor/github.com/gordonklaus/ineffassign .

lint:
	$(GO) run ./vendor/golang.org/x/lint/golint -set_exit_status $$($(GO) list ./...)

staticcheck:
	$(GO) run ./vendor/honnef.co/go/tools/cmd/staticcheck -checks="all,-ST1000" $$($(GO) list ./...)

vet:
	$(GO) vet $$($(GO) list ./...)

.PHONY: all $(TARGET) install clean test bench check gofmt ineffassign lint staticcheck vet
