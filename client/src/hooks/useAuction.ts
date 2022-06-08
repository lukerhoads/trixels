import { useCall, useEthers } from '@usedapp/core';
import { BigNumber, Contract, utils } from 'ethers';
import { useCallback, useEffect, useMemo, useState } from 'react';
import config from '../config';
import AuctionHouseABI from '../abi/contracts/AuctionHouse.sol/AuctionHouse.json';
import DistributorABI from '../abi/contracts/Distributor.sol/Distributor.json';
import TokenABI from '../abi/contracts/Token.sol/Token.json';

// import { AuctionHouse__factory, AuctionHouse } from '../../typechain-types';

const auctionHouseInterface = new utils.Interface(AuctionHouseABI);
const distributorInterface = new utils.Interface(DistributorABI);
const tokenInterface = new utils.Interface(TokenABI);

const { addresses } = config;

export type ContractAuction = {
  startDate: BigNumber;
  endDate: BigNumber;
  highestBid: BigNumber;
  highestBidder: string;
  tokenId: BigNumber;
  settled: boolean;
};

type AuctionInfo = {
  auction?: ContractAuction;
  placeBid: (tokenId: BigNumber) => Promise<void>;
  subscribeToNewBid: (callback: (highestBidder: string, highestBid: BigNumber) => void) => void;
};

export const useAuction = (): AuctionInfo => {
  const { library } = useEthers();

  const [auction, setAuction] = useState<ContractAuction | undefined>(undefined);

  const auctionHouseContract = useMemo((): Contract | undefined => {
    if (!library || !addresses.auctionHouse) return;
    // return AuctionHouse__factory.connect(addresses.auctionHouse, library);
    return new Contract(addresses.auctionHouse, auctionHouseInterface, library);
  }, [library]);

  useEffect(() => {
    if (!auctionHouseContract) return;
    auctionHouseContract.auction().then((auction: ContractAuction) => {
      setAuction(auction);
    });
  }, [auctionHouseContract]);

  const placeBid = useCallback(
    async (tokenID: BigNumber) => {
      if (!auctionHouseContract || !library) return;
      const auctionHouseConnected = auctionHouseContract.connect(library.getSigner());
      const tx = await auctionHouseConnected.placeBid(tokenID.toString());
      await tx.wait();
    },
    [auctionHouseContract]
  );

  const subscribeToNewBid = useCallback(
    (callback: (highestBidder: string, highestBid: BigNumber) => void) => {
      if (!auctionHouseContract || !library) return;
      auctionHouseContract.on('BidPlaced', (_, highestBidder, highestBid) => callback(highestBidder, highestBid));
    },
    [auctionHouseContract]
  );

  return {
    auction,
    placeBid,
    subscribeToNewBid,
  };
};

export type PastAuction = {
  trixelId: number;
  salePrice: number;
  saleDate: string;
  winner: string;
  owner: string;
};

export type PastAuctionInfo = {
  auction?: PastAuction;
};

export const usePastAuction = (tokenID: number): PastAuctionInfo => {
  const { library } = useEthers();

  const [auction, setAuction] = useState<PastAuction | undefined>(undefined);

  const auctionHouseContract = useMemo((): Contract | undefined => {
    if (!library || !addresses.auctionHouse) return;
    // return AuctionHouse__factory.connect(addresses.auctionHouse, library);
    return new Contract(addresses.auctionHouse, auctionHouseInterface, library);
  }, [library]);

  const distributorContract = useMemo((): Contract | undefined => {
    if (!library || !addresses.auctionHouse) return;
    // return AuctionHouse__factory.connect(addresses.auctionHouse, library);
    return new Contract(addresses.distributor, distributorInterface, library);
  }, [library]);

  const tokenContract = useMemo((): Contract | undefined => {
    if (!library || !addresses.auctionHouse) return;
    // return AuctionHouse__factory.connect(addresses.auctionHouse, library);
    return new Contract(addresses.token, tokenInterface, library);
  }, [library]);

  useEffect(() => {
    if (!auctionHouseContract || !tokenContract || !distributorContract) return;
    let resAuction: PastAuction;

    distributorContract.sales(tokenID).then((salePrice: BigNumber) => (resAuction.salePrice = salePrice.toNumber()));
  }, [auctionHouseContract, distributorContract, tokenContract]);

  return {
    auction,
  };
};
