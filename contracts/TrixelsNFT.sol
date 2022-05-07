// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

contract TrixelsNFT is ERC721 {
    address public minter;
    uint256 private _currentTrixelId;

    constructor(address _minter) ERC721('Trixels', 'TRIX') {
        minter = _minter;
    }

    modifier onlyMinter() {
        require(msg.sender == minter, 'Sender is not the minter');
        _;
    }

    function mint(bytes64 skyNetID) public onlyMinter returns (uint256) {
        _mint(minter, _currentTrixelId++, skyNetID);
        return _currentTrixelId;
    }
}