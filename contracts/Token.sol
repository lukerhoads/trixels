// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./interfaces/IToken.sol";

// Token is the underlying Trixels NFT contract that represents the sold NFTs
contract Token is Ownable, ERC721, IToken {
    uint public tokenQuantity;

    address public auctionHouse;

    constructor() ERC721("Trixels", "TRIX") {
        tokenQuantity = 0;
    }

    modifier onlyAuctionHouse() {
        require(msg.sender == auctionHouse, "Sender not auction house");
        _;
    }

    function setAuctionHouse(address _auctionHouse) external override onlyOwner {
        auctionHouse = _auctionHouse;
    }

    /*
     * mint mints a new token
     * @return tokenID the ID of the minted token
     */
    function mint(address to) external override onlyAuctionHouse returns (uint) {
        tokenQuantity++;
        _safeMint(to, tokenQuantity);
        return tokenQuantity;
    }

    /*
     * burn burns a token
     * @param tokenID the ID of the token to burn
     */
    function burn(uint256 tokenId) external override onlyAuctionHouse {
        tokenQuantity--;
        return _burn(tokenId);
    }

    function _baseURI() internal pure override returns (string memory) {
        return ""; // TODO: insert redirect URL for Skynet metadata URLs
    }
}