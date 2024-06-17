#!/bin/bash
# Script to run all tests

echo "Running tests..."
go test ./...

if [ $? -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Some tests failed!"
    exit 1
fi
