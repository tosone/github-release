# github-release

### config
``` yaml
Username: tosone
Repo: test
Branch: master
Token: token
ClientID: ClientID
ClientSecret: ClientSecret
Rewrite: true
Draft: false
Prerelease: false
Runtime:
  Timeout: 10
  Debug: true
Release:
  Files:
    - release/*
    - debug/*
    - releasefile
  Compress: true
  CompressWith:
    - with/*
    - LICENSE
```
``` bash
release2github create
```
``` bash
release --help
```
``` bash
Release files and changelog to github release page.

Usage:
  release2github [command]

Available Commands:
  create      Create a new release on github release page.
  delete      Delete a tag release from github release page.
  help        Help about any command
  version     Get version

Flags:
  -h, --help   help for release2github

Use "release2github [command] --help" for more information about a command.
```