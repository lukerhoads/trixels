// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "./TrixelsToken.sol";
import "./ERC721.sol";
import { Ownable } from '@openzeppelin/contracts/access/Ownable.sol';

interface ITrixels {
    struct Pixel {
        uint16 x;
        uint16 y;
        bytes3 color;
        address lastEditor;
    }

    event PixelChanged(uint indexed x, uint indexed y, bytes3 oldColor, bytes8 indexed newColor, address user);

    function changePixel(uint x, uint y, bytes3 newColor) external;

    function massChangePixels(Pixel[] calldata newPixels) external;
}

contract Trixels is ITrixels, Ownable {
    // Constants
    uint constant DIMENSION = 200;

    bytes3[DIMENSION][DIMENSION] public colors;
    address[DIMENSION][DIMENSION] public lastEditors;

    // function initialize() external {
    //     for (uint i=0; i<DIMENSION; i++) {
    //         for (uint j=0; i<DIMENSION; j++) {
    //             colors[i][j] = "#000";
    //             lastEditors[i][j] = address(0);
    //         }
    //     }
    // }

    function changePixel(uint x, uint y, bytes3 newColor) external {
        // Validate X, Y, and newColor
        require(x < DIMENSION, "x coord invalid");
        require(y < DIMENSION, "y coord invalid");
        bytes3 oldColor = colors[x][y];
        colors[x][y] = newColor;
        lastEditors[x][y] = msg.sender; 
        emit PixelChanged(x, y, oldColor, newColor, msg.sender);
    }

    function massChangePixels(ITrixels.Pixel[] calldata newPixels) external onlyOwner {
        for (uint i=0; i<newPixels.length; i++) {
            Pixel memory pixel = newPixels[i];
            colors[pixel.x][pixel.y] = pixel.color;
            lastEditors[pixel.x][pixel.y] = pixel.lastEditor;
        }
    }
}