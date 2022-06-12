export type Proposal = {
    proposalID: number
    proposalHash: string 
    recipient: string 
    amount: number 
    description: string 
    passed: boolean 
    yay: number 
    nay: number 
    creator: string 
    createdAt: number
}