import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import store from '../store'
import '../styles/active-auction.scss'
import { ethers } from 'ethers'
import config from '../config'

import Token  from '../../../artifacts/contracts/Token.sol/Token.json'
import AuctionHouse  from '../../../artifacts/contracts/AuctionHouse.sol/AuctionHouse.json'

const ActiveAuction = () => {
    const [auctionHighestBidder, setAuctionHighestBidder] = useState("")
    const [currentOwner, setCurrentOwner] = useState("")

    useEffect(() => {
        findTokenCurrentOwner()

        // Setup listener to update auction highest bidder
    }, [])

    const findTokenCurrentOwner = async () => {
        if (!store.activeAuction) return 
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        const token = new ethers.Contract(config.tokenAddress, Token.abi, provider)
        const tokenOwner = await token.ownerOf(store.activeAuction.tokenID)
        setCurrentOwner(tokenOwner)
    }

    const findAuctionHighestBidder = async () => {
        if (!store.activeAuction) return 
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        const auctionHouse = new ethers.Contract(config.auctionHouseAddress, AuctionHouse.abi, provider)
        const auction = await auctionHouse.auction() 
        setAuctionHighestBidder(auction.highestBidder)
    }

    return (
        <div className='active-auction'>
            { store.activeAuction ? (
                <div className='auc-wrap'>
                    <img className='active-auction-image' src={"https://cdn.britannica.com/91/181391-050-1DA18304/cat-toes-paw-number-paws-tiger-tabby.jpg?q=60"}  />
                    <div className='active-auction-data'>
                        <div className="section">
                            <p className="caption">Trixel ID:</p>
                            <p className='value'>#{store.activeAuction.tokenID}</p>
                            <p className="caption">Sale price:</p>
                            <p className='value'>{store.activeAuction.saleValue} ETH</p>
                            <p className="caption">Highest bidder:</p>
                            <p className='value'>{auctionHighestBidder}</p>
                        </div>
                        <div className="section">
                            <p className="caption">Sale date:</p>
                            <p className='value'>{store.activeAuction.createdAt}</p>
                            <p className="caption">Current owner:</p>
                            <p className='value'>{currentOwner}</p>
                        </div>
                    </div>
                </div>
            ) : (
                <>
                    <p>No currently active auction</p>
                </>
            )}
        </div>
    )
}

export default observer(ActiveAuction)