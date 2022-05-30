// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IDistributor.sol";
import "./utility/ETHMover.sol";

// Distributor is the contract that distributes the partitioned sale price to contributors.
contract Distributor is Ownable, ETHMover {
    uint constant TOTAL_CONTRIBS = 900; // 30 * 30

    mapping(uint => uint256) public sales;
    mapping(address => uint256) public balances;

    constructor(address _weth) ETHMover(_weth) {
    }

    /*
     * @notice allows the AuctionHouse to deposit the distribution
     *         for contributors
     * @param _tokenID the token ID of the sale
     */
    function recordSale(uint _tokenID) external payable {
        sales[_tokenID] = msg.value;
    }

    /*
     * @notice is the action performed by the server in order 
     *            to distribute the sale price among contributors
     * @dev only allows owner
     * @param contributors the contributors of the sold token
     * @param _tokenID the token ID of the distributing token
     */
    function distribute(IDistributor.Contributor[] memory _contributors, uint _tokenID) external onlyOwner {
        for (uint i = 0; i < _contributors.length; i++) {
            balances[_contributors[i].addr] += (_contributors[i].numContribs / TOTAL_CONTRIBS) * sales[_tokenID];
        }
    }

    /*
     * @notice allows contributors to withdraw their cumulative earnings
     */
    function withdraw() external {
        require(balances[msg.sender] > 0, "No balance to withdraw");
        _safeTransferETHWithFallback(msg.sender, balances[msg.sender]);
    }
}