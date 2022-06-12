import { Pixel } from 'types/pixel';

export type Coordinate = {
    x: number;
    y: number;
};

export type Trixel = {
    tokenID: number;
    metadataUrl: string;
    createdAt: string;
};

export type Metadata = {
    name: string;
    image: string;
    description: string;
};

class ApiClient {
    private url: string;

    constructor(url: string) {
        this.url = url;
    }

    async getLiveTrixel(): Promise<Trixel> {
        return fetch(this.url + '/trixel/live')
            .then((res) => res.json())
            .then((trixel) => trixel as Trixel);
    }

    async getTrixels(): Promise<Trixel[]> {
        return fetch(this.url + '/trixels')
            .then((res) => res.json())
            .then((trixels) => trixels as Trixel[]);
    }

    async getTrixelById(trixelId: number): Promise<Trixel> {
        return fetch(this.url + `/trixels/find/${trixelId}`)
            .then((res) => res.json())
            .then((trixel) => trixel as Trixel);
    }

    async getPixelByCoord(coord: Coordinate): Promise<Pixel> {
        return fetch(this.url + `/pixel/${coord.x}-${coord.y}`)
            .then((res) => res.json())
            .then((pixel) => pixel as Pixel);
    }

    async getPixels(): Promise<Pixel[]> {
        return fetch(this.url + `/pixels`)
            .then((res) => res.json())
            .then((pixels) => pixels as Pixel[]);
    }

    async updatePixel(newPixel: Pixel) {
        return fetch(this.url + `/pixels`, {
            method: 'POST',
            body: JSON.stringify(newPixel),
        }).then(res => {
            console.log(res)
        });
    }

    async getMetadata(metaUrl: string): Promise<Metadata> {
        return fetch(metaUrl)
            .then((res) => res.json())
            .then((meta) => meta as Metadata);
    }
}

export default ApiClient;
