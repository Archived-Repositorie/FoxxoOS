#!/usr/bin/env bin

nix-env -iA go

go build

firefox --new-tab localhost:8080 & ./FoxxoOS