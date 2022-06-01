import fs from "fs"
import os from "os"
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
    setEnvValue("AUCTION_HOUSE_ADDRESS", auctionHouse.address)
}

const setEnvValue = (key: string, value: string) => {
    const ENV_VARS = fs.readFileSync(".env", "utf8").split(os.EOL)
    const target = ENV_VARS.indexOf(ENV_VARS.find(line => {
        const keyValRegex = new RegExp(`(?<!#\\s*)${key}(?==)`)
        return line.match(keyValRegex)
    }) || "")

    if (target !== -1) {
        ENV_VARS.splice(target, 1, `${key}=${value}`)
    } else {
        ENV_VARS.push(`${key}=${value}`)
    }

    fs.writeFileSync(".env", ENV_VARS.join(os.EOL))
}

main()
    .then(() => process.exit(0))
    .catch(err => {
        console.error(err)
        process.exit(1)
    })