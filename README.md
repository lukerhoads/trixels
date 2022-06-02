# Trixels
Trixels is a semi-decentralized application that brings together the world through NFTs. Ever heard of r/place or NounsDAO? Trixels is a mix of both of these cool projects. Users contribute to a pixel canvas which is stored on chain (not right now). Every two weeks (configurable by DAO) the canvas is snapshotted and auctioned off as an NFT. The auction lasts one day and proceeds go to the DAO and to the pixel contributors. The DAO has control over how much goes to the DAO, and how to weight contributions.  

## Design
- Every one day, pixels are updated on-chain
- Every two weeks, a new NFT is minted and auctioned
- 500x500 board
- 5 minute timeout on pixel edits

## Inspiration
- https://nouns.wtf/
- https://www.reddit.com/r/place/

## Todo
- Figure out a way to store onchain
- Make a good GUI
- Write tests for server
- Dockerize server
- Edit makefile, add flags

## Design Notes
- Pixel values are stored centrally
- Images are simply put up for auction on CRON by central server
- Everything else is decentralized

## Data structures for postgres
- Key is hash of string of pixel coordinate 
    - ex: "500,500" --> 61d0c4cea3f6428da3fef7d53df6c434af12240dc8d051eec0063d16c1aba428
- values are hex color, last editor address, last edited time

## Frontend process
- Sign in with metamask, get address
- Edit pixel with address