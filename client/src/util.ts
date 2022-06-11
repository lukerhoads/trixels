import addresses from '../addresses.json';
import config from './config'

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

export const addDays = (date1: Date, numMinutes: number): Date => {
    return new Date(date1.getTime() + config.params.editTimeoutMinutes * 60000);
}

export const getDateDiff = (date1: Date, date2: Date) => {
    return new Date(date1.getTime() - date2.getTime());
}