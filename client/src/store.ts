import { makeAutoObservable } from 'mobx';
import { SimplePixel, Pixel } from 'types/pixel';
import config from './config';

class Store {
  // Haptics state
  pixels: Pixel[] = [];
  activePixel: Pixel | undefined = undefined;
  hoverPixel: SimplePixel | undefined = undefined;
  userAddress: string = '';

  // Canvas state
  scale: number = 1;
  isDragging: boolean = false;
  xOffset: number = window.innerWidth / 2;
  yOffset: number = window.innerHeight / 2;
  dragStartX: number = 0;
  dragStartY: number = 0;

  constructor() {
    makeAutoObservable(this);
  }

  async fetchPixels() {
    fetch(config.apiUrl + '/pixels')
      .then((resp) => resp.json())
      .then((data) => (this.pixels = data.data))
      .catch((err) => console.error(err));
  }

  async fetchPixel(x: number, y: number): Promise<Pixel> {
    return fetch(config.apiUrl + `/pixel/${x}-${y}`)
      .then((resp) => resp.json())
      .then((data) => {
        return data;
      })
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

  setIsDragging(newIsDragging: boolean) {
    this.isDragging = newIsDragging;
  }

  setActivePixel(pixel: Pixel) {
    this.activePixel = pixel;
  }

  setUserAddress(address: string) {
    this.userAddress = address;
  }

  setScale(newScale: number) {
    this.scale = newScale;
  }
}

const store = new Store();
export default store;
