pragma solidity ^0.8.0



contract Trixels {
    address owner;
    bytes8[200][200] colors;
    address[200][200] lastEditors;

    constructor(bytes8[] colors) {
        owner = msg.sender;
        for (uint i=0; i<pixels.length; i++) {
            for (uint j=0; i<pixels[0].length; j++) {
                int index = (i * pixels.length) + j;
                bytes8 memory color = "#000";
                if (index < colors.length) {
                    color = colors[index];
                }
                colors[i][j] = color;
                lastEditors[i][j] = address(0);
            }
        }
    }

    function changePixel(int x, int y, bytes8 newColor) external {
        // Validate X, Y, and newColor
        require(x < 200, "x coord invalid");
        require(y < 200, "y coord invalid");
        colors[x][y] = newColor;
        lastEditors[x][y] = msg.sender; 
    }

    // use nouns contract for auction logic
    function getPixels() external returns (bytes8[][]) {
        return colors;
    }

    // Add auction functionality

    modifier onlyUpdater {
        require(msg.sender == owner);
        _;
    }    

    struct Pixel {
        int x;
        int y;
        bytes8 color;
        address lastEditor;
    }

    function massUpdatePixels(Pixel[] newPixels) external onlyUpdater {
        for (int i=0; i<newPixels.length; i++) {
            Pixel pixel = newPixels[i];
            colors[pixel.x][pixel.y] = pixel.color;
            lastEditors[pixel.x][pixel.y] = pixel.lastEditor;
        }
    }
}