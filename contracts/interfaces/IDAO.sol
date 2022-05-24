// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface IDAO {
    struct Proposal {
        address recipient;
        uint amount;
        string description;
        bool passed;
        mapping (address => bool) votedFor;
        mapping (address => bool) votedAgainst;
        uint yay;
        uint nay;
        address creator;
        uint createdAt;
    }

    function makeProposal(address _recipient, uint _amount, string calldata _description) external returns (uint proposalID);

    function vote(uint _proposalID, bool _inSupport) external;

    event ProposalAdded(address recipient, uint amount, string description);

    event Voted(uint indexed proposalID, bool inSupport, address voter);

    event ProposalExecuted(uint indexed proposalID);
}