// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

interface IToken is IERC721 {
    function setAuctionHouse(address _auctionHouse) external;

    function mint(address to) external returns (uint tokenID);

    function burn(uint256 tokenId) external;
}