import { ethers, network } from "hardhat";
import { Signer } from "ethers";
import { expect } from "chai";
import { deployContract } from "ethereum-waffle";

import { Token } from '../typechain-types/contracts/Token'
import TokenArtifact from '../artifacts/contracts/Token.sol/Token.json'

describe("Token", () => {
    let accounts: Signer[]
    let token: Token

    beforeEach(async () => {
        accounts = await ethers.getSigners();
        token = (await deployContract(accounts[0], TokenArtifact)) as Token
    });

    const mint = async () => {
        let mintTx = await token.mint(await accounts[0].getAddress())
        let trace = await network.provider.send('debug_traceTransaction', [mintTx.hash])
        const [tokenID] = ethers.utils.defaultAbiCoder.decode(
            ['uint'],
            `0x${trace.returnValue}`
        )
        return tokenID
    }

    describe("mint", async () => {
        it("should mint", async () => {
            let tokenQuantity = await token.tokenQuantity()
            expect(tokenQuantity).to.eq(0)
            let tokenID = await mint()
            expect(tokenID).to.eq(1)
            let newTokenQuantity = await token.tokenQuantity()
            expect(newTokenQuantity).to.eq(1)
            let balance = await token.balanceOf(await accounts[0].getAddress())
            expect(balance).to.eq(1)
        })
    });

    describe("burn", async () => {
        it("should burn", async () => {
            let tokenID = await mint()
            let tokenQuantity = await token.tokenQuantity()
            expect(tokenQuantity).to.eq(1)
            await token.burn(tokenID)
            tokenQuantity = await token.tokenQuantity()
            expect(tokenQuantity).to.eq(0)
        })
    });

    it("should return proper token URI", async () => {
        await mint()
        let tokenUri = await token.tokenURI(1)
        expect(tokenUri).to.eq("")
    });
})