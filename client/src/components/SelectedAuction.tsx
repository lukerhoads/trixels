import { observer } from 'mobx-react';
import { useEffect } from 'react';

// - Trixel ID (complete)
// - Sale price (source from distributor contract)
// - Auction winner (event in auctionHouse contract)
// - Current owner (source from token contract)

type SelectedAuctionProps = {
  tokenID: string;
};

const SelectedAuction = ({ tokenID }: SelectedAuctionProps) => {
  useEffect(() => {
    // Verify tokenID by checking token contract and making sure tokenID is less than tokenQuantity
  }, []);

  return (
    <div className='active-auction'>
      <div className='auc-wrap'>
        <img className='active-image' src={'https://cdn.britannica.com/91/181391-050-1DA18304/cat-toes-paw-number-paws-tiger-tabby.jpg?q=60'} />
        <div className='active-data'>
          <div className='section'>
            <p className='caption'>Trixel ID:</p>
            <p className='value'>{tokenID}</p>
            <p className='caption'>Sale price:</p>
            <p className='value'></p>
          </div>
          <div className='section'>
            <p className='caption'>Sale date:</p>
            <p className='value'></p>
          </div>
        </div>
        <p className='caption'>Auction winner:</p>
        <p className='value'></p>
        <p className='caption'>Current owner:</p>
        <p className='value'></p>
      </div>
    </div>
  );
};

export default observer(SelectedAuction);
