#!/usr/bin/env bin

nix-env -iA nixos.go

go build

firefox --new-tab localhost:8080 & sudo ./FoxxoOS