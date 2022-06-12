import { BigNumber } from "ethers";

export type Proposal = {
    proposalID: number;
    proposalHash?: string;
    recipient?: string;
    amount?: number;
    description?: string;
    passed?: boolean;
    yay?: number;
    nay?: number;
    creator?: string;
    createdAt: number;
    transactionData?: string;
};

export type ProposalSubmit = {
    recipient?: string;
    amount?: number;
    description?: string;
    transactionData?: string;
};

export type DaoStats = {
    numProposals: number;
    votingPeriod: string;
    etherBalance?: BigNumber
};
