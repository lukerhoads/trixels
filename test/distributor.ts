import { ethers, waffle } from "hardhat";
import { BigNumber, Signer } from "ethers";
import { expect } from "chai";

import { Distributor, IDistributor } from '../typechain-types/contracts/Distributor'
import { WETH } from '../typechain-types/contracts/WETH'
import WETHArtifact from '../artifacts/contracts/WETH.sol/WETH.json'

const { deployContract } = waffle

describe("Distributor", () => {
    let accounts: Signer[]
    let weth: WETH
    let distributor: Distributor

    type TestContributor = {
        accountIdx: number,
        numContribs: number
        expectedBalance: string
    }

    let contributors: TestContributor[]
    let expectedBalances: number[]

    beforeEach(async () => {
        accounts = await ethers.getSigners()
        weth = (await deployContract(accounts[0], WETHArtifact)) as WETH
        const distrib = await ethers.getContractFactory("Distributor")
        distributor = await distrib.deploy(weth.address)
        contributors =  [
            {
                accountIdx: 1,
                numContribs: 10,
                expectedBalance: "11111111111111111"
            },
            {
                accountIdx: 2,
                numContribs: 20,
                expectedBalance: "22222222222222222"
            },
            {
                accountIdx: 3,
                numContribs: 40,
                expectedBalance: "44444444444444444"
            },
            {
                accountIdx: 4,
                numContribs: 60,
                expectedBalance: "66666666666666666"
            }
        ]
    })

    const recordSale = async (tokenID: number, amt: string) => {
        let amount = ethers.utils.parseEther(amt)
        let depositTx = await distributor.recordSale(tokenID, { value: amount })
        await depositTx.wait()
    }

    const distributeToContributors = async (tokenID: number) => {
        const convContributors: IDistributor.ContributorStruct[] = await Promise.all(contributors.map(async contributor => ({
            addr: await accounts[contributor.accountIdx].getAddress(),
            numContribs: contributor.numContribs,
        })))

        let distributeTx = await distributor.distribute(convContributors, tokenID)
        await distributeTx.wait()
    }

    describe("recordSale", async () => {
        it("should record sale", async () => {
            await recordSale(1, "1.0")
            let balance = await distributor.sales(1)
            expect(balance).to.eq(ethers.utils.parseEther("1.0"))
        })
    })

    describe("distribute", async () => {
        it("should properly distribute to contributors", async () => {
            await recordSale(1, "1.0")
            await distributeToContributors(1)
            contributors.forEach(async (contributor, idx) => {
                let balance = await distributor.balances(await accounts[contributor.accountIdx].getAddress())
                expect(balance).to.eq(expectedBalances[idx])
            })
        })
    })

    describe("withdraw", async () => {
        it("should enable contributors to withdraw their earnings", async () => {
            await recordSale(1, "1.0")
            await distributeToContributors(1)
            contributors.forEach(async (contributor, idx) => {
                const contract = distributor.connect(accounts[contributor.accountIdx])
                let withdrawTx = await contract.withdraw()
                await withdrawTx.wait()
                const balance = await accounts[contributor.accountIdx].getBalance()
                expect(balance.toString()).to.eq(expectedBalances[idx])
            })
        })
    })
})