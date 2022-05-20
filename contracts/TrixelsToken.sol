// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import { Ownable } from '@openzeppelin/contracts/access/Ownable.sol';
import "./ERC721.sol";

interface ITrixelsToken {
    event TrixelCreated(uint256 indexed tokenId);

    event TrixelBurned(uint256 indexed tokenId);

    event MinterUpdated(address minter);

    event MinterLocked();

    function mint(bytes23 skyNetID) external returns (uint256);

    function burn(uint256 tokenId) external;
}

contract TrixelsToken is ITrixelsToken, ERC721, Ownable {
    address public minter;
    uint256 private _currentTrixelId;

    bool public minterLocked;

    constructor(address _minter) ERC721('Trixels', 'TRIX') {
        minter = _minter;
    }

    modifier onlyMinter() {
        require(msg.sender == minter, 'Sender is not the minter');
        _;
    }

    modifier whenMinterNotLocked() {
        require(!minterLocked, 'Minter is locked');
        _;
    }

    function mint(bytes23 skyNetID) public override onlyMinter returns (uint256) {
        _mint(minter, _currentTrixelId++, skyNetID);
        emit TrixelCreated(_currentTrixelId);
        return _currentTrixelId;
    }

    function burn(uint256 trixelId) public override onlyMinter {
        _burn(trixelId);
        emit TrixelBurned(trixelId);
    }

    function setMinter(address _minter) external onlyOwner whenMinterNotLocked {
        minter = _minter;
        emit MinterUpdated(_minter);
    }

    function lockMinter() external onlyOwner whenMinterNotLocked {
        minterLocked = true;
        emit MinterLocked();
    }
}