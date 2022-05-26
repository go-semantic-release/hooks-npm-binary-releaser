#!/usr/bin/env bash

set -euo pipefail

pluginDir=".semrel/$(go env GOOS)_$(go env GOARCH)/hooks-npm-binary-releaser/0.0.0-dev/"
[[ ! -d "$pluginDir" ]] && {
  echo "creating $pluginDir"
  mkdir -p "$pluginDir"
}

go build -o "$pluginDir/npm-binary-releaser" ./cmd/hooks-npm-binary-releaser
