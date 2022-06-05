import React from 'react';
import Header from 'layout/header';
import '../styles/auction.scss';
import { Link } from 'react-router-dom';

const Auction = () => {
  return (
    <div className='auction'>
      <Header>
        <Link to='/'>Home</Link>
      </Header>
    </div>
  );
};

export default Auction;
