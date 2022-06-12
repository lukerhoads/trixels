import { useEthers } from '@usedapp/core';
import ApiClient from 'api-client';
import { BigNumber, Contract, utils } from 'ethers';
import Header from 'layout/header';
import { useEffect, useMemo, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { LiveAuction, PastAuction, PastAuctionPreview } from 'types/auction';
import config from '../config';
import '../styles/auction.scss';

import Card from 'components/auction/Card';
import Main from 'components/auction/Main';
import Haptics from 'components/Haptics';
import { observer } from 'mobx-react';
import store from 'store';
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
        await Promise.all(trixels.map(async (trixel) => {
            let metadata = await apiClient.getMetadata(trixel.metadataUrl);
            pastAuctions.push({
                tokenID: trixel.tokenID,
                imageUrl: metadata.image,
                mintDate: trixel.createdAt,
            });
        }))

        setPastAuctions(prev => [...prev, ...pastAuctions]);
        setIsLoading(false);
    };

    const fetchActiveAuction = async () => {
        if (!auctionHouseContract || !tokenID) return;
        const auction = await auctionHouseContract.auction();
        if (auction.settled) return;
        console.log("auction: ", auction)
        console.log("Getting live")
        const liveAuction = await apiClient.getLiveTrixel();
        const metadata = await apiClient.getMetadata(liveAuction.metadataUrl);
        const bidIncrementPercentage = await auctionHouseContract.minBidIncrementPercentage();
        let minNextBid;
        const reservePrice = await auctionHouseContract.reservePrice();
        let highestBid = parseFloat(utils.formatEther(auction.highestBid))
        if (highestBid < reservePrice) {
            minNextBid = reservePrice.toNumber() * (1 + bidIncrementPercentage / 100);
        } else {
            minNextBid = highestBid * (1 + bidIncrementPercentage / 100);
        }
        setLiveAuction({
            tokenID: liveAuction.tokenID,
            imageUrl: metadata.image,
            endingDate: auction.endDate.toString(),
            highestBid: highestBid,
            highestBidder: auction.highestBidder,
            minNextBid: minNextBid,
        });
        setIsLoading(false);
    };

    const fetchPastAuction = async () => {
        if (!tokenID || !tokenContract || !auctionHouseContract) return;
        let trixel = await apiClient.getTrixelById(parseInt(tokenID));
        let metadata = await apiClient.getMetadata(trixel.metadataUrl);
        let filteredEvent = auctionHouseContract.filters.AuctionEnded(BigNumber.from(tokenID), null, null);
        console.log('Filtered event: ', filteredEvent);
        let filteredBurn = tokenContract.filters.Burn(BigNumber.from(1))
        console.log("burn: ", filteredBurn)
        let currentOwner = await tokenContract.ownerOf(tokenID);
        setPassedAuction({
            tokenID: parseInt(tokenID),
            imageUrl: metadata.image,
            salePrice: 0,
            mintDate: trixel.createdAt,
            winner: filteredEvent.address ? filteredEvent.address : "Nobody",
            currentOwner: currentOwner,
        });
        setIsLoading(false);
    };

    const placeBid = async () => {
        if (!auctionHouseContract || !library || !liveAuction) return;
        try {
            const auctionHouseConnected = auctionHouseContract.connect(library.getSigner());
            const tx = await auctionHouseConnected.placeBid(liveAuction.tokenID, { value: utils.parseEther(liveAuction.minNextBid.toString()) });
            await tx.wait();
        } catch (err: any) {
            store.pushToLogs({
                mood: 'error',
                message: `Error submitting bid: ${err.message}`
            })
            return
        }
        
        store.pushToLogs({
            mood: 'success',
            message: `Successfully bid ${liveAuction.minNextBid}`
        })

        fetchActiveAuction()
    };

    return (
        <Haptics type="logs">
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
        </Haptics>
    );
};

export default observer(Auction);
