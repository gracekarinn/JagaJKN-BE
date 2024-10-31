const MedicalRecords = artifacts.require("MedicalRecords");

module.exports = function (deployer) {
  deployer.deploy(MedicalRecords).then(() => {
    console.log("Contract deployed to:", MedicalRecords.address);
  });
};
