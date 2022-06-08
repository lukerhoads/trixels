import React, { useEffect, useState } from 'react';
import '../styles/active-auction.scss';

import { useAuction } from 'hooks/useAuction';

// TODO:
// - Figure out how to do typed contracts
// - Add the bid button with a warning that a bid is non-retractable
// - Add footer
// - Make auction image square
// - Have invalid tokenIDs route to current auction
// - Catch metamask decline transaction

const ActiveAuction = () => {
  const { auction, placeBid, subscribeToNewBid } = useAuction();
  const [isLiveAuction, setIsLiveAuction] = useState(false);

  useEffect(() => {
    if (!auction) return;
    if (auction && !auction.settled) {
      setIsLiveAuction(true);
    }

    // Setup listener to update auction highest bidder
    subscribeToNewBid((highestBidder, highestBid) => {
      auction.highestBidder = highestBidder;
      auction.highestBid = highestBid;
    });
  }, [auction]);

  return (
    <div className='active-auction'>
      <div className='auc-wrap'>
        {auction && isLiveAuction ? (
          <>
            <img className='active-image' src={'https://cdn.britannica.com/91/181391-050-1DA18304/cat-toes-paw-number-paws-tiger-tabby.jpg?q=60'} />
            <div className='active-data'>
              <div className='section'>
                <p className='caption'>Trixel ID:</p>
                <p className='value'>{auction.tokenId.toString()}</p>
                <p className='caption'>Highest bid:</p>
                <p className='value'>{auction.highestBid.toString()}</p>
              </div>
              <div className='section'>
                <p className='caption'>Ending date:</p>
                <p className='value'>{auction.endDate.toString()}</p>
                <p className='caption'>Minimum bid increment:</p>
                <p className='value'></p>
              </div>
            </div>
            <p className='caption'>Highest bidder:</p>
            <p className='value'>{auction.highestBidder}</p>
            <button onClick={() => placeBid(auction.tokenId)}>Bid</button>
          </>
        ) : (
          <p>No active auction, come back in</p>
        )}
      </div>
    </div>
  );
};

export default ActiveAuction;
