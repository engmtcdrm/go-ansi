#!/bin/bash

# Install whatever version of Go is specified in go.mod
go_version=$(grep -E '^go' go.mod | awk '{print $2}')
go_version_file="go${go_version}.linux-amd64.tar.gz"

sudo rm -rf "/usr/local/go"
wget "https://go.dev/dl/${go_version_file}"
sudo tar -C "/usr/local" -xzf "${go_version_file}"
rm "${go_version_file}"
go version
