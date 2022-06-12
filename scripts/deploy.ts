import fs from "fs"
import { ethers } from "hardhat"
import os from "os"

const AUCTION_HOUSE_DURATION_SECONDS = 604800
const AUCTION_HOUSE_MIN_BID_INCREMENT_PERCENT = 5
const AUCTION_HOUSE_DAO_CUT = 10

const main = async () => {
    let randomWal = ethers.Wallet.createRandom()
    let randomWallet = new ethers.Wallet(randomWal.privateKey, ethers.provider)
    let signers = await ethers.getSigners()
    const tx = await signers[0].sendTransaction({
        to: randomWallet.address,
        value: ethers.utils.parseEther("500.0")
    })
    await tx.wait()
    setEnvValue("PROXY_PRIVATE_KEY", randomWallet.privateKey.slice(2))

    const WETH = await ethers.getContractFactory("WETH")
    const weth = await WETH.deploy()
    await weth.deployed()
    const Token = await ethers.getContractFactory("Token")
    const token = await Token.deploy()
    await token.deployed()
    const Distributor = await ethers.getContractFactory("Distributor")
    const distributor = await Distributor.deploy(weth.address)
    await distributor.deployed()
    const DAO = await ethers.getContractFactory("DAO", randomWallet)
    const dao = await DAO.deploy(token.address, weth.address)
    await dao.deployed()
    const AuctionHouse = await ethers.getContractFactory("AuctionHouse")
    const auctionHouse = await AuctionHouse.deploy(token.address, dao.address, distributor.address, weth.address, AUCTION_HOUSE_DURATION_SECONDS, AUCTION_HOUSE_MIN_BID_INCREMENT_PERCENT, AUCTION_HOUSE_DAO_CUT)
    await auctionHouse.deployed()
    const setReservePriceTx = await auctionHouse.setReservePrice(1)
    await setReservePriceTx.wait()
    const setNewOwnerTx = await auctionHouse.transferOwnership(randomWallet.address)
    await setNewOwnerTx.wait()
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