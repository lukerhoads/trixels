const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Auction House", () => {
    const trixelId = "VADq5G7mozFHJgx9k7tv2Q4FQzf66CSTXWQe8lhDMGVigw"
    const minter = "0x03d21F78f63BB844438f8F87cBFcF68f984eee19"
    const weth = "0xd2e83e87D0f663117451fBf2838B7685e175e669"

    const setupAuction = async () => {
        const Token = await ethers.getContractFactory("TrixelsToken")
        const token = await Token.deploy("0x03d21F78f63BB844438f8F87cBFcF68f984eee19") // minter address
        const AuctionHouse = await ethers.getContractFactory("TrixelsAuctionHouse")
        const auctionHouse = await AuctionHouse.deploy()
        await auctionHouse.deployed()
        const initTx = auctionHouse.initialize(token, 100, 5, 5, "0xd2e83e87D0f663117451fBf2838B7685e175e669") // weth address
        await initTx.wait()
        const startAuctionTx = auctionHouse.startAuction(trixelId)
        await startAuctionTx.wait()
        return auctionHouse
    }

    it("Should start an auction correctly", async () => {
        const auctionHouse = await setupAuction()
        const auction = await auctionHouse.auction()
        expect(auction.trixelId).to.equal(trixelId)
        expect(auction.amount).to.equal(5)
        expect(auction.bidder).to.equal("0x0000000000000000000000000000000000000000")
        expect(auction.settled).to.equal(false)
    })

    it("Should settle an auction correctly", async () => {
        const auctionHouse = await setupAuction()
        const settleAuctionTx = auctionHouse.settleAuction()
        await settleAuctionTx.wait()
        const auction = await auctionHouse.auction()
        expect(auction.settled).to.equal(true)
    })

    it("Should create a bid correctly", async () => {
        const auctionHouse = await setupAuction()
        const createBidTx = auctionHouse.createBid(trixelId)
        await createBidTx.wait()
        const auction = await auctionHouse.auction()
        expect(auction.amount).to.equal(5)
        expect(auction.bidder).to.equal("0x0000000000000000000000000000000000000000")
        expect(auction.settled).to.equal(false)
    })
})