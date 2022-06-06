import { makeAutoObservable } from 'mobx';
import { SimplePixel, Pixel } from 'types/pixel';
import { Log } from 'types/log';
import config from './config';

class Store {
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

  // Other state
  logs: Log[] = [];

  constructor() {
    makeAutoObservable(this);
  }

  async fetchPixels() {
    fetch(config.apiUrl + '/pixels')
      .then((resp) => resp.json())
      .then((data) => {
        this.pixels = data
      })
      .catch((err) => console.error(err));
  }

  async fetchPixel(x: number, y: number): Promise<Pixel> {
    const apiUrl = config.apiUrl + `/pixel/${x}-${y}`
    return fetch(apiUrl)
      .then((resp) => resp.json())
      .then((data) => data)
      .catch((err) => console.error(err));
  }

  async subscribeToPixelChanges() {}

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
    this.logs.shift()
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
    }

    return fetch(config.apiUrl + `/pixels`, {
      method: 'POST',
      body: JSON.stringify(newActivePixel),
    })
      .then((resp) => resp.json())
      .then((data) => data)
      .catch((err) => console.error(err));
  }
}

const store = new Store();
export default store;
