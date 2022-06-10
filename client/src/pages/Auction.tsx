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

    const [tokenQuantity, setTokenQuantity] = useState(0)
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
            fetchActiveAuction()
        } else {
            fetchPastAuction()
        }
        
        fetchPastAuctions()
        fetchTokenQuantity()
    }, [])

    useEffect(() => {
        if (!auctionHouseContract) return
        auctionHouseContract.on("BidPlaced", () => {
            console.log("Bid placed")
        })

        return () => {
            auctionHouseContract.removeAllListeners("BidPlaced")
        }
    }, [auctionHouseContract])

    useEffect(() => {
        if (!tokenID) {
            setValidTokenID(false)
            return
        }

        setValidTokenID(parseInt(tokenID) < tokenQuantity)
    }, [tokenQuantity])

    const fetchTokenQuantity = async () => {
        if (!tokenContract) return
        const tokenQuantity = await tokenContract.tokenQuantity()
        setTokenQuantity(tokenQuantity.toNumber())
    }

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
        setIsLoading(false)
    }

    const fetchActiveAuction = async () => {
        if (!auctionHouseContract) return 
        const auction = await auctionHouseContract.auction()
        console.log("Auction: ", auction)
        console.log("Auction startDate: ", auction.startDate.toNumber())
        if (auction.settled) return
        setLiveAuction(auction)
        setIsLoading(false)
    }

    const fetchPastAuction = async () => {
        if (!tokenID || !tokenContract || !auctionHouseContract) return
        let trixel = await apiClient.getTrixelById(parseInt(tokenID))
        let metadata = await apiClient.getMetadata(trixel.metadataUrl)
        let currentOwner = await tokenContract.ownerOf(tokenID)
        // Need to get the sale value along with the winner (event filter)
        let filteredEvent = auctionHouseContract.filters.AuctionEnded(tokenID, null, null)
        console.log("Filtered event: ", filteredEvent)
        setPassedAuction({
            tokenID: parseInt(tokenID),
            imageUrl: metadata.image,
            salePrice: 0,
            mintDate: trixel.createdAt,
            winner: "",
            currentOwner: currentOwner,
        })
        setIsLoading(false)
    }

    const placeBid = useCallback(async () => {
        if (!auctionHouseContract || !library) return;
        const auctionHouseConnected = auctionHouseContract.connect(library.getSigner());
        const tx = await auctionHouseConnected.placeBid(BigNumber.from(tokenID).toString());
        await tx.wait();
    }, [auctionHouseContract])

    return (
        <div className='auction'>
            <Header>
                <Link to='/'>Home</Link>
            </Header>
            { loading ? (
                <p>Loading...</p>
            ) : (
                <div className='auction-wrapper'>
                    <div className='active-auction'>
                        <Main live={isLive} liveAuction={isLive ? liveAuction : undefined} pastAuction={isLive ? undefined : passedAuction} placeBid={placeBid} />
                    </div>
                    <div className='auction-spacer' />
                    <div className='past-auctions'>
                        {pastAuctions.map((pastAuction, idx) => (
                            <a key={idx} href={`/auction/${pastAuction.tokenID}`}>
                                <div className="past-auction">
                                    <img className="thumbnail" src={pastAuction.imageUrl} />
                                    <div className="description">
                                        <p>Trixel #{pastAuction.tokenID}</p>
                                        <p>Minted on {pastAuction.mintDate}</p>
                                    </div>
                                </div>
                            </a>
                        ))}
                        {pastAuctions.length ? null : <p>No past auctions</p>}
                    </div>
                </div>
            )}
        </div>
    );
};

export default Auction;
