#!/bin/bash
# Script to build the application

echo "Building the application..."
go build -o bin/blockchain-vm cmd/main.go

if [ $? -eq 0 ]; then
    echo "Build successful!"
else
    echo "Build failed!"
    exit 1
fi
