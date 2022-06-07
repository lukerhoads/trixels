import React from 'react';
import Header from 'layout/header';
import '../styles/auction.scss';
import { Link } from 'react-router-dom';
import ActiveAuction from 'components/ActiveAuction';
import PastAuctions from 'components/PastAuctions';

const Auction = () => {
  return (
    <div className='auction'>
      <Header>
        <Link to='/'>Home</Link>
      </Header>
      <div className="auction-wrapper">
        <ActiveAuction />
        <div className="auction-spacer" />
        <PastAuctions />
      </div>
    </div>
  );
};

export default Auction;
