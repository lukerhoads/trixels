import React, { useEffect, useState } from 'react';
import { observer } from 'mobx-react';
import '../styles/past-auctions.scss';
import store from 'store';
import { ethers } from 'ethers';

type PastAuctionsProps = {
  selectedAuction?: string
}

const PastAuctions = ({ selectedAuction }: PastAuctionsProps) => {
  const [selAuction, setSelAuction] = useState()

  useEffect(() => {
    store.fetchPastAuctions();
  }, []);

  return (
    <div className='past-auctions'>
      {store.pastAuctions.map((auction, idx) => (
        <a href={`/auction/${auction.tokenID}`}>
          <div key={idx} className={'past-auction' + (auction.tokenID == selectedAuction ? ' active' : '')}>
            <img className='thumbnail' src={auction.imageUrl} />
            <div className='description'>
              <b>Trixel #{auction.tokenID}</b>
              <p>Sold for {auction.saleValue} ETH to </p>
            </div>
          </div>
        </a>
      ))}
    </div>
  );
};

export default observer(PastAuctions);
