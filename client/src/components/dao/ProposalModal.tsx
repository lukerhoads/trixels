import { useEthers } from '@usedapp/core';
import { Contract, utils } from 'ethers';
import React, { useMemo, useRef, useState } from 'react';
import { ProposalSubmit } from 'types/dao';
import DAOABI from '../../abi/contracts/DAO.sol/DAO.json';
import config from '../../config';
import '../../styles/proposal-modal.scss';
import Close from '../Close';

const daoInterface = new utils.Interface(DAOABI);
const { addresses } = config;

export type ProposalModalProps = {
    onExit: (e: React.MouseEvent<HTMLDivElement>) => void;
};

const ProposalModal = ({ onExit }: ProposalModalProps) => {
    const { library } = useEthers();
    const [error, setError] = useState<string | undefined>(undefined);
    const recipientRef = useRef<HTMLInputElement>(null);
    const amountRef = useRef<HTMLInputElement>(null);
    const descriptionRef = useRef<HTMLInputElement>(null);
    const txDataRef = useRef<HTMLInputElement>(null);

    const daoContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.dao, daoInterface, library);
    }, [library]);

    const makeProposal = async (proposal: ProposalSubmit) => {
        if (!daoContract || !library) return;
        try {
            const daoConnected = daoContract.connect(library.getSigner());
            const makeProposalTx = await daoConnected.makeProposal(
                proposal.recipient,
                proposal.amount,
                proposal.description,
                utils.formatBytes32String(proposal.transactionData || '')
            );
            await makeProposalTx.wait();
        } catch (err: any) {
            setError(err.data.data.message);
        }
    };

    const submitClick = () => {
        if (!recipientRef || !amountRef || !descriptionRef || !txDataRef) return;
        let proposal = {
            recipient: recipientRef.current?.value,
            amount: amountRef.current?.value ? parseInt(amountRef.current?.value) : undefined,
            description: descriptionRef.current?.value,
            transactionData: txDataRef.current?.value,
        };
        makeProposal(proposal);
    };

    const onDivClick = (e: React.MouseEvent<HTMLDivElement>) => {
        e.stopPropagation();
    };

    return (
        <div className='proposal-modal' onClick={onDivClick}>
            <Close onExit={onExit} />
            <div className='proposal-modal-wrap'>
                <p>Make Proposal</p>
                <input placeholder='Recipient' ref={recipientRef} />
                <input placeholder='Amount' ref={amountRef} />
                <input placeholder='Description' ref={descriptionRef} />
                <input placeholder='Transaction data' ref={txDataRef} />
                <button onClick={submitClick}>Submit</button>
                {error && <p className='error'>{error}</p>}
            </div>
        </div>
    );
};

export default ProposalModal;
