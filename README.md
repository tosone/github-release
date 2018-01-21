# release2github [![Build Status](https://travis-ci.org/tosone/release2github.svg?branch=v0.0.1)](https://travis-ci.org/tosone/release2github)

Release files and changelog to github release page.

### Usage

Release with a config file, default is `.release`. And execute path should be a git working repository.
Before push release files you should add a tag for you git repository, or you will fail to push any to github.

``` yaml
Username: tosone # Github username
Repo: release2github # repo name 
Branch: master # Build branch 
Token: token # Github token, generate a token here https://github.com/settings/tokens
ClientID: ClientID # New a OAuth app that can visit https://api.github.com more times. https://github.com/settings/developers
ClientSecret: ClientSecret # OAuth app client Secret.
Rewrite: true # Is rewrite the release or not. 
Draft: false # Is just a Draft or not.
Prerelease: false # Is prerelease or not.
Runtime:
  Timeout: 10 # Wait for visit https://api.github.com max timeout.
  Debug: true # Print the debug information.
Release:
  Files: # All of the files that will be upload release page.
    - release/*
    - debug/*
    - releasefile
  Compress: true # Is compress the upload files or not.
  CompressWith: # The files that will compress with upload file.
    - with/*
    - LICENSE
```

``` bash
release2github --help

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

### Env

- All of the config can be set in environment with prefix `RELEASE` and Separate with `_`.
- `Token`, `ClientID`, `ClientSecret` can be set as a variable in environment.
- `Token` can be set as `RELEASE_TOKEN`.
