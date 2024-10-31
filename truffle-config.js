require("dotenv").config();
const HDWalletProvider = require("@truffle/hdwallet-provider");

const provider = () => {
  if (!process.env.BLOCKCHAIN_PRIVATE_KEY || !process.env.BLOCKCHAIN_PROVIDER) {
    console.error("Environment variables not found:");
    console.error(
      "BLOCKCHAIN_PRIVATE_KEY:",
      !!process.env.BLOCKCHAIN_PRIVATE_KEY
    );
    console.error("BLOCKCHAIN_PROVIDER:", !!process.env.BLOCKCHAIN_PROVIDER);
    throw new Error("Missing environment variables");
  }

  return new HDWalletProvider({
    privateKeys: [process.env.BLOCKCHAIN_PRIVATE_KEY],
    providerOrUrl: process.env.BLOCKCHAIN_PROVIDER,
  });
};

module.exports = {
  networks: {
    sepolia: {
      provider: provider,
      network_id: 11155111,
      gas: 3000000,
      gasPrice: 20000000000,
      confirmations: 2,
      timeoutBlocks: 200,
      skipDryRun: true,
    },
  },
  compilers: {
    solc: {
      version: "0.8.19",
      settings: {
        optimizer: {
          enabled: true,
          runs: 200,
        },
      },
    },
  },
};
