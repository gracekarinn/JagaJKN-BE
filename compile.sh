#!/bin/bash

echo "Cleaning build directory..."
rm -rf build
mkdir -p build

echo "Compiling contract..."
solc --abi --bin contracts/MedicalRecords.sol -o build --overwrite

echo "Generating Go bindings..."
abigen --bin=build/MedicalRecords.bin --abi=build/MedicalRecords.abi --pkg=contracts --out=internal/blockchain/contracts/medical_records.go

echo "Compilation complete!"
