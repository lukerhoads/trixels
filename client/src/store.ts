import { makeAutoObservable } from "mobx"
import { SimplePixel, Pixel } from "types/pixel";
import config from './config'

class Store {
    pixels: Pixel[] = [];
    activePixel: Pixel | undefined = undefined
    hoverPixel: SimplePixel | undefined = undefined
    userAddress: string = ""
    scale: number = 1

    constructor() {
        makeAutoObservable(this)
    }

    async fetchPixels() {
        fetch(config.apiUrl + "/pixels")
            .then(resp => resp.json())
            .then(data => this.pixels = data.data)
            .catch(err => console.error(err))
    }

    async fetchPixel(x: number, y: number): Promise<Pixel> {
        return fetch(config.apiUrl + `/pixel/${x}-${y}`)
            .then(resp => resp.json())
            .then(data => data)
            .catch(err => console.error(err))
    }

    subscribeToPixelChanges() {

    }

    setActivePixel(pixel: Pixel) {
        this.activePixel = pixel
    }

    setUserAddress(address: string) {
        this.userAddress = address
    }

    setScale(newScale: number) {
        this.scale = newScale
    }
}

const store = new Store()
export default store