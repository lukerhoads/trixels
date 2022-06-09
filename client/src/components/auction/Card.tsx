import React from 'react';
import '../styles/auction-card.scss';

type CardProps = {
    tokenID: number;
    saleValue: number;
    imageUrl: string;
    active: boolean;
};

const Card = ({ tokenID, saleValue, imageUrl, active }: CardProps) => {
    return (
        <a href={`/auction/${tokenID}`}>
            <div className={'past-auction' + (active ? ' active' : '')}>
                <img className='thumbnail' src={imageUrl} />
                <div className='description'>
                    <b>Trixel #{tokenID}</b>
                    <p>Sold for {saleValue} ETH to </p>
                </div>
            </div>
        </a>
    );
};

export default Card;
