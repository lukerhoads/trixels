import { useEthers } from '@usedapp/core';
import { Contract, utils } from 'ethers';
import Header from 'layout/header';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import Card from '../components/dao/Card';
import config from '../config';

const { addresses } = config;

import { Proposal } from 'types/dao';
import DAOABI from '../abi/contracts/DAO.sol/DAO.json';

const daoInterface = new utils.Interface(DAOABI);

const Dao = () => {
    const { proposalID } = useParams()
    const [loading, setLoading] = useState(false)
    const [proposals, setProposals] = useState<Proposal[]>([]);
    const { library } = useEthers();

    const daoContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.dao, daoInterface, library);
    }, [library]);

    useEffect(() => {
      listProposals()
    }, [daoContract])

    const listProposals = async () => {
        if (!daoContract) return
        const numProposals = await daoContract.numProposals();
        let proposals: Proposal[] = [];
        for (let i=numProposals - 20; i++; i<numProposals) {
            let proposal = await daoContract.proposals(i)
            console.log("proposal: ", proposal)
            proposals.push({
                proposalID: i,
                ...proposal,
            })
        }

        setProposals(proposals)
    };

    const vote = useCallback(async () => {
        if (!daoContract || !library || !proposalID) return;
        const daoConnected = daoContract.connect(library.getSigner());
    }, [daoContract]);

    return (
        <div className='dao'>
            <Header>
                <Link to='/'>Home</Link>
            </Header>
            <div className='dao-wrapper'>
                { proposalID != '' && (
                    <p>Active proposal</p>
                )}
                <div className='dao-spacer' />
                <div className='proposals'>
                    {proposals.map((proposal, idx) => (
                        <Card key={idx} proposalID={proposal.proposalID} proposedDate={proposal.createdAt.toString()} active={!proposal.passed} />
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Dao;
