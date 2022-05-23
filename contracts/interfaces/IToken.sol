// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

interface IToken is IERC721 {
    function mint(address to, uint256 tokenId) external;

    function burn(uint256 tokenId) external;
}