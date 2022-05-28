// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "./interfaces/IToken.sol";
import "./interfaces/IDAO.sol";
import "./utility/ETHMover.sol";

contract DAO is IDAO, ETHMover {
    uint constant minProposalVotingPeriod = 2 weeks;

    IToken public token;
    IDAO.Proposal[] public proposals;

    constructor(address _token, address _weth) ETHMover(_weth) {
        token = IToken(_token);
    }

    function isMember(address person) public view returns (bool) {
        return token.balanceOf(person) > 0;
    }

    function proposalCount() external override view returns (uint) {
        return proposals.length;
    }

    /*
     * @notice only allows users who hold a Trixel. 
     */
    modifier onlyTokenHolders() {
        require(isMember(msg.sender), "Not part of DAO");
        _;
    }

    receive() external payable {
    }

    /*
     * @notice enables DAO members to make a proposal to the DAO. 
     * @dev only allows DAO members
     * @param _recipient  the recipient of the proposal transaction
     * @param _amount the amount to be sent on approval
     * @param _description description of the proposal
     * @return proposalID the ID assigned to the proposal
     */
    function makeProposal(address _recipient, uint256 _amount, string calldata _description, bytes calldata _transactionData) external override onlyTokenHolders returns (uint proposalID) {
        proposalID = proposals.length + 1;
        IDAO.Proposal storage p = proposals[proposalID];
        p.proposalHash = keccak256(abi.encodePacked(_recipient, _amount, _transactionData));
        p.recipient = _recipient;
        p.amount = _amount;
        p.description = _description;
        p.passed = false; 
        p.creator = msg.sender;
        p.createdAt = block.timestamp;
        emit ProposalAdded(proposalID, _recipient, _amount, _description);
    }

    /*
     * @notice enables DAO members to vote on an active proposal. 
     * @dev only allows DAO members
     * @param _proposalID  the ID of the proposal to vote on
     * @param _inSupport whether the member supports the proposal
     */
    function vote(uint _proposalID, bool _inSupport) external override onlyTokenHolders {
        IDAO.Proposal storage p = proposals[_proposalID];
        if (_inSupport) {
            p.yay += token.balanceOf(msg.sender);
            p.votedFor[msg.sender] = true;
        } else {
            p.nay += token.balanceOf(msg.sender);
            p.votedAgainst[msg.sender] = true;
        }

        emit Voted(_proposalID, _inSupport, msg.sender);
    }

    /*
     * @notice enables DAO members to unvote on an active proposal.
     * @dev only allows DAO members 
     * @param _proposalID  the ID of the proposal to vote on
     */
    function unVote(uint _proposalID) external override onlyTokenHolders {
        IDAO.Proposal storage p = proposals[_proposalID];
        if (p.votedFor[msg.sender]) {
            p.yay -= token.balanceOf(msg.sender);
            p.votedAgainst[msg.sender] = false;
        } 

        if (p.votedAgainst[msg.sender]) {
            p.nay -= token.balanceOf(msg.sender);
            p.votedAgainst[msg.sender] = false;
        }
    }

    /*
     * @notice enables DAO members to execute a proposal if approved. 
     * @dev only allows DAO members
     * @param _proposalID  the ID of the proposal to pass
     * @param _transactionData the transactionData previously entered
     * @return success whether the proposal could be successfully executed
     */
    function executeProposal(uint _proposalID, bytes calldata _transactionData) external override onlyTokenHolders returns (bool success) {
        IDAO.Proposal storage p = proposals[_proposalID];
        require(p.createdAt + minProposalVotingPeriod < block.timestamp, "Cannot execute proposal, voting period has not ended");
        require(p.yay > p.nay, "Yay votes do not exceed Nay votes");
        require(!p.passed, "Proposal has already been passed");
        require(checkProposalHash(_proposalID, p.recipient, p.amount, _transactionData), "Invalid transaction data");
        success = _safeTransferETHWithData(p.recipient, p.amount, _transactionData);
        emit ProposalExecuted(_proposalID, success);
    }

    /*
     * @notice enables DAO members to check a proposals hash. 
     * @param _proposalID  the ID of the proposal to check
     * @param _recipient the address of the proposal recipient
     * @param _amount the amount of the proposal
     * @param _transactionData the transaction data being verified
     * @return validHash whether the hash is valid
     */
    function checkProposalHash(uint _proposalID, address _recipient, uint256 _amount, bytes calldata _transactionData) public view returns (bool validHash) {
        IDAO.Proposal storage p = proposals[_proposalID];
        return p.proposalHash == keccak256(abi.encodePacked(_recipient, _amount, _transactionData));
    }
}