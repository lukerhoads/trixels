// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface IAuctionHouse {
    struct Auction {
        uint startDate;
        uint endDate;
        uint256 highestBid;
        address highestBidder;
        uint256 tokenId;
        bool settled;
    }

    function startAuction(uint256 tokenId) external;

    function endAuction() external;

    function placeBid() external;

    event AuctionStarted(uint256 tokenId, uint startDate, uint endDate);

    event AuctionEnded(address winner);

    event BidPlaced(address bidder, uint256 amount);
}