import { useEthers } from '@usedapp/core';
import ApiClient from 'api-client';
import { BigNumber, Contract, utils } from 'ethers';
import Header from 'layout/header';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { LiveAuction, PastAuction, PastAuctionPreview } from 'types/auction';
import config from '../config';
import '../styles/auction.scss';

import Card from 'components/auction/Card';
import Main from 'components/auction/Main';
import AuctionHouseABI from '../abi/contracts/AuctionHouse.sol/AuctionHouse.json';
import TokenABI from '../abi/contracts/Token.sol/Token.json';

const { addresses } = config;
const auctionHouseInterface = new utils.Interface(AuctionHouseABI);
const tokenInterface = new utils.Interface(TokenABI);

const Auction = () => {
    let { tokenID } = useParams();
    const { library } = useEthers();
    const [loading, setIsLoading] = useState(true);
    const [validTokenID, setValidTokenID] = useState(true);
    const [isLive, setIsLive] = useState(false);

    const [tokenQuantity, setTokenQuantity] = useState(0);
    const [liveAuction, setLiveAuction] = useState<LiveAuction | undefined>(undefined);
    const [passedAuction, setPassedAuction] = useState<PastAuction | undefined>(undefined);
    const [pastAuctions, setPastAuctions] = useState<PastAuctionPreview[]>([]);

    const auctionHouseContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.auctionHouse, auctionHouseInterface, library);
    }, [library]);

    const tokenContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.token, tokenInterface, library);
    }, [library]);

    const apiClient = useMemo((): ApiClient => {
        return new ApiClient(config.alt.apiUrl);
    }, []);

    useEffect(() => {
        if (tokenID === 'current') {
            setIsLive(true);
            fetchActiveAuction();
        } else {
            fetchPastAuction();
        }

        fetchPastAuctions();
        fetchTokenQuantity();
    }, []);

    useEffect(() => {
        if (!auctionHouseContract) return;
        auctionHouseContract.on('BidPlaced', () => {
            console.log('Bid placed');
        });

        return () => {
            auctionHouseContract.removeAllListeners('BidPlaced');
        };
    }, [auctionHouseContract]);

    useEffect(() => {
        if (!tokenID) {
            setValidTokenID(false);
            return;
        }

        setValidTokenID(parseInt(tokenID) < tokenQuantity);
    }, [tokenQuantity]);

    const fetchTokenQuantity = async () => {
        if (!tokenContract) return;
        const tokenQuantity = await tokenContract.tokenQuantity();
        setTokenQuantity(tokenQuantity.toNumber());
    };

    const fetchPastAuctions = async () => {
        let trixels = await apiClient.getTrixels();
        let pastAuctions: PastAuctionPreview[] = [];
        trixels.forEach(async (trixel) => {
            let metadata = await apiClient.getMetadata(trixel.metadataUrl);
            pastAuctions.push({
                tokenID: trixel.tokenID,
                imageUrl: metadata.image,
                mintDate: trixel.createdAt,
            });
        });

        setPastAuctions(pastAuctions);
        setIsLoading(false);
    };

    const fetchActiveAuction = async () => {
        if (!auctionHouseContract || !tokenID) return;
        const auction = await auctionHouseContract.auction();
        console.log('Auction: ', auction);
        if (auction.settled) return;
        const liveAuction = await apiClient.getLiveTrixel();
        const metadata = await apiClient.getMetadata(liveAuction.metadataUrl);
        const bidIncrementPercentage = await auctionHouseContract.minBidIncrementPercentage();
        let minNextBid;
        const reservePrice = await auctionHouseContract.reservePrice();
        if (auction.highestBid.toNumber() < reservePrice) {
            minNextBid = reservePrice + (1 + bidIncrementPercentage / 100) * reservePrice;
        } else {
            minNextBid = auction.highestBid.toNumber() + (1 + bidIncrementPercentage / 100) * auction.highestBid.toNumber();
        }
        setLiveAuction({
            tokenID: liveAuction.tokenID,
            imageUrl: metadata.image,
            endingDate: auction.endDate.toString(),
            highestBid: auction.highestBid.toNumber(),
            highestBidder: auction.highestBidder,
            minNextBid: minNextBid,
        });
        setIsLoading(false);
    };

    const fetchPastAuction = async () => {
        if (!tokenID || !tokenContract || !auctionHouseContract) return;
        let trixel = await apiClient.getTrixelById(parseInt(tokenID));
        let metadata = await apiClient.getMetadata(trixel.metadataUrl);
        let currentOwner = await tokenContract.ownerOf(tokenID);
        let filteredEvent = auctionHouseContract.filters.AuctionEnded(tokenID, null, null);
        console.log('Filtered event: ', filteredEvent);
        setPassedAuction({
            tokenID: parseInt(tokenID),
            imageUrl: metadata.image,
            salePrice: 0,
            mintDate: trixel.createdAt,
            winner: '',
            currentOwner: currentOwner,
        });
        setIsLoading(false);
    };

    const placeBid = useCallback(async () => {
        if (!auctionHouseContract || !library) return;
        const auctionHouseConnected = auctionHouseContract.connect(library.getSigner());
        const tx = await auctionHouseConnected.placeBid(BigNumber.from(tokenID).toString());
        await tx.wait();
    }, [auctionHouseContract]);

    return (
        <div className='auction'>
            <Header>
                <Link to='/'>Home</Link>
                <Link to='/dao'>Dao</Link>
            </Header>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <div className='auction-wrapper'>
                    <div className='active-auction'>
                        <Main
                            live={isLive}
                            liveAuction={liveAuction ? liveAuction : undefined}
                            pastAuction={passedAuction ? passedAuction : undefined}
                            placeBid={placeBid}
                        />
                    </div>
                    <div className='auction-spacer' />
                    <div className='past-auctions'>
                        {pastAuctions.map((pastAuction, idx) => (
                            <Card key={idx} tokenID={pastAuction.tokenID} imageUrl={pastAuction.imageUrl} mintDate={pastAuction.mintDate} active={false} />
                        ))}
                        {pastAuctions.length ? null : <p>No past auctions</p>}
                    </div>
                </div>
            )}
        </div>
    );
};

export default Auction;
