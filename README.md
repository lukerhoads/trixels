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
- At the moment, the pixel values will have to be stored on a server
- Figure out how to efficiently store hex values
- Write code for DAO
- Deploy frontend for auction
- Add timeouts for pixel changes
- Make a good GUI

## Design Notes
- Pixel values are stored centrally
- Images are simply put up for auction on CRON by central server
- Everything else is decentralized

## Data structures for postgres
- Key is hash of string of pixel coordinate 
    - ex: "500,500" --> 61d0c4cea3f6428da3fef7d53df6c434af12240dc8d051eec0063d16c1aba428
- values are hex color, last editor address, last edited time

## Frontend process
- sign in with metamask, get address
- edit pixel with address