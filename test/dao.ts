import { ethers, waffle } from "hardhat";
import { BigNumber, Signer } from "ethers";
import { expect } from "chai";
import { solidityKeccak256 } from "ethers/lib/utils";

import { DAO } from '../typechain-types/contracts/DAO'
import { IDAO } from "../typechain-types/contracts/interfaces/IDAO";
import { Token } from '../typechain-types/contracts/Token'
import TokenArtifact from '../artifacts/contracts/Token.sol/Token.json'
import { WETH } from '../typechain-types/contracts/WETH'
import WETHArtifact from '../artifacts/contracts/WETH.sol/WETH.json'

const { deployContract } = waffle

describe("DAO", () => {
    let accounts: Signer[]
    let weth: WETH
    let token: Token
    let dao: DAO

    beforeEach(async () => {
        accounts = await ethers.getSigners()
        weth = (await deployContract(accounts[0], WETHArtifact)) as WETH 
        token = (await deployContract(accounts[0], TokenArtifact)) as Token 
        const daoContract = await ethers.getContractFactory("DAO")
        dao = await daoContract.deploy(token.address, weth.address)
    })

    describe("isMember", async () => {
        it("should return true when user is a token holder", async () => {
            let accountAddr = await accounts[1].getAddress()
            let mintTx = await token.mint(accountAddr)
            await mintTx.wait()
            let isMember = await dao.isMember(accountAddr)
            expect(isMember).to.be.true
        })

        it("should return false when user does not hold a token", async () => {
            let accountAddr = await accounts[1].getAddress()
            let isMember = await dao.isMember(accountAddr)
            expect(isMember).to.be.false
        })
    })

    type Proposal = {
        recipient: string,
        amount: string,
        description: string,
        transactionData: any[]
    }

    const mintTo = async (account: Signer) => {
        const mintTx = await token.mint(await account.getAddress())
        await mintTx.wait()
    }

    const makeProposals = async () => {
        let proposals: Proposal[] = [
            {
                recipient: await accounts[2].getAddress(),
                amount: "1.0",
                description: "send to accounts 2",
                transactionData: []
            },
            {
                recipient: await accounts[3].getAddress(),
                amount: "2.0",
                description: "send to accounts 3",
                transactionData: []
            }
        ]

        proposals.forEach(async proposal => {
            let proposalTx = await dao.makeProposal(proposal.recipient, ethers.utils.parseEther(proposal.amount), proposal.description, proposal.transactionData)
            await proposalTx.wait()
        })
    }

    describe("numProposals", async () => {
        it("should return zero when there are no proposals", async () => {
            let proposalCount = await dao.numProposals()
            expect(proposalCount).to.eq(0)
        })

        it("should return the amount of proposals", async () => {
            await mintTo(accounts[0])
            await makeProposals()
            let proposalCount = await dao.numProposals()
            expect(proposalCount).to.eq("2")
        })
    })

    describe("vote", async () => {
        it("should count a member's vote", async () => {
            await mintTo(accounts[0])
            await makeProposals()
            let voteTx = await dao.vote(1, true)
            await voteTx.wait()
            let proposal = await dao.proposals(1)
            expect(proposal.yay).to.eq(1)
        })

        it("should count a member with more tokens as having more votes", async () => {
            await mintTo(accounts[0])
            await mintTo(accounts[0])
            await makeProposals()
            let voteTx = await dao.vote(1, true)
            await voteTx.wait()
            let proposal = await dao.proposals(1)
            expect(proposal.yay).to.eq(2)
        })

        it("should not allow a non-token holder to vote", async () => {
            await expect(dao.vote(1, true)).to.be.reverted
        })
    })

    describe("unVote", async () => {
        it("should unVote a member", async () => {
            await mintTo(accounts[0])
            await makeProposals()
            let voteTx = await dao.vote(1, true)
            await voteTx.wait()
            let unVoteTx = await dao.unVote(1)
            await unVoteTx.wait()
            let proposal = await dao.proposals(1)
            expect(proposal.yay).to.eq(0)
        })
    })

    describe("executeProposal", async () => {
        it("should execute a proposal", async () => {
            await accounts[0].sendTransaction({
                to: dao.address,
                value: ethers.utils.parseEther("1.0"),
            })
            await mintTo(accounts[0])
            await makeProposals()
            let voteTx = await dao.vote(0, true)
            await voteTx.wait()
            await waffle.provider.send("evm_increaseTime", [1209600])
            await waffle.provider.send("evm_mine", [])
            let executeProposalTx = await dao.executeProposal(0, [])
            await executeProposalTx.wait()
            expect((await accounts[2].getBalance()).toString()).to.eq("10001000000000000000000")
        })
    })

    describe("checkProposalhash", async () => {
        it("should return true if the proposal hash is present", async () => {
            await mintTo(accounts[0])
            await makeProposals()
            let isPresent = await dao.checkProposalHash(0, await accounts[2].getAddress(), ethers.utils.parseEther("1.0"), [])
            expect(isPresent).to.be.true
        })

        it("should return false if the proposal hash is not present", async () => {
            await mintTo(accounts[0])
            await makeProposals()
            let isPresent = await dao.checkProposalHash(0, await accounts[2].getAddress(), ethers.utils.parseEther("3.0"), [])
            expect(isPresent).to.be.false
        })
    })
})