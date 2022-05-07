const Trixels = artifacts.require("Trixels");
const TrixelsNFT = artifacts.require("TrixelsNFT");
const TrixelsAuctionHouse = artifacts.require("TrixelsAuctionHouse");

module.exports = function (deployer) {
  deployer.deploy(Trixels);
  deployer.deploy(TrixelsNFT);
  deployer.deploy(TrixelsAuctionHouse);
};
