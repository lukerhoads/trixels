import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import '../styles/past-auctions.scss'
import store from 'store'
import { ethers } from 'ethers'
import config from '../config'

import AuctionHouse from '../../../artifacts/contracts/AuctionHouse.sol/AuctionHouse.json'
import { Contract } from 'ethers'

const PastAuctions = () => {
    const [auctionHouse, setAuctionHouse] = useState<Contract | undefined>(undefined)

    useEffect(() => {
        // store.fetchPastAuctions()
        const prov = new ethers.providers.Web3Provider(window.ethereum)
        const aucHouse = new ethers.Contract(config.auctionHouseAddress, AuctionHouse.abi, prov)
        setAuctionHouse(aucHouse)
        console.log(store.pastAuctions)
    }, [])

    const getAuctionWinner = async (tokenID: number) => {
        if (!auctionHouse) return
        let filter = auctionHouse?.filters.AuctionEnded(tokenID, null, null)
        let events = await auctionHouse?.queryFilter(filter)
        console.log(events)
        return ""
    }

    return (
        <div className='past-auctions'>
            {store.pastAuctions.map((auction, idx) => (
                <a href={`/auction/${auction.tokenID}`}>
                    <div key={idx} className='past-auction'>
                        <img className='thumbnail' src={auction.imageUrl} />
                        <div className='description'>
                            <b>Trixel #{auction.tokenID}</b>
                            <p>Sold for {auction.saleValue} ETH to </p>
                        </div>
                    </div>
                </a>
            ))}
        </div>
    )
}

export default observer(PastAuctions)