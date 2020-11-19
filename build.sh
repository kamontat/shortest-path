#!/bin/bash

export BUILDERS=(
  "darwin/amd64"
  "linux/amd64"
  "windows/amd64"
)

tmp="$PWD"

cd "app" || exit 1

for builder in "${BUILDERS[@]}"; do
  os="${builder%%/*}"
  arch="${builder##*/}"

  export GOOS="$os"
  export GOARCH="$arch"

  go build -o "shortest-path-${os}-${arch}"
done

cd "$tmp" || exit 1
