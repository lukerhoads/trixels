import '../../styles/auction-card.scss';

type CardProps = {
    tokenID: number;
    mintDate: string;
    imageUrl: string;
    active: boolean;
};

const Card = ({ tokenID, mintDate, imageUrl, active }: CardProps) => {
    return (
        <a href={`/auction/${tokenID}`}>
            <div className={'past-auction' + (active ? ' active' : '')}>
                <img className='thumbnail' src={imageUrl} />
                <div className='description'>
                    <b>Trixel #{tokenID}</b>
                    <p>Minted on {mintDate}</p>
                </div>
            </div>
        </a>
    );
};

export default Card;
