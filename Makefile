BuildStamp = $(shell date '+%Y-%m-%d_%H:%M:%S')
GitHash    = $(shell git rev-parse HEAD)
Version    = $(shell git describe --abbrev=0 --tags --always)
Target     = $(shell basename $(abspath $(dir $$PWD)))

.PHONY: build
build: clean
	go build -v -o release/${Target} -ldflags \
	"-X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: release
release:
	go build -v -o release/${Target} -ldflags \
	"-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: authors
authors:
	printf "Authors\n=======\n\nProject's contributors:\n\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: clean
clean:
	$(RM) -r release
