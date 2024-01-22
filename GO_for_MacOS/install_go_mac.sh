#!/bin/bash

# Check if Go is already installed
if command -v go &> /dev/null
then
    echo "Go is already installed."
    exit
fi

# Install Go using Homebrew if available
if command -v brew &> /dev/null
then
    brew install go
    exit
fi

# Install Go using MacPorts if available
if command -v port &> /dev/null
then
    sudo port install go
    exit
fi

# Download and install Go from the official website
echo "Downloading and installing Go..."
url="https://golang.org/dl/"
latest_version=$(curl -sSL $url | grep -o -E 'go[0-9]+\.[0-9]+(\.[0-9]+)?' | sort -V | tail -n 1)
download_url="${url}${latest_version}.darwin-amd64.tar.gz"
curl -sSL $download_url | sudo tar -C /usr/local -xz

# Add Go binary directory to PATH
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bash_profile
source ~/.bash_profile

# Print Go version
go version

