// MNEMONIC="" npm run deploy-wormhole-implementation-only -- --network acala

const BridgeImplementation = artifacts.require("NFTBridgeImplementation");
module.exports = async function(deployer, network) {
  if (network === "test") return;
  await deployer.deploy(BridgeImplementation);
};
