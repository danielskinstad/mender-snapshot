DESTDIR ?= /
prefix ?= $(DESTDIR)
bindir=/usr/bin

GO ?= go
V ?=
PKGS = $(shell go list ./...)
PKGFILES = $(shell find . \( -path ./vendor \) -prune \
		-o -type f -name '*.go' -print)
PKGFILES_notest = $(shell echo $(PKGFILES) | tr ' ' '\n' | grep -v '_test.go' )

VERSION = $(shell git describe --tags --dirty --exact-match 2>/dev/null || git rev-parse --short HEAD)

GO_LDFLAGS = \
	-ldflags "-X github.com/mendersoftware/mender-snapshot/config.Version=$(VERSION)"

ifeq ($(V),1)
BUILDV = -v
endif

build: mender-snapshot

clean:
	@$(GO) clean

mender-snapshot: $(PKGFILES)
	@$(GO) build $(GO_LDFLAGS) $(BUILDV)

install: install-bin

install-bin: mender-snapshot
	@install -m 755 -d $(prefix)$(bindir)
	@install -m 755 mender-snapshot $(prefix)$(bindir)/

uninstall: uninstall-bin

uninstall-bin:
	@rm -f $(prefix)$(bindir)/mender-snapshot
	@-rmdir -p $(prefix)$(bindir)

check: test

test:
	@$(GO) test $(BUILDV) $(PKGS)

.PHONY: build
.PHONY: clean
.PHONY: install
.PHONY: install-bin
.PHONY: uninstall
.PHONY: uninstall-bin
.PHONY: test
.PHONY: check
