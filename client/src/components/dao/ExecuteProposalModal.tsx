import { useEthers } from '@usedapp/core';
import { Contract, utils } from 'ethers';
import React, { useMemo, useRef, useState } from 'react';
import { useParams } from 'react-router-dom';
import DAOABI from '../../abi/contracts/DAO.sol/DAO.json';
import config from '../../config';
import '../../styles/proposal-modal.scss';
import Close from '../Close';

const daoInterface = new utils.Interface(DAOABI);
const { addresses } = config;

export type ExecuteProposalModalProps = {
    onExit: (e: React.MouseEvent<HTMLDivElement>) => void;
};

const ExecuteProposalModal = ({ onExit }: ExecuteProposalModalProps) => {
    const { library } = useEthers();
    const { proposalID } = useParams();
    const txDataRef = useRef<HTMLInputElement>(null);
    const [error, setError] = useState<string | undefined>(undefined);

    const daoContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.dao, daoInterface, library);
    }, [library]);

    const executeProposal = async (proposalID: string, transactionData: string) => {
        if (!daoContract || !library) return;
        try {
            const daoConnected = await daoContract.connect(library.getSigner());
            const makeProposalTx = await daoConnected.executeProposal(proposalID, transactionData);
            await makeProposalTx.wait();
        } catch (err: any) {
            setError(err.data.data.message);
        }
    };

    const submitClick = () => {
        if (!proposalID || !txDataRef || !txDataRef.current?.value) return;
        executeProposal(proposalID, txDataRef.current?.value);
    };

    const onDivClick = (e: React.MouseEvent<HTMLDivElement>) => {
        e.stopPropagation();
    };

    return (
        <div className='proposal-modal' onClick={onDivClick}>
            <Close onExit={onExit} />
            <div className='proposal-modal-wrap'>
                <p>Execute Proposal</p>
                <input placeholder='Transaction data' ref={txDataRef} />
                <button onClick={submitClick}>Submit</button>
                {error && <p className='error'>{error}</p>}
            </div>
        </div>
    );
};

export default ExecuteProposalModal;
