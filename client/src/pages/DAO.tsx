import { useEthers } from '@usedapp/core';
import { Contract, utils } from 'ethers';
import Header from 'layout/header';
import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import Card from '../components/dao/Card';
import config from '../config';
import '../styles/dao.scss';

const { addresses } = config;

import Haptics from 'components/Haptics';
import Overlay from 'components/Overlay';
import { observer } from 'mobx-react';
import store from 'store';
import { DaoStats, Proposal } from 'types/dao';
import DAOABI from '../abi/contracts/DAO.sol/DAO.json';

const daoInterface = new utils.Interface(DAOABI);

const Dao = () => {
    const { account } = useEthers();
    const { proposalID } = useParams();
    const [loading, setLoading] = useState(false);
    const [proposals, setProposals] = useState<Proposal[]>([]);
    const [daoStats, setDaoStats] = useState<DaoStats | undefined>(undefined);
    const [canExecuteProposal, setCanExecuteProposal] = useState(false);
    const { library } = useEthers();
    const voteCheckboxRef = useRef<HTMLInputElement>(null);
    const [userIsMember, setUserIsMember] = useState(false)
    const [userVotedForActive, setUserVotedForActive] = useState(false);
    const [validProposalID, setValidProposalID] = useState(false);

    const daoContract = useMemo((): Contract | undefined => {
        if (!library || !addresses.auctionHouse) return;
        return new Contract(addresses.dao, daoInterface, library);
    }, [library]);

    useEffect(() => {
        if (proposalID) {
            fetchCanExecuteProposal();
        }
        fetchUserIsMember()
        fetchStats();
        fetchProposals();
    }, [daoContract]);

    useEffect(() => {
        validateProposalId();
        calculateIfVotedFor();
    }, [account]);

    const fetchUserIsMember = async () => {
        if (!daoContract) return
        const isMember = await daoContract.isMember(account)
        setUserIsMember(isMember)
    }

    const validateProposalId = async () => {
        if (!daoContract || !proposalID) return;
        const numProposals = await daoContract.numProposals();
        setValidProposalID(parseInt(proposalID) < numProposals);
    };

    const calculateIfVotedFor = async () => {
        if (!daoContract || !proposalID) return;
        const proposal = await daoContract.proposals(proposalID);
        console.log('proposal: ', proposal);
    };

    const fetchProposals = async () => {
        if (!daoContract) return;
        const numProposals = await daoContract.numProposals();
        let proposals: Proposal[] = [];
        for (let i = numProposals; i--; i > 0) {
            console.log('Proposal #' + i);
            let proposal = await daoContract.proposals(i);
            console.log('proposal: ', proposal);
            proposals.push({
                proposalID: i,
                ...proposal,
            });
        }

        setProposals(proposals);
        setLoading(false);
    };

    const fetchStats = async () => {
        if (!daoContract) return;
        const numProposals = await daoContract.numProposals();
        const balance = await library?.getBalance(daoContract.address)
        setDaoStats({
            numProposals: numProposals.toNumber(),
            votingPeriod: '2 weeks',
            etherBalance: balance
        });
    };

    const makeProposalClick = () => {
        if (!account) {
            store.setOverlayInfo({
                type: 'auth-modal',
            });
        } else {
            store.setOverlayInfo({
                type: 'proposal-modal',
            });
        }
    };

    const executeProposalClick = () => {
        if (!account) {
            store.setOverlayInfo({
                type: 'auth-modal',
            });
        } else {
            store.setOverlayInfo({
                type: 'execute-proposal-modal',
            });
        }
    };

    const voteClick = () => {
        if (!voteCheckboxRef || !voteCheckboxRef.current) return;
        vote(voteCheckboxRef.current.checked);
    };

    const unVoteClick = () => {
        unVote()
    };

    const fetchCanExecuteProposal = async () => {
        if (!daoContract || !proposalID) return;
        let canExecuteProposal = await daoContract.canExecuteProposal(proposalID);
        setCanExecuteProposal(canExecuteProposal);
    };

    const vote = useCallback(
        async (inSupport: boolean) => {
            if (!daoContract || !library || !proposalID) return;
            const voteTx = await daoContract.vote(proposalID, inSupport);
            await voteTx.wait();
        },
        [daoContract]
    );

    const unVote = useCallback(async () => {
        if (!daoContract || !library || !proposalID) return;
        const voteTx = await daoContract.unVote(proposalID);
        await voteTx.wait();
    }, [daoContract]);

    return (
        <Haptics type="logs">
            <div className='dao'>
                <Overlay />
                <Header>
                    <Link to='/'>Home</Link>
                    <Link to='/auction/current'>Auction</Link>
                </Header>
                <div className='dao-wrapper'>
                    <div className='dao-top-section'>
                        <div className='dao-stats'>
                            <p>Trixels DAO Stats</p>
                            <p className='caption'># of Proposals:</p>
                            <p className='value'>{daoStats?.numProposals}</p>
                            <p className='caption'>Voting period:</p>
                            <p className='value'>{daoStats?.votingPeriod}</p>
                            <p className='caption'>Balance:</p>
                            <p className='value'>{daoStats?.etherBalance ? utils.formatEther(daoStats.etherBalance) : "0"} ETH</p>
                            {userIsMember && <button onClick={() => makeProposalClick()}>Make proposal</button>}
                        </div>
                        <div className='dao-spacer' />
                        {proposalID && userIsMember && (
                            <div className='active-proposal'>
                                Active proposal
                                <input type='checkbox' ref={voteCheckboxRef}>
                                    Approve?
                                </input>
                                <button onClick={voteClick}>Vote</button>
                                <button onClick={unVoteClick}>Unvote</button>
                                <button onClick={executeProposalClick}>Execute proposal</button>
                            </div>
                        )}
                    </div>

                    <div className='proposals'>
                        {proposals.map((proposal, idx) => (
                            <Card key={idx} proposalID={proposal.proposalID} proposedDate={proposal.createdAt.toString()} active={!proposal.passed} />
                        ))}
                    </div>
                </div>
            </div>
        </Haptics>
    );
};

export default observer(Dao);
