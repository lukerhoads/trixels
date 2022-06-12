import { useEthers } from '@usedapp/core';
import { CHAIN_ID } from 'config';
import { useTransition } from 'react';

export type Web3AuthInfo = {
    authenticating: boolean;
    authenticate: () => void;
    connectChain: () => void;
};

type WindowInstanceWithEthereum = Window & typeof globalThis & { ethereum?: any };

export const useWeb3Auth = (): Web3AuthInfo => {
    const { chainId, activateBrowserWallet } = useEthers();
    const [authenticating, startAuthenticating] = useTransition();

    const connectChain = async () => {
        const encodedChainId = '0x' + CHAIN_ID.toString(16);
        const ethereum = (window as WindowInstanceWithEthereum).ethereum;
        try {
            await ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: encodedChainId }],
            });
        } catch (err: any) {
            console.error(err.code);
        }
    };

    const authenticate = async () => {
        activateBrowserWallet();
        if (chainId != CHAIN_ID) {
            connectChain();
        }
    };

    return {
        authenticating,
        connectChain,
        authenticate,
    };
};
