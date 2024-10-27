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
    
    // Struktur untuk user verification
    struct UserRegistration {
        bytes32 userHash;
        uint256 timestamp;
        bool exists;
    }
    
    // Storage buat records: SEP number -> Record data
    mapping(string => RecordHash) private records;
    
    // Storage buat user verification: NIK -> User data
    mapping(string => UserRegistration) private users;
    
    // Events
    event RecordAdded(
        string noSEP,
        string userNIK,
        uint256 timestamp
    );

    event UserRegistered(
        string nik,
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
        require(users[userNIK].exists, "User not registered"); 
        
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

    // Register new user
    function addUser(
        string memory nik,
        bytes32 userHash
    ) public {
        require(!users[nik].exists, "User already registered");
        
        users[nik] = UserRegistration({
            userHash: userHash,
            timestamp: block.timestamp,
            exists: true
        });
        
        emit UserRegistered(nik, block.timestamp);
    }

    // Verify user
    function verifyUser(
        string memory nik,
        bytes32 userHash
    ) public view returns (bool) {
        require(users[nik].exists, "User not registered");
        
        return users[nik].userHash == userHash;
    }

    // Optional: Get user registration timestamp
    function getUserRegistrationTime(
        string memory nik
    ) public view returns (uint256) {
        require(users[nik].exists, "User not registered");
        
        return users[nik].timestamp;
    }
}