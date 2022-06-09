import React, { useEffect, useState } from 'react';
import { observer } from 'mobx-react';
import '../styles/past-auctions.scss';
import store from 'store';
import PastAuction from './PastAuction'

type PastAuctionsProps = {
  selectedAuction?: string
}

const PastAuctions = ({ selectedAuction }: PastAuctionsProps) => {
  useEffect(() => {
    store.fetchPastAuctions();
  }, []);

  return (
    <div className='past-auctions'>
      {store.pastAuctions.map((auction, idx) => (
        <>
        <PastAuction key={idx} tokenID={auction.tokenID} saleValue={auction.saleValue} imageUrl={auction.imageUrl} active={auction.tokenID.toString() == selectedAuction} />
        </>
      ))}
    </div>
  );
};

export default observer(PastAuctions);
