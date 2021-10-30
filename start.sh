#!/usr/bin/env bin

nix-env -iA nixos.go

git pull

go build

firefox --new-tab localhost:8080 & sudo ./FoxxoOS