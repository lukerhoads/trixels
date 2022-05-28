import { ethers } from "hardhat";
import { Signer } from "ethers";

import { Token } from './bindings/contracts/Token'
import TokenArtifact from '../artifacts/contracts/Token.sol/Token.json'
import { deployContract } from "ethereum-waffle";
import { expect } from "chai";

describe("Token", () => {
    let accounts: Signer[]
    let token: Token

    beforeEach(async () => {
        accounts = await ethers.getSigners();
        token = (await deployContract(accounts[0], TokenArtifact)) as Token
    });

    describe("mint", async () => {
        it("should mint", async () => {
            let tokenQuantity = await token.tokenQuantity()
            expect(tokenQuantity).to.eq(0)
            let tokenID = await token.mint(await accounts[0].getAddress())
            expect(tokenID).to.eq(1)
            let newTokenQuantity = await token.tokenQuantity()
            expect(newTokenQuantity).to.eq(1)
            let balance = await token.balanceOf(await accounts[0].getAddress())
            expect(balance).to.eq(1)
        })
    });

    describe("burn", async () => {
        it("should burn", async () => {
            let tokenID = await token.mint(await accounts[0].getAddress())
            await token.burn(tokenID)
        })
    });

    it("should return proper token URI", async () => {
        let tokenUri = await token.tokenURI(0)
        expect(tokenUri).to.eq("")
    });
})