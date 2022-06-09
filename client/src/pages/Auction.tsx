import React, { useEffect, useState } from 'react';
import Header from 'layout/header';
import '../styles/auction.scss';
import { Link, useParams } from 'react-router-dom';
import ActiveAuction from 'components/ActiveAuction';
import PastAuctions from 'components/PastAuctions';
import SelectedAuction from 'components/SelectedAuction';

const Auction = () => {
  let { tokenID } = useParams();
  let [validTokenID, setValidTokenID] = useState(false)
  useEffect(() => {
    // validTokenID if it is either current or less than the number of tokens minted
    setValidTokenID(true)
  }, [])

  return (
    <div className='auction'>
      <Header>
        <Link to='/'>Home</Link>
      </Header>
      <div className='auction-wrapper'>
        {!tokenID || tokenID == 'current' ? <ActiveAuction /> : <SelectedAuction tokenID={tokenID} />}
        <div className='auction-spacer' />
        <PastAuctions selectedAuction={validTokenID ? tokenID : undefined} />
      </div>
    </div>
  );
};

export default Auction;
