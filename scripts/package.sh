#!/usr/bin/env bash
set -euo pipefail
APP=ctfmini
VERSION=${VERSION:-0.1.0}
OUT=dist
mkdir -p "$OUT"

# Linux amd64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$OUT/$APP-linux-amd64"
( cd "$OUT" && zip -9 "$APP-$VERSION-linux-amd64.zip" "$APP-linux-amd64" )

# macOS arm64
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o "$OUT/$APP-macos-arm64"
( cd "$OUT" && zip -9 "$APP-$VERSION-macos-arm64.zip" "$APP-macos-arm64" )

# macOS amd64
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "$OUT/$APP-macos-amd64"
( cd "$OUT" && zip -9 "$APP-$VERSION-macos-amd64.zip" "$APP-macos-amd64" )

# Windows amd64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$OUT/$APP-windows-amd64.exe"
( cd "$OUT" && zip -9 "$APP-$VERSION-windows-amd64.zip" "$APP-windows-amd64.exe" )

echo "Built zips in $OUT/"
