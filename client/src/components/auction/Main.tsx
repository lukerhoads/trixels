import { BigNumber } from 'ethers';
import React, { useState, useEffect } from 'react';
import { LiveAuction, PastAuction } from 'types/auction';

// - Trixel ID (complete)
// - Sale price (source from distributor contract)
// - Auction winner (event in auctionHouse contract)
// - Current owner (source from token contract)

type MainProps = {
    live: boolean
    liveAuction?: LiveAuction
    pastAuction?: PastAuction
    placeBid: () => void
}

const Main = ({ live, liveAuction, pastAuction, placeBid }: MainProps) => {
    return (
        <div className='auction-wrapper'>
            <img className='active-image' src={live ? liveAuction?.imageUrl : pastAuction?.imageUrl} />
            <div className='active-data'>
                <div className='split'>
                    <div className='section'>
                        {live ? (
                            <>
                                <p className='caption'>Trixel ID:</p>
                                <p className='value'>{liveAuction?.tokenID}</p>
                                <p className='caption'>Ending date:</p>
                                <p className='value'>{liveAuction?.endingDate}</p>
                            </>
                        ) : (
                            <>
                                <p className='caption'>Trixel ID:</p>
                                <p className='value'>{pastAuction?.tokenID}</p>
                                <p className='caption'>Sale date:</p>
                                <p className='value'>{pastAuction?.mintDate}</p>
                            </>
                        )}
                    </div>
                    <div className='section'>
                        {live ? (
                            <>
                                <p className='caption'>Highest Bid:</p>
                                <p className='value'>{liveAuction?.tokenID}</p>
                                <p className='caption'>Minimum Bid Increment:</p>
                                <p className='value'>{liveAuction?.endingDate}</p>
                            </>
                        ) : (
                            <>
                                <p className='caption'>Sale price:</p>
                                <p className='value'>{pastAuction?.salePrice}</p>
                            </>
                        )}
                    </div>
                </div>
                { live ? (
                    <>
                        <p className='caption'>Highest Bidder: </p>
                        <p className='value'>{liveAuction?.highestBidder}</p>
                        <button onClick={() => placeBid()}>Bid</button>
                    </>
                ) : (
                    <>
                        <p className='caption'>Auction Winner: </p>
                        <p className='value'>{pastAuction?.winner}</p>
                        <p className='caption'>Current Owner: </p>
                        <p className='value'>{pastAuction?.currentOwner}</p>
                    </>
                )}
            </div>
        </div>
    );
};

export default Main;
