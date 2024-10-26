// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MedicalRecords {
    // Struktur untuk rekam medis
    struct RecordHash {
        string noSEP;         
        string userNIK;       
        bytes32 dataHash;     
        uint256 timestamp;    
        bool exists;         
    }
    
    // Storage buat records: SEP number -> Record data
    mapping(string => RecordHash) private records;
    
    // Event that will be emitted when a new record is added
    event RecordAdded(
        string noSEP,
        string userNIK,
        uint256 timestamp
    );
    
    // Add rekam medis
    function addRecord(
        string memory noSEP,
        string memory userNIK,
        bytes32 dataHash
    ) public {
        // Validasi
        require(!records[noSEP].exists, "Record already exists");
        
        // Buat record
        records[noSEP] = RecordHash({
            noSEP: noSEP,
            userNIK: userNIK,
            dataHash: dataHash,
            timestamp: block.timestamp,
            exists: true
        });
        
        // Emit event
        emit RecordAdded(noSEP, userNIK, block.timestamp);
    }
    
    // Verif hash
    function verifyRecord(
        string memory noSEP,
        bytes32 dataHash
    ) public view returns (bool) {
        // Check if record exists
        require(records[noSEP].exists, "Record does not exist");
        
       
        return records[noSEP].dataHash == dataHash;
    }
}