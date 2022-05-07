// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

// Most code copied from NounsDAO
// https://github.com/nounsDAO/nouns-monorepo/blob/master/packages/nouns-contracts/contracts/NounsAuctionHouse.sol

import { PausableUpgradeable } from '@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol';
import { ReentrancyGuardUpgradeable } from '@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol';
import { OwnableUpgradeable } from '@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol';

interface ITrixelsAuctionHouse {
    struct Auction {
        uint256 trixelId;
        uint256 amount;
        uint256 startTime;
        uint256 endTime;
        address payable bidder;
        bool settled;
    }

    event AuctionCreated(uint256 indexed trixelId, uint256 startTime, uint256 endTime);

    event AuctionBid(uint256 indexed trixelId, address sender, uint256 value, bool extended);

    event AuctionExtended(uint256 indexed nounId, uint256 endTime);

    event AuctionSettled(uint256 indexed trixelId, address winner, uint256 amount);
}

contract TrixelsAuctionHouse is ITrixelsAuctionHouse, PausableUpgradeable, ReentrancyGuardUpgradeable, OwnableUpgradeable {
    uint256 public duration;
    uint256 public timeBuffer;
    uint8 public minBidIncrementPercentage;
    uint256 public reservePrice;

    ITrixelsAuctionHouse.Auction auction;
    TrixelsNFT trixels;

    constructor(TrixelsNFT _trixels, uint256 _duration, uint256 _timeBuffer, uint256 _reservePrice) {
        __Pausable_init();
        __ReentrancyGuard_init();
        __Ownable_init();
        _pause();
        trixels = _trixels;
        timeBuffer = _timeBuffer;
        duration = _duration;
        reservePrice = _reservePrice;
    }

    function startAuction(bytes64 skyNetID) external nonReentrant whenNotPaused {
        _startAuction(skyNetID);
    }

    function settleAuction() external nonReentrant whenNotPaused {
        _settleAuction();
    }

    function createBid(uint256 trixelId) external payable nonReentrant {
        INounsAuctionHouse.Auction memory _auction = auction;

        require(_auction.trixelId == trixelId, 'Trixel not up for auction');
        require(block.timestamp < _auction.endTime, 'Auction expired');
        require(msg.value >= reservePrice, 'Must send at least reservePrice');
        require(
            msg.value >= _auction.amount + ((_auction.amount * minBidIncrementPercentage) / 100),
            'Must send more than last bid by minBidIncrementPercentage amount'
        );

        address payable lastBidder = _auction.bidder;

        // Refund the last bidder, if applicable
        if (lastBidder != address(0)) {
            _safeTransferETHWithFallback(lastBidder, _auction.amount);
        }

        auction.amount = msg.value;
        auction.bidder = payable(msg.sender);

        // Extend the auction if the bid was received within `timeBuffer` of the auction end time
        bool extended = _auction.endTime - block.timestamp < timeBuffer;
        if (extended) {
            auction.endTime = _auction.endTime = block.timestamp + timeBuffer;
        }

        emit AuctionBid(_auction.nounId, msg.sender, msg.value, extended);

        if (extended) {
            emit AuctionExtended(_auction.nounId, _auction.endTime);
        }
    }   

    function _startAuction(bytes64 skyNetID) internal {
        uint256 trixelId = trixels.mint(skyNetID);
        uint256 startTime = block.timestamp;
        uint256 endTime = startTime + duration;

        auction = Auction({
            trixelId: trixelId,
            amount: 0,
            startTime: startTime,
            endTime: endTime,
            bidder: payable(0),
            settled: false
        });

        emit AuctionCreated(trixelId, startTime, endTime);
    }

    function _settleAuction() internal {
        ITrixelsAuctionHouse.Auction memory _auction = auction;

        require(_auction.startTime != 0, "Auction hasn't begun");
        require(!_auction.settled, 'Auction has already been settled');
        require(block.timestamp >= _auction.endTime, "Auction hasn't completed");

        auction.settled = true;

        if (_auction.bidder == address(0)) {
            trixels.burn(_auction.nounId);
        } else {
            nouns.transferFrom(address(this), _auction.bidder, _auction.nounId);
        }

        if (_auction.amount > 0) {
            _safeTransferETHWithFallback(owner(), _auction.amount);
        }

        emit AuctionSettled(_auction.trixelId, _auction.bidder, _auction.amount);
    }

    function pause() external override onlyOwner {
        _pause();
    }

    function unpause() external override onlyOwner {
        _unpause();

        if (auction.startTime == 0 || auction.settled) {
            _createAuction();
        }
    }

    function setTimeBuffer(uint256 _timeBuffer) external override onlyOwner {
        timeBuffer = _timeBuffer;

        emit AuctionTimeBufferUpdated(_timeBuffer);
    }

    function setReservePrice(uint256 _reservePrice) external override onlyOwner {
        reservePrice = _reservePrice;

        emit AuctionReservePriceUpdated(_reservePrice);
    }
}