require("@nomiclabs/hardhat-waffle");
require("dotenv").config();

module.exports = {
  solidity: "0.8.19",
  networks: {
    sepolia: {
      url: "https://rpc.ankr.com/eth_sepolia",
      accounts: [process.env.BLOCKCHAIN_PRIVATE_KEY],
    },
    hardhat: {
      chainId: 1337,
    },
  },
};
