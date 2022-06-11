interface BaseAuction {
    tokenID: number 
    imageUrl: string
}

export type LiveAuction = BaseAuction & {
    endingDate: string;
    highestBid: number;
    highestBidder: string;
    minNextBid: number;
};

export type PastAuctionPreview = BaseAuction & {
    mintDate: string
}

export type PastAuction = BaseAuction & {
    salePrice: number;
    mintDate: string;
    winner: string;
    currentOwner: string;
};