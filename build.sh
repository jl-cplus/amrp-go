#!/bin/bash

# remove build cache
go clean -cache

# module dependencies
go mod tidy -C src

# cli build tool
go install fyne.io/tools/cmd/fyne@latest

# reset build environment
rm -rf dist && mkdir dist

# change working directory
cd dist

# create executable binary package
fyne package --src ../src --os linux --release

# generate SHA-256 checksum
sha256sum amrp-go.tar.xz > amrp-go.tar.xz.sha256sum