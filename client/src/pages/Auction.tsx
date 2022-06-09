import React, { useCallback, useEffect, useMemo, useState } from 'react';
import Header from 'layout/header';
import '../styles/auction.scss';
import { Link, useParams } from 'react-router-dom';
import { BigNumber, Contract, utils } from 'ethers';
import config from '../config';
import { useEthers } from '@usedapp/core';
import ApiClient from 'api-client';
import { LiveAuction, PastAuction, PastAuctionPreview } from 'types/auction';

import AuctionHouseABI from '../abi/contracts/AuctionHouse.sol/AuctionHouse.json';
import DistributorABI from '../abi/contracts/Distributor.sol/Distributor.json';
import TokenABI from '../abi/contracts/Token.sol/Token.json';
import Main from 'components/auction/Main';

const { addresses } = config;
const auctionHouseInterface = new utils.Interface(AuctionHouseABI);
const distributorInterface = new utils.Interface(DistributorABI);
const tokenInterface = new utils.Interface(TokenABI);

const Auction = () => {
    let { tokenID } = useParams();
    const [loading, setIsLoading] = useState(true)
    const [validTokenID, setValidTokenID] = useState(true)
    const [isLive, setIsLive] = useState(false)

    const [liveAuction, setLiveAuction] = useState<LiveAuction | undefined>(undefined)
    const [passedAuction, setPassedAuction] = useState<PastAuction | undefined>(undefined)
    const [pastAuctions, setPastAuctions] = useState<PastAuctionPreview[]>([])

    const { library } = useEthers()

    const auctionHouseContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.auctionHouse, auctionHouseInterface, library);
    }, [library]);

    const distributorContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.distributor, distributorInterface, library);
    }, [library]);

    const tokenContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.token, tokenInterface, library);
    }, [library]);

    const apiClient = useMemo((): ApiClient => {
        return new ApiClient(config.alt.apiUrl)
    }, [])

    useEffect(() => {
        if (tokenID === 'current') {
            setIsLive(true)
        }
        
        fetchPastAuctions()
        if (isLive) {
            fetchActiveAuction()
        } else {

        }
    }, [])

    useEffect(() => {
        if (!auctionHouseContract) return
        auctionHouseContract.on("BidPlaced", () => {
            console.log("Bid placed")
        })
    }, [auctionHouseContract])

    const fetchPastAuctions = async () => {
        let trixels = await apiClient.getTrixels()
        let pastAuctions: PastAuctionPreview[] = []
        trixels.forEach(async trixel => {
            let metadata = await apiClient.getMetadata(trixel.metadataUrl)
            pastAuctions.push({
                tokenID: trixel.tokenID,
                imageUrl: metadata.image,
                mintDate: trixel.createdAt // This is the mint date, not create date
            })
        })

        setPastAuctions(pastAuctions)
    }

    const fetchActiveAuction = async () => {
        if (!auctionHouseContract) return 
        setLiveAuction(await auctionHouseContract.auction())
    }

    const fetchPastAuction = async () => {
        
    }

    const placeBid = useCallback(async () => {
        if (!auctionHouseContract || !library) return;
        const auctionHouseConnected = auctionHouseContract.connect(library.getSigner());
        const tx = await auctionHouseConnected.placeBid(BigNumber.from(tokenID).toString());
        await tx.wait();
    }, [auctionHouseContract])

    // Here is the plan:
    // Load data in here. Data we need is:
    // - Past auctions (tokenId, saleValue, winner)
    // - Active auction (stuff in contract)
    // - Past auction, but more in depth (sale date, current owner along with original data)
    // - and of course, the image url.

    return (
        <div className='auction'>
            <Header>
                <Link to='/'>Home</Link>
            </Header>
            { loading ? (
                <p>Loading...</p>
            ) : (
                <div className='auction-wrapper'>
                    <div className='auction'>
                        <Main live={isLive} liveAuction={isLive ? liveAuction : undefined} pastAuction={isLive ? undefined : passedAuction} placeBid={placeBid} />
                    </div>
                    <div className='auction-spacer' />
                    <div className='past-auctions'>
                        {pastAuctions.map(pastAuction => (
                            <a href={`/auction/${pastAuction.tokenID}`}>
                                <div className="past-auction">
                                    <img className="thumbnail" src={pastAuction.imageUrl} />
                                    <div className="description">
                                        <p>Trixel #{pastAuction.tokenID}</p>
                                        <p>Minted on {pastAuction.mintDate}</p>
                                    </div>
                                </div>
                            </a>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default Auction;
