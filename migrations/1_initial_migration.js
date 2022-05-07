const Trixels = artifacts.require("Trixels");
const TrixelsToken = artifacts.require("TrixelsToken");
const TrixelsAuctionHouse = artifacts.require("TrixelsAuctionHouse");

module.exports = async (deployer) => {
  await deployer.deploy(Trixels);
  // trixelsInstance = await Trixels.deployed();
  // await trixelsInstance.initialize();
  await deployer.deploy(TrixelsToken, "0x31b653c95756Aac850eA1e096f844e4478Aa8E97");
  const trixelToken = await TrixelsToken.deployed();
  await deployer.deploy(TrixelsAuctionHouse);
  const trixelAuctionHouse = await TrixelsAuctionHouse.deployed();
  await trixelAuctionHouse.initialize(trixelToken.address, 1000, 5, 10, "0x36811fA4e42C483F73F52D32960354Be023267dc");
};
