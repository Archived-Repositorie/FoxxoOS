#!/usr/bin/env bash
#run script using command `bash`

nix-env -iA nixos.go
nix-env -iA electron

git pull

go build

cp $(which electron) ./foxxoos/vendor/electron-linux-amd64/electron

sudo ./FoxxoOS
