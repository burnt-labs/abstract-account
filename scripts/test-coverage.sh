#!/bin/bash

# Script to run tests with coverage, excluding generated protobuf files

echo "Running tests with coverage (excluding .pb.go files)..."

# Run tests with coverage
go test ./x/abstractaccount/... -coverprofile=coverage.out

# Filter out .pb.go and .pb.gw.go files from coverage report
grep -v "\.pb\.go:" coverage.out | grep -v "\.pb\.gw\.go:" > coverage_filtered.out

# Show coverage report without .pb.go files
echo "Coverage report (excluding generated files):"
go tool cover -func=coverage_filtered.out

# Generate HTML report without .pb.go files
go tool cover -html=coverage_filtered.out -o coverage.html

echo "HTML coverage report generated: coverage.html"
echo "Filtered coverage file: coverage_filtered.out"
