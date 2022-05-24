// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "./interfaces/IToken.sol";
import "./interfaces/IDAO.sol";

contract DAO is IDAO {
    uint constant minProposalVotingPeriod = 2 weeks;

    IToken public token;
    IDAO.Proposal[] public proposals;

    constructor(address _token) {
        token = IToken(_token);
    }

    function isMember(address person) public view returns (bool) {
        return token.balanceOf(person) > 0;
    }

    modifier onlyTokenHolders() {
        require(isMember(msg.sender), "Not part of DAO");
        _;
    }

    receive() external payable {
    }

    function makeProposal(address _recipient, uint256 _amount, string calldata _description) external onlyTokenHolders returns (uint proposalID) {
        proposalID = proposals.length++;
        IDAO.Proposal storage p = proposals[proposalID];
        p.recipient = _recipient;
        p.amount = _amount;
        p.description = _description;
        p.passed = false; 
        p.creator = msg.sender;
        p.createdAt = block.timestamp;
        emit ProposalAdded(_recipient, _amount, _description);
    }

    function vote(uint _proposalID, bool _inSupport) external {
        IDAO.Proposal storage p = proposals[_proposalID];
        if (_inSupport) {
            p.yay++;
            p.votedFor[msg.sender] = true;
        } else {
            p.nay++;
            p.votedAgainst[msg.sender] = true;
        }

        emit Voted(_proposalID, _inSupport, msg.sender);
    }

    function executeProposal(uint _proposalID, bytes memory _transactionData) external {
        IDAO.Proposal storage p = proposals[_proposalID];
        require(p.createdAt + minProposalVotingPeriod < block.timestamp, "Cannot execute proposal, voting period has not ended");
        require(p.yay > p.nay, "Yay votes do not exceed Nay votes");

        emit ProposalExecuted(_proposalID);
    }
}