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

    function deposit(uint256 _tokenId) external payable {
        sales[_tokenId] = msg.value;
    }

    function distribute(IDistributor.Contributor[] memory _contributors, uint256 _salePrice) external onlyOwner {
        for (uint i = 0; i < _contributors.length; i++) {
            balances[_contributors[i].addr] += (_contributors[i].numContribs / TOTAL_CONTRIBS) * _salePrice;
        }
    }

    function withdraw() external {
        require(balances[msg.sender] > 0, "No balance to withdraw");
        _safeTransferETHWithFallback(msg.sender, balances[msg.sender]);
    }
}