import React, { useEffect, useState } from 'react';
import Header from 'layout/header';
import '../styles/auction.scss';
import { Link, useParams } from 'react-router-dom';
import { BigNumber, Contract, utils } from 'ethers'
import ActiveAuction from 'components/ActiveAuction';
import PastAuctions from 'components/PastAuctions';
import SelectedAuction from 'components/SelectedAuction';
import config from '../config';
import TokenABI from '../abi/contracts/Token.sol/Token.json';
import { useEthers } from '@usedapp/core';
import { Navigate } from 'react-router-dom';

const { addresses } = config
const tokenInterface = new utils.Interface(TokenABI);

const Auction = () => {
  const { library } = useEthers()
  let { tokenID } = useParams();
  let [validTokenID, setValidTokenID] = useState(false)
  let [dataLoaded, setDataLoaded] = useState(false)
  useEffect(() => {
    // validTokenID if it is either current or less than the number of tokens minted
    if (!tokenID) {
      // Redirect to /current
      return
    }
    const tokenContract = new Contract(addresses.token, tokenInterface, library)
    tokenContract.tokenQuantity().then((quant: BigNumber) => {
      setValidTokenID(quant.toString() < tokenID!)
      setDataLoaded(true)
    })
  }, [tokenID])

  if (!dataLoaded) {
    return <div>Loading...</div>
  }

  if (!tokenID || !validTokenID) {
    return <Navigate to="/auction/current" />
  }

  return (
    <div className='auction'>
      <Header>
        <Link to='/'>Home</Link>
      </Header>
      <div className='auction-wrapper'>
        {!tokenID || tokenID == 'current' ? <ActiveAuction /> : <SelectedAuction tokenID={tokenID} />}
        <div className='auction-spacer' />
        <PastAuctions selectedAuction={tokenID} />
      </div>
    </div>
  );
};

export default Auction;
