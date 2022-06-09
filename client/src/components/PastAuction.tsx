import React from 'react'
import '../styles/past-auction.scss'

type PastAuctionProps = {
    tokenID: number
    saleValue: number
    imageUrl: string
    active: boolean
}

const PastAuction = ({ tokenID, saleValue, imageUrl, active }: PastAuctionProps) => {
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
    )
}

export default PastAuction