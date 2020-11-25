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
  ext=""
  [[ $os == "windows" ]] && ext=".exe"

  export GOOS="$os"
  export GOARCH="$arch"

  go build -o "./${os}/shortest-path$ext"
done

cd "$tmp" || exit 1
