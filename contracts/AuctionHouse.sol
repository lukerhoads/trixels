// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IAuctionHouse.sol";
import "./interfaces/IToken.sol";
import "./interfaces/IDAO.sol";
import "./interfaces/IDistributor.sol";
import "./utility/ETHMover.sol";

// AuctionHouse auctions off the latest Trixel
contract AuctionHouse is Ownable, IAuctionHouse, ETHMover {
    IToken public token;
    IDAO public dao;
    IDistributor public distributor;
    IAuctionHouse.Auction auction;

    // The duration of every auction
    uint256 public duration;

    // The reserve price of every auction
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
     * @notice enables the owner to start a new auction. 
     * @dev only callable by the owner
     * @param _tokenID  the recipient of the proposal transaction
     */
    function startAuction() external onlyOwner {
        IAuctionHouse.Auction memory _auction = auction;
        require(_auction.settled, "Another auction has already been started");
        uint tokenID = token.mint(address(this));
        uint256 startDate = block.timestamp;
        uint256 endDate = startDate + duration;
        auction = IAuctionHouse.Auction({
            startDate: startDate,
            endDate: endDate,
            highestBid: 0,
            highestBidder: payable(0),
            tokenId: tokenID,
            settled: false
        });

        emit AuctionStarted(tokenID, startDate, endDate);
    }

    /*
     * @notice enables the owner to end an active auction.
     * @dev only callable by the owner
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
            distributor.deposit{ value: _auction.highestBid }(_auction.tokenId, _auction.highestBid);
            require(_safeTransferETH(address(dao), _auction.highestBid), "Could not transfer to DAO contract");
        }

        emit AuctionEnded(_auction.highestBidder, _auction.highestBid);
    }

    /*
     * @notice enables anyone to bid on an active auction.
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

    /*
     * @notice sets the auction duration to a new duration.
     * @dev only callable by the owner
     * @param _duration the new duration
     */
    function setDuration(uint256 _duration) external onlyOwner {
        duration = _duration;
    }

    /*
     * @notice sets the auction duration to a new duration.
     * @dev only callable by the owner
     * @param _reservePrice the new reserve price 
     */
    function setReservePrice(uint256 _reservePrice) external onlyOwner {
        reservePrice = _reservePrice;
    }

    /*
     * @notice sets the minimum bid increment percentage.
     * @dev only callable by the owner
     * @param _minBidIncrementPercentage the new bid increment percentage
     */
    function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) external onlyOwner {
        minBidIncrementPercentage = _minBidIncrementPercentage;
    }
}