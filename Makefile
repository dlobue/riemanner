SHELL := /bin/bash
PKG = github.com/Clever/riemanner
SUBPKGS = $(addprefix $(PKG)/,riemanner)
PKGS = $(PKG) $(SUBPKGS)
.PHONY: test $(PKGS)

test: $(PKGS)

$(PKGS):
ifeq ($(LINT),1)
	golint $(GOPATH)/src/$@*/**.go
endif
	go get -d -t $@
ifeq ($(COVERAGE),1)
	go test -cover -coverprofile=$(GOPATH)/src/$@/c.out $@ -test.v
	go tool cover -html=$(GOPATH)/src/$@/c.out
else
	go test $@ -test.v
endif
	go install -v $@

install:
	sudo cp $(GOPATH)/bin/riemanner /usr/local/bin/riemanner
