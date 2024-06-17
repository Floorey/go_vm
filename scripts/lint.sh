#!/bin/bash
# Script to run linter

echo "Running linter..."
golangci-lint run

if [ $? -eq 0 ]; then
    echo "Linting passed!"
else
    echo "Linting failed!"
    exit 1
fi
