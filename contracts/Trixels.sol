// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "./TrixelsNFT.sol";
import "./ERC721.sol";

contract Trixels {
    // Constants
    uint constant DIMENSION = 200;

    // Other contracts
    TrixelsNFT public trixels; 

    address public owner;
    bytes8[DIMENSION][DIMENSION] public colors;
    address[DIMENSION][DIMENSION] public lastEditors;

    struct Pixel {
        uint16 x;
        uint16 y;
        bytes8 color;
        address lastEditor;
    }

    constructor(bytes8[] memory newColors) {
        owner = msg.sender;
        for (uint i=0; i<DIMENSION; i++) {
            for (uint j=0; i<DIMENSION; j++) {
                uint index = (i * DIMENSION) + j;
                bytes8 color = "#000";
                if (index < colors.length) {
                    color = newColors[index];
                }
                colors[i][j] = color;
                lastEditors[i][j] = address(0);
            }
        }
    }

    function changePixel(uint x, uint y, bytes8 newColor) external {
        // Validate X, Y, and newColor
        require(x < DIMENSION, "x coord invalid");
        require(y < DIMENSION, "y coord invalid");
        require(isValidHexCode(newColor), "hex code invalid");
        colors[x][y] = newColor;
        lastEditors[x][y] = msg.sender; 
    }

    function massChangePixels(Pixel[] calldata newPixels) external onlyUpdater {
        for (uint i=0; i<newPixels.length; i++) {
            Pixel memory pixel = newPixels[i];
            colors[pixel.x][pixel.y] = pixel.color;
            lastEditors[pixel.x][pixel.y] = pixel.lastEditor;
        }
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }    

    function isValidHexCode(bytes8 hexCode) private returns (bool) {
        return true;
    }
}