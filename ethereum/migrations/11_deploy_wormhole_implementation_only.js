// MNEMONIC="" npm run deploy-wormhole-implementation-only -- --network acala

const Implementation = artifacts.require("Implementation");
module.exports = async function(deployer) {
  await deployer.deploy(Implementation);
};
