// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface IAuctionHouse {
    struct Auction {
        uint startDate;
        uint endDate;
        uint256 highestBid;
        address payable highestBidder;
        uint256 tokenId;
        bool settled;
    }

    function startAuction() external returns (uint);

    function endAuction() external;

    function placeBid(uint _tokenId) external payable;

    event AuctionStarted(uint indexed tokenId, uint startDate, uint endDate);

    event AuctionEnded(uint indexed tokenId, address winner, uint256 amount);

    event BidPlaced(uint indexed tokenId, address bidder, uint256 amount);
}