#!/bin/bash

msg() {
	echo 2>&1 "$0": $@
}

build() {
	local pkg="$1"

	gofmt -s -w "$pkg"
	go fix "$pkg"
	go vet -vettool="$(which shadow)" "$pkg"

	#hash gosimple >/dev/null && gosimple "$pkg"
	hash golint >/dev/null && golint "$pkg"
	#hash staticcheck >/dev/null && staticcheck "$pkg"

	go test "$pkg"
	go install -v "$pkg"
}

build ./suco-d
