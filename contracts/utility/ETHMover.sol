// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IWETH.sol";

contract ETHMover {
    address weth;

    constructor(address _weth) {
        weth = _weth;
    }

    function _safeTransferETH(address to, uint256 value) internal returns (bool) {
        (bool success, ) = to.call{ value: value, gas: 30_000 }("");
        return success;
    }

    function _safeTransferETHWithData(address to, uint256 value, bytes calldata _transactionData) internal returns (bool) {
        (bool success, ) = to.call{ value: value, gas: 30_000 }(_transactionData);
        return success;
    }

    function _safeTransferETHWithFallback(address to, uint256 value) internal {
        if (!_safeTransferETH(to, value)) {
            IWETH(weth).deposit{ value: value }();
            IERC20(weth).transfer(to, value);
        }
    }
}