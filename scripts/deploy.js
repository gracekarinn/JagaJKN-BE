const main = async () => {
  const MedicalRecords = await hre.ethers.getContractFactory("MedicalRecords");
  const medicalRecords = await MedicalRecords.deploy();
  await medicalRecords.deployed();
  console.log("Contract deployed to:", medicalRecords.address);
};

main().catch((error) => {
  console.error(error);
  process.exit(1);
});
