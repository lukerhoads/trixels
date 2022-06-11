import { useEthers } from '@usedapp/core';
import { BigNumber, Contract, utils } from 'ethers';
import { useEffect, useMemo, useState } from 'react';
import config from '../config';
import AuctionHouseABI from '../abi/contracts/AuctionHouse.sol/AuctionHouse.json';
import DistributorABI from '../abi/contracts/Distributor.sol/Distributor.json';
import TokenABI from '../abi/contracts/Token.sol/Token.json';

const { addresses } = config;
const auctionHouseInterface = new utils.Interface(AuctionHouseABI);
const distributorInterface = new utils.Interface(DistributorABI);
const tokenInterface = new utils.Interface(TokenABI);

export type PastAuction = {
    trixelId: number;
    salePrice: number;
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
        let resAuction: PastAuction = {
            trixelId: tokenID,
            salePrice: 0,
            winner: '',
            owner: '',
        };
        distributorContract.sales(tokenID).then((salePrice: BigNumber) => (resAuction.salePrice = salePrice.toNumber()));
        // Source auction winner from the event
        const event = auctionHouseContract.filters.AuctionEnded(tokenID, null, null);
        if (event.topics) {
            resAuction.winner = event.topics[2][0];
        }
        // Source the current owner from the token contract
        tokenContract.ownerOf(tokenID).then((ownerAddress: string) => (resAuction.owner = ownerAddress));
        setAuction(resAuction);
    }, [auctionHouseContract, distributorContract, tokenContract]);

    return {
        auction,
    };
};
