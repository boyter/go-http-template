#!/bin/bash

set -e

if [ -t 1 ]
then
  YELLOW='\033[0;33m'
  GREEN='\033[0;32m'
  RED='\033[0;31m'
  NC='\033[0m'
fi

yellow() { printf "${YELLOW}%s${NC}" "$*"; }
green() { printf "${GREEN}%s${NC}" "$*"; }
red() { printf "${RED}%s${NC}" "$*"; }

good() {
  echo "$(green "● success:")" "$@"
}

bad() {
  ret=$1
  shift
  echo "$(red "● failed:")" "$@"
  exit "$ret"
}

try() {
  "$@" || bad $? "$@" && good "$@"
}

cmd_exists() {
  type "$1" >/dev/null 2>&1
}

cmd_exists go || bad 1 "uhhh, where's your go install? what are you even doing here!"
cmd_exists gofmt || bad 1 "so you have go installed, but not gofmt? wtf mate^^"
cmd_exists golangci-lint || bad 1 "cannot find golangci-lint; check PATH or go to https://github.com/golangci/golangci-lint"
cmd_exists gitleaks || bad 1 "cannot find gitleaks 8.8.8 https://github.com/zricethezav/gitleaks"

try ./unique_code.py
try gitleaks detect -v -c gitleaks.toml
try gitleaks protect -v -c gitleaks.toml
try go build ./...
try go test -count=1 --tags=integration ./...
{
 {
   opt='shopt -s extglob nullglob'
   gofmt='gofmt -s -w -l !(vendor)/ *.go'
   notice="    running: ( $opt; $gofmt; )"
   prefix="    $(yellow modified)"
   trap 'echo "$notice"; $opt; $gofmt | sed -e "s#^#$prefix #g"' EXIT
 }

 # comma separate linters (e.g. "gofmt,stylecheck")
 additional_linters="gofmt"
 try golangci-lint run --enable $additional_linters ./...
 trap '' EXIT
}
