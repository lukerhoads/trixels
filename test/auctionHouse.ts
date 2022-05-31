import { ethers, waffle } from "hardhat";
import { Signer } from "ethers";
import { expect } from "chai";

import { AuctionHouse, Token, DAO, WETH, Distributor } from "../typechain-types";
import TokenArtifact from '../artifacts/contracts/Token.sol/Token.json'
import WETHArtifact from '../artifacts/contracts/WETH.sol/WETH.json'

const { deployContract } = waffle

describe("Auction House", () => {
    let accounts: Signer[]
    let token: Token
    let weth: WETH
    let dao: DAO
    let distributor: Distributor
    let auctionHouse: AuctionHouse

    beforeEach(async () => {
        accounts = await ethers.getSigners()
        weth = (await deployContract(accounts[0], WETHArtifact)) as WETH 
        token = (await deployContract(accounts[0], TokenArtifact)) as Token 
        const distributorContract = await ethers.getContractFactory("Distributor")
        distributor = await distributorContract.deploy(weth.address)
        const daoContract = await ethers.getContractFactory("DAO")
        dao = await daoContract.deploy(token.address, weth.address)
        let auctionHouseContract = await ethers.getContractFactory("AuctionHouse")
        auctionHouse = await auctionHouseContract.deploy(token.address, dao.address, distributor.address, weth.address, 604800, 5, 10)
        let setAuctionHouseTx = await token.setAuctionHouse(auctionHouse.address)
        await setAuctionHouseTx.wait()
    })

    describe("startAuction", async () => {
        it("should start an auction and mint a new token", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            let auctionHouseTokenBalance = await token.balanceOf(auctionHouse.address)
            expect(auctionHouseTokenBalance).to.eq(1)
            let auction = await auctionHouse.auction()
            expect(auction.highestBid).to.eq(0)
            expect(auction.highestBidder).to.eq(ethers.constants.AddressZero)
            expect(auction.tokenId).to.eq(1)
            expect(auction.settled).to.eq(false)
        })
    })

    describe("placeBid", async () => {
        it("correctly places a bid", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            let placeBidTx = await auctionHouse.placeBid(0, { value: ethers.utils.parseEther("1.0") })
            await placeBidTx.wait()
            let auction = await auctionHouse.auction()
            expect(auction.highestBid.toString()).to.eq("1000000000000000000")
        })

        it("reverts when bidder is already highest bidder", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            let placeBidTx = await auctionHouse.placeBid(0, { value: ethers.utils.parseEther("1.0") })
            await placeBidTx.wait()
            expect(await auctionHouse.placeBid(0, { value: ethers.utils.parseEther("2.0") })).to.be.reverted
        })

        it("reverse when auction has already expired", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            await waffle.provider.send("evm_increaseTime", [1209600])
            await waffle.provider.send("evm_mine", [])
            expect(await auctionHouse.placeBid(0, { value: ethers.utils.parseEther("1.0") })).to.be.reverted
        })

        it("reverts when auction is already settled", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            let endAuctionTx = await auctionHouse.endAuction()
            await endAuctionTx.wait()
            expect(await auctionHouse.placeBid(0, { value: ethers.utils.parseEther("1.0") })).to.be.reverted
        })  
    })

    describe("endAuction", async () => {
        it("should end an auction, burn the token when there are no bids", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            await waffle.provider.send("evm_increaseTime", [604800])
            await waffle.provider.send("evm_mine", [])
            let endAuctionTx = await auctionHouse.endAuction()
            await endAuctionTx.wait()
            let auction = await auctionHouse.auction()
            expect(auction.settled).to.be.true
            let auctionHouseTokenBalance = await token.balanceOf(auctionHouse.address)
            expect(auctionHouseTokenBalance).to.eq(0)
        })
       
        it("should end an auction, distribute the token, and send the balance to the correct parties when there are bids", async () => {
            let startAuctionTx = await auctionHouse.startAuction()
            await startAuctionTx.wait()
            let auction = await auctionHouse.auction()
            let placeBidTx = await auctionHouse.placeBid(auction.tokenId, { value: ethers.utils.parseEther("1.0") })
            await placeBidTx.wait()
            await waffle.provider.send("evm_increaseTime", [604800])
            await waffle.provider.send("evm_mine", [])
            let endAuctionTx = await auctionHouse.endAuction()
            await endAuctionTx.wait()
            auction = await auctionHouse.auction()
            expect(auction.settled).to.be.true
            let highestBidderTokenBalance = await token.balanceOf(await accounts[0].getAddress())
            expect(highestBidderTokenBalance).to.eq(1)
            let daoBalance = await waffle.provider.getBalance(dao.address)
            expect(daoBalance.toString()).to.eq("100000000000000000")
            let distributorBalance = await waffle.provider.getBalance(distributor.address)
            expect(distributorBalance.toString()).to.eq("900000000000000000")
        })
    })
})