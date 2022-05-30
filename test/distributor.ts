import { ethers } from "hardhat";
import { Signer } from "ethers";
import { expect } from "chai";
import { deployContract } from "ethereum-waffle";

import { Distributor, IDistributor } from '../typechain-types/contracts/Distributor'
import DistributorArtifact from '../artifacts/contracts/Distributor.sol/Distributor.json'

describe("Distributor", () => {
    let accounts: Signer[]
    let distributor: Distributor

    beforeEach(async () => {
        accounts = await ethers.getSigners()
        distributor = (await deployContract(accounts[0], DistributorArtifact)) as Distributor
    });

    const recordSale = async (tokenID: number, amt: string) => {
        let amount = ethers.utils.parseEther(amt)
        let depositTx = await distributor.recordSale(tokenID, { value: amount })
        await depositTx.wait()
    }

    const distributeToContributors = async () => {
        const contributors: IDistributor.ContributorStruct[] = [
            {
                addr: await accounts[0].getAddress(),
                numContribs: 10
            },
            {
                addr: await accounts[1].getAddress(),
                numContribs: 20
            },
            {
                addr: await accounts[2].getAddress(),
                numContribs: 40
            },
            {
                addr: await accounts[3].getAddress(),
                numContribs: 60
            }
        ]

        let distributeTx = await distributor.distribute(contributors, 1)
        await distributeTx.wait()
    }

    describe("recordSale", async () => {
        await recordSale(1, "1.0")
        let balance = await distributor.sales(0)
        expect(balance).to.eq("1.0")
    })

    describe("distribute", async () => {
        // await deposit(1, "1.0")
        let amount = ethers.utils.parseEther("1.0")
        let depositTx = await distributor.recordSale(1, { value: amount })
        await depositTx.wait()
        // await distributeToContributors()
        // let expectedBalances = [10, 10, 10, 10]
        // contributors.forEach(async (contributor, idx) => {
        //     let balance = await distributor.balances(contributor.addr)
        //     expect(balance).to.eq(expectedBalances[idx])
        // })
    })

    // describe("distrib", async () => {
    //     // await deposit(1, "1.0")
    // })

    describe("withdraw", async () => {

    })
})