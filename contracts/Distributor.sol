// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IDistributor.sol";
import "./utility/ETHMover.sol";

contract Distributor is Ownable, ETHMover {
    uint constant TOTAL_CONTRIBS = 900; // 30 * 30

    mapping(uint256 => uint256) public sales;
    mapping(address => uint256) public balances;

    constructor(address _weth) ETHMover(_weth) {
    }

    /*
     * deposit allows the AuctionHouse to deposit the distribution
     *         for contributors
     * @param _tokenID the token ID of the sale
     */
    function deposit(uint256 _tokenID) external payable {
        sales[_tokenID] = msg.value;
    }

    /*
     * distribute is the action performed by the server in order 
     *            to distribute the sale price among contributors
     * @param contributors the contributors of the sold token
     * @param _tokenID the token ID of the distributing token
     */
    function distribute(IDistributor.Contributor[] memory _contributors, uint _tokenID) external onlyOwner {
        for (uint i = 0; i < _contributors.length; i++) {
            balances[_contributors[i].addr] += (_contributors[i].numContribs / TOTAL_CONTRIBS) * sales[_tokenID];
        }
    }

    /*
     * withdraw allows contributors to withdraw their cumulative earnings
     */
    function withdraw() external {
        require(balances[msg.sender] > 0, "No balance to withdraw");
        _safeTransferETHWithFallback(msg.sender, balances[msg.sender]);
    }
}