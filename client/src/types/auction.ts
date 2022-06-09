export type Auction = {
  tokenID: number;
  metadataUrl: string;
  imageUrl: string;
  createdAt: string;
  saleValue: number;
};

export type PastAuction = {
  tokenID: number;
  salePrice: number;
  saleDate: string;
  winner: string;
  currentOwner: string;
};

export type LiveAuction = {
  tokenID: number;
  endingDate: string;
  highestBid: number;
  highestBidder: string;
  minBidIncrement: number;
};
