#!/bin/bash

# Check if protoc is already installed
if ! command -v protoc &> /dev/null; then
    echo "Installing protoc compiler..."
    
# Install protoc compiler
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip
unzip protoc-30.2-linux-x86_64.zip -d $HOME/.local

# Add protoc to PATH
echo 'export PATH="$PATH:$HOME/.local/bin"' >> ~/.bashrc
source ~/.bashrc
    
    echo "Protoc installation completed!"
else
    echo "Protoc is already installed. Skipping installation."
fi
