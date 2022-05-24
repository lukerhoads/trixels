// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface IDAO {
    struct Proposal {
        bytes32 proposalHash;
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

    function makeProposal(address _recipient, uint _amount, string calldata _description, bytes calldata _transactionData) external returns (uint proposalID);

    function vote(uint _proposalID, bool _inSupport) external;

    function unVote(uint _proposalID) external;

    function executeProposal(uint proposalID, bytes calldata transactionData) external returns (bool success);

    function proposalCount() external view returns (uint);

    event ProposalAdded(uint indexed proposalID, address recipient, uint amount, string description);

    event Voted(uint indexed proposalID, bool inSupport, address indexed voter);

    event ProposalExecuted(uint indexed proposalID, bool result);
}