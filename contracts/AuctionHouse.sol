// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./interfaces/IAuctionHouse.sol";
import "./interfaces/IToken.sol";

contract AuctionHouse is Ownable, IAuctionHouse {
    IToken public token;
    uint256 public duration;
    IAuctionHouse.Auction auction;

    constructor(address _token, uint256 _duration) {
        token = IToken(_token);
        duration = _duration;
    }

    function startAuction(uint256 tokenId) external onlyOwner {
        address tokenOwner = token.ownerOf(tokenId);
        require(tokenOwner == address(this), "Auction house does not own desired token");
        uint256 startDate = block.timestamp;
        uint256 endDate = startDate + duration;
        auction = IAuctionHouse.Auction({
            startDate: startDate,
            endDate: endDate,
            highestBid: 0,
            highestBidder: address(0),
            tokenId: tokenId,
            settled: false
        });

        emit AuctionStarted(tokenId, startDate, endDate);
    }

    function endAuction() external onlyOwner {
        IAuctionHouse.Auction memory _auction = auction;
        require(_auction.startDate != 0, "Auction hasn't started");
        require(!_auction.settled, "Auction has already been settled");
        require(block.timestamp > _auction.endDate, "Auction cannot be ended yet");
        auction.settled = true;
        if (auction.highestBidder == address(0)) {
            token.burn(_auction.tokenId);
        } else {
            token.transferFrom(address(this), _auction.highestBidder, _auction.tokenId);
        }

        emit AuctionEnded(_auction.highestBidder);
    }

    function placeBid(uint256 tokenId) external payable {
        IAuctionHouse.Auction memory _auction = auction;
        require(tokenId == _auction.tokenId, "Token not up for auction");
        require(msg.value > _auction.highestBid, "Value not enough to outbid");
        require(msg.sender != _auction.highestBidder, "Already have the highest bid");
        require(block.timestamp < _auction.endDate, "Auction has expired");
        require(!_auction.settled, "Auction has already been settled");
        auction.highestBid = msg.value;
        auction.highestBidder = msg.sender;
        emit BidPlaced(msg.sender, msg.value);
    }
}