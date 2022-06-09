import addresses from '../addresses.json';

export const validateHexCode = (hex: string) => {
    return hex.match('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$');
};

type ContractAddresses = {
    auctionHouse: string;
    dao: string;
    distributor: string;
    token: string;
};

export const getContractAddressesForChain = (chainId: number) => {
    const _addresses: Record<string, ContractAddresses> = addresses;
    if (!_addresses[chainId]) {
        throw new Error(`Unknown chain id ${chainId}`);
    }

    return _addresses[chainId];
};
