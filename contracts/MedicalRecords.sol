// Made by Grace
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MedicalRecords {
    // Structures
    struct UserInfo {
        bytes32 dataHash;
        uint256 timestamp;
        bool exists;
    }

    struct RecordHash {
        string noSEP;
        string userNIK;
        bytes32 dataHash;
        uint256 timestamp;
        bool exists;
    }

    // Mappings
    mapping(string => UserInfo) private users;        // NIK -> UserInfo
    mapping(string => RecordHash) private records;    // SEP -> RecordHash

    // Events
    event UserRegistered(string nik, uint256 timestamp);
    event RecordAdded(string noSEP, string userNIK, uint256 timestamp);

    // User registration
    function addUser(string memory nik, bytes32 userHash) public {
        require(!users[nik].exists, "User already registered");

        users[nik] = UserInfo({
            dataHash: userHash,
            timestamp: block.timestamp,
            exists: true
        });

        emit UserRegistered(nik, block.timestamp);
    }

    // User verification
    function verifyUser(string memory nik, bytes32 userHash) public view returns (bool) {
        require(users[nik].exists, "User not registered");
        return users[nik].dataHash == userHash;
    }

    // Medical record functions
    function addRecord(string memory noSEP, string memory userNIK, bytes32 dataHash) public {
        require(!records[noSEP].exists, "Record already exists");
        require(users[userNIK].exists, "User not registered");

        records[noSEP] = RecordHash({
            noSEP: noSEP,
            userNIK: userNIK,
            dataHash: dataHash,
            timestamp: block.timestamp,
            exists: true
        });

        emit RecordAdded(noSEP, userNIK, block.timestamp);
    }

    // Record verification
    function verifyRecord(string memory noSEP, bytes32 dataHash) public view returns (bool) {
        require(records[noSEP].exists, "Record does not exist");
        return records[noSEP].dataHash == dataHash;
    }

    // Helper function to check if user exists
    function isUserRegistered(string memory nik) public view returns (bool) {
        return users[nik].exists;
    }

}