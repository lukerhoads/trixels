import { ethers } from 'ethers';
import { makeAutoObservable } from 'mobx';
import { SimplePixel, Pixel } from 'types/pixel';
import { Log } from 'types/log';
import config from './config';
import { Auction } from 'types/auction';
import ApiClient from 'api-client';

class Store {
    // Client
    client: ApiClient

    // Haptics state
    pixels: Pixel[] = [];
    activePixel: Pixel | undefined = undefined;
    activePixelOriginalColor: string = '';
    hoverPixel: SimplePixel | undefined = undefined;
    userAddress: string = '';

    // Canvas state
    scale: number = 1;
    isDragging: boolean = false;
    xOffset: number = window.innerWidth / 2;
    yOffset: number = window.innerHeight / 2;
    dragStartX: number = 0;
    dragStartY: number = 0;

    // Auction state
    activeAuction: Auction | undefined = {
        tokenID: 1,
        metadataUrl: 'https://penis.com',
        imageUrl: 'https://cdn.britannica.com/91/181391-050-1DA18304/cat-toes-paw-number-paws-tiger-tabby.jpg?q=60',
        createdAt: '',
        saleValue: 10,
    };
    pastAuctions: Auction[] = [
        {
            tokenID: 1,
            metadataUrl: 'https://penis.com',
            imageUrl: 'https://cdn.britannica.com/91/181391-050-1DA18304/cat-toes-paw-number-paws-tiger-tabby.jpg?q=60',
            createdAt: '',
            saleValue: 10,
        },
    ];

    // Other state
    logs: Log[] = [];

    constructor(apiUrl: string) {
        makeAutoObservable(this);
        this.client = new ApiClient(apiUrl)
    }

    async fetchPixels() {
        let pixels = await this.client.getPixels()
        this.pixels = pixels
    }

    async fetchPixel(x: number, y: number): Promise<Pixel> {
        return this.client.getPixelByCoord({ x: x, y: y })
    }

    setXOffset(xOff: number) {
        this.xOffset = xOff;
    }

    setYOffset(yOff: number) {
        this.yOffset = yOff;
    }

    setDragStartX(newDrag: number) {
        this.dragStartX = newDrag;
    }

    setDragStartY(newDrag: number) {
        this.dragStartY = newDrag;
    }

    pushToLogs(newLog: Log) {
        this.logs.push(newLog);
    }

    popFromLogs() {
        this.logs.shift();
    }

    setIsDragging(newIsDragging: boolean) {
        this.isDragging = newIsDragging;
    }

    setActiveColor(newColor: string) {
        if (!this.activePixel) return;
        this.activePixel.color = newColor;
    }

    setActivePixel(pixel: Pixel | undefined) {
        if (pixel) {
            this.activePixelOriginalColor = pixel.color;
        }
        this.activePixel = pixel;
    }

    setUserAddress(address: string) {
        this.userAddress = address;
    }

    setScale(newScale: number) {
        this.scale = newScale;
    }

    async setActivePixelColor(editor: string) {
        if (!this.activePixel) return;
        if (!this.activePixel.color.match('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$')) {
            this.pushToLogs({
                mood: 'error',
                message: 'New pixel color is not a valid hex code.',
            });
        }

        const newActivePixel: Pixel = {
            x: this.activePixel.x,
            y: this.activePixel.y,
            color: this.activePixel.color,
            editor: editor,
        };

        return this.client.updatePixel(newActivePixel)
    }
}

const store = new Store(config.alt.apiUrl);
export default store;
