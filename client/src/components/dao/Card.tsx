import '../../styles/proposal-card.scss';

type CardProps = {
    proposalID: number;
    proposedDate: string;
    active: boolean;
};

const Card = ({ proposalID, proposedDate, active }: CardProps) => {
    return (
        <a href={`/dao/${proposalID}`}>
            <div className={'proposal' + (active ? ' active' : '')}>
                <div className='description'>
                    <b>Proposal #{proposalID}</b>
                    <p>
                        Proposed on {proposedDate} and is active: {active}
                    </p>
                </div>
            </div>
        </a>
    );
};

export default Card;
