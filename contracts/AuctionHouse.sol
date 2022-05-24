// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IAuctionHouse.sol";
import "./interfaces/IToken.sol";
import "./interfaces/IDAO.sol";
import "./utility/ETHMover.sol";

contract AuctionHouse is Ownable, IAuctionHouse, ETHMover {
    IToken public token;
    IDAO public dao;
    IAuctionHouse.Auction auction;

    uint256 public duration;
    uint256 public reservePrice;
    uint8 public minBidIncrementPercentage;
    uint8 public daoCut;

    constructor(address _token, address _dao, uint256 _duration, uint8 _minBidIncrementPercentage, uint8 _daoCut, address _weth) ETHMover(_weth) {
        token = IToken(_token);
        dao = IDAO(_dao);
        duration = _duration;
        minBidIncrementPercentage = _minBidIncrementPercentage;
    }

    /*
     * startAuction enables the owner to start a new auction. 
     * @param _tokenID  the recipient of the proposal transaction
     */
    function startAuction(uint256 _tokenID) external onlyOwner {
        IAuctionHouse.Auction memory _auction = auction;
        require(_auction.settled, "Another auction has already been started");
        address tokenOwner = token.ownerOf(_tokenID);
        require(tokenOwner == address(this), "Auction house does not own desired token");

        uint256 startDate = block.timestamp;
        uint256 endDate = startDate + duration;
        auction = IAuctionHouse.Auction({
            startDate: startDate,
            endDate: endDate,
            highestBid: 0,
            highestBidder: payable(0),
            tokenId: _tokenID,
            settled: false
        });

        emit AuctionStarted(_tokenID, startDate, endDate);
    }

    /*
     * endAuction enables the owner to end an active auction.
     */
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

        if (_auction.highestBid > 0) {
            // Send paycut to DAO
            require(_safeTransferETH(address(dao), _auction.highestBid), "Could not transfer to DAO contract");
        }

        emit AuctionEnded(_auction.highestBidder, _auction.highestBid);
    }

    /*
     * placeBid enables anyone to bid on an active auction.
     * @param _tokenID  the recipient of the proposal transaction
     */
    function placeBid(uint256 _tokenID) external payable {
        IAuctionHouse.Auction memory _auction = auction;
        require(_tokenID == _auction.tokenId, "Token not up for auction");
        require(msg.value > reservePrice, "Bid does not meet reserve price");
        require(
            msg.value >= _auction.highestBid + ((_auction.highestBid * minBidIncrementPercentage) / 100), 
            "Bid amount insufficient"
        );
        require(msg.sender != _auction.highestBidder, "Already have the highest bid");
        require(block.timestamp < _auction.endDate, "Auction has expired");
        require(!_auction.settled, "Auction has already been settled");

        address payable lastBidder = _auction.highestBidder;
        if (_auction.highestBidder != address(0)) {
            // Transfers cut to DAO
            _safeTransferETHWithFallback(lastBidder, (daoCut * _auction.highestBid) / 100);
        }

        auction.highestBid = msg.value;
        auction.highestBidder = payable(msg.sender);
        emit BidPlaced(msg.sender, msg.value);
    }

    function setReservePrice(uint256 _reservePrice) external onlyOwner {
        reservePrice = _reservePrice;
    }

    function setMinIncrementPercentage(uint256 _reservePrice) external onlyOwner {
        reservePrice = _reservePrice;
    }
}