import ApiClient from 'api-client';
import { makeAutoObservable } from 'mobx';
import { Log } from 'types/log';
import { OverlayInfo } from 'types/overlay';
import { Pixel, SimplePixel } from 'types/pixel';
import config from './config';

class Store {
    // Client
    client: ApiClient;

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

    // Overlay state
    overlayInfo: OverlayInfo | undefined = undefined;

    // Other state
    logs: Log[] = [];

    constructor(apiUrl: string) {
        makeAutoObservable(this);
        this.client = new ApiClient(apiUrl);
    }

    async fetchPixels() {
        let pixels = await this.client.getPixels();
        this.pixels = pixels;
    }

    async fetchPixel(x: number, y: number): Promise<Pixel> {
        return this.client.getPixelByCoord({ x: x, y: y });
    }

    setOverlayInfo(newInfo: OverlayInfo | undefined) {
        if (!newInfo) {
            this.overlayInfo = undefined;
            return;
        }
        this.overlayInfo = {
            ...newInfo,
            ...this.overlayInfo,
        };
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

        return this.client.updatePixel(newActivePixel);
    }
}

const store = new Store(config.alt.apiUrl);
export default store;
