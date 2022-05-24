// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface IDistributor {
    struct Contributor {
        address addr;
        uint numContribs;
    }

    function distribute(Contributor[] memory _contributors, uint256 _salePrice) external;

    function withdraw() external;
}