#!/bin/bash

# remove build cache
go clean -cache

# module dependencies
go mod tidy -C src

# cli build tool
go install fyne.io/tools/cmd/fyne@latest

# reset build environment
rm -rf dist && mkdir dist

# create executable binary package
fyne package --src src --os windows --release

# change working directory
cd dist

# move the built executable binary to the dist folder
# workaround for fyne --exe issue preventing PE file icon rendering
mv ../src/amrp-go.exe .

# use compression to reduce executable size
upx -9 amrp-go.exe 

# generate SHA-256 checksum
sha256sum amrp-go.exe > amrp-go.exe.sha256sum