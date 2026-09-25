#! /usr/bin/env nix-shell
#! nix-shell -i bash -p bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely
