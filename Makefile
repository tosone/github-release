BuildStamp = $(shell date '+%Y-%m-%d_%H:%M:%S')
GitHash    = $(shell git rev-parse HEAD)
Version    = $(shell git describe --abbrev=0 --tags --always)
Target     = $(shell basename $(abspath $(dir $$PWD)))
Suffix     =

ifeq ($(OS),Windows_NT)
	OSName = windows
	Suffix = .exe
else
	OSName = $(shell echo $(shell uname -s) | tr '[:upper:]' '[:lower:]')
endif

${OSName}: clean
	GOOS=$@ go build -v -o ${Target}-$@${Suffix} -ldflags "-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

release: clean
	xgo -v -out ${Target}-${Version} --targets=windows/*,darwin/*,linux/* -ldflags "-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}" github.com/EffDataAly/GithubTraveler
	mkdir release
	mv ${Target}-${Version}-* release

authors:
	printf "Authors\n=======\n\nProject's contributors:\n\n" > AUTHORS.md
	git log --raw | grep "^Author: " | cut -d ' ' -f2- | cut -d '<' -f1 | sed 's/^/- /' | sort | uniq >> AUTHORS.md

lint:
	gometalinter.v2 ./...

clean:
	$(RM) -r release

.PHONY: release all
