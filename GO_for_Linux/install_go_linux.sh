#!/bin/bash

# Check if Go is already installed
if command -v go &> /dev/null
then
    echo "Go is already installed."
    exit
fi

# Install Go on Ubuntu-based systems
if command -v apt-get &> /dev/null
then
    sudo apt-get update
    sudo apt-get install -y golang-go
fi

# Install Go on systems using dnf (Fedora, RHEL, CentOS)
if command -v dnf &> /dev/null
then
    sudo dnf install -y golang
fi

# Install Go on systems using yum (older versions of RHEL, CentOS)
if command -v yum &> /dev/null
then
    sudo yum install -y golang
fi

# Print Go version
go version
