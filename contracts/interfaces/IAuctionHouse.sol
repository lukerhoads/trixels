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

    function placeBid(uint256 _tokenId) external payable;

    event AuctionStarted(uint256 tokenId, uint startDate, uint endDate);

    event AuctionEnded(address winner, uint256 amount);

    event BidPlaced(address bidder, uint256 amount);
}