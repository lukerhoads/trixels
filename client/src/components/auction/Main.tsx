import { useEthers } from '@usedapp/core';
import Address from 'components/Address';
import { observer } from 'mobx-react';
import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import store from 'store';
import { LiveAuction, PastAuction } from 'types/auction';

// - Trixel ID (complete)
// - Sale price (source from distributor contract)
// - Auction winner (event in auctionHouse contract)
// - Current owner (source from token contract)

type MainProps = {
    live: boolean;
    liveAuction?: LiveAuction;
    pastAuction?: PastAuction;
    placeBid: () => void;
};

const Main = ({ live, liveAuction, pastAuction, placeBid }: MainProps) => {
    const { account, deactivate } = useEthers();
    const [bothUndefined, setBothUndefined] = useState(false);

    useEffect(() => {
        if (!liveAuction && !pastAuction) {
            setBothUndefined(true);
        } else {
            setBothUndefined(false);
        }
    }, [liveAuction, pastAuction]);

    if (bothUndefined) {
        return <p>No live or past auction</p>;
    }

    const authClick = () => {
        store.setOverlayInfo({
            type: 'auth-modal',
        });
    };

    const parseEthereumDate = (date: string | undefined) => {
        if (!date) return ""
        return new Date(parseInt(date) * 1000).toUTCString()
    }

    return (
        <div>
            <img className='active-image' src={live ? liveAuction?.imageUrl : pastAuction?.imageUrl} />
            <div className='active-data'>
                <div className='split'>
                    <div className='section'>
                        {live ? (
                            <>
                                <p className='caption'>Trixel ID:</p>
                                <p className='value'>{liveAuction?.tokenID}</p>
                                <p className='caption'>Ending date:</p>
                                <p className='value'>{parseEthereumDate(liveAuction?.endingDate)}</p>
                            </>
                        ) : (
                            <>
                                <p className='caption'>Trixel ID:</p>
                                <p className='value'>{pastAuction?.tokenID}</p>
                                <p className='caption'>Sale date:</p>
                                <p className='value'>{pastAuction?.mintDate}</p>
                            </>
                        )}
                    </div>
                    <div className='section'>
                        {live ? (
                            <>
                                <p className='caption'>Highest Bid:</p>
                                <p className='value'>{liveAuction?.highestBid}</p>
                                <p className='caption'>Minimum Next Bid:</p>
                                <p className='value'>{liveAuction?.minNextBid}</p>
                            </>
                        ) : (
                            <>
                                <p className='caption'>Sale price:</p>
                                <p className='value'>{pastAuction?.salePrice}</p>
                            </>
                        )}
                    </div>
                </div>
            </div>
            {live ? (
                <>
                    <p className='caption'>Highest Bidder: </p>
                    <Address className='value'>{liveAuction?.highestBidder}</Address>
                    {account ? <button onClick={() => placeBid()}>Bid</button> : <button onClick={() => authClick()}>Sign in</button>}
                    <p className="clicktext" onClick={() => deactivate()}>Disconnect</p>
                </>
            ) : (
                <>
                    <p className='caption'>Auction Winner: </p>
                    <p className='value'>{pastAuction?.winner}</p>
                    <p className='caption'>Current Owner: </p>
                    <p className='value'>{pastAuction?.currentOwner}</p>
                    <Link to="/auction/current">Back to current</Link>
                </>
            )}
        </div>
    );
};

export default observer(Main);
