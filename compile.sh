#!/bin/bash

# Buat cek
rm -rf build
mkdir build

# Compile contract
solc --abi --bin contracts/MedicalRecords.sol -o build

echo "Contract compiled successfully!"