import { ethers } from "hardhat"

import { WETH, Token, Distributor, AuctionHouse } from "../typechain-types"

const AUCTION_HOUSE_DURATION_SECONDS = 604800
const AUCTION_HOUSE_MIN_BID_INCREMENT_PERCENT = 5
const AUCTION_HOUSE_DAO_CUT = 10

const main = async () => {
    const WETH = await ethers.getContractFactory("WETH")
    const weth = await WETH.deploy() as WETH
    const Token = await ethers.getContractFactory("Token")
    const token = await Token.deploy() as Token
    const Distributor = await ethers.getContractFactory("Distributor")
    const distributor = await Distributor.deploy(weth.address) as Distributor
    const DAO = await ethers.getContractFactory("DAO")
    const dao = await DAO.deploy(token.address, weth.address)
    const AuctionHouse = await ethers.getContractFactory("AuctionHouse")
    const auctionHouse = await AuctionHouse.deploy(token.address, dao.address, distributor.address, weth.address, AUCTION_HOUSE_DURATION_SECONDS, AUCTION_HOUSE_MIN_BID_INCREMENT_PERCENT, AUCTION_HOUSE_DAO_CUT) as AuctionHouse
    const setAuctionHouseTx = await token.setAuctionHouse(auctionHouse.address)
    await setAuctionHouseTx.wait()
}

main()
    .then(() => process.exit(0))
    .catch(err => {
        console.error(err)
        process.exit(1)
    })