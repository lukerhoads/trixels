// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./interfaces/IToken.sol";

contract Token is Ownable, ERC721, IToken {
    constructor() ERC721("Trixels", "TRIX") {}

    function mint(address to, uint256 tokenId) external onlyOwner {
        return _safeMint(to, tokenId);
    }

    function burn(uint256 tokenId) external onlyOwner {
        return _burn(tokenId);
    }

    function _baseURI() internal pure override returns (string memory) {
        return ""; // TODO: insert redirect URL for Skynet metadata URLs
    }
}