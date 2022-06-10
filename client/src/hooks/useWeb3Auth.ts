import { useTransition } from 'react';
import { useEthers } from '@usedapp/core';
import config, { CHAIN_ID } from 'config';

export type Web3AuthInfo = {
    authenticating: boolean;
    authenticate: () => void;
};

export const useWeb3Auth = (): Web3AuthInfo => {
    const { activateBrowserWallet, account } = useEthers();
    const [authenticating, startAuthenticating] = useTransition();

    const authenticate = () => {
        const encodedChainId = '0x' + CHAIN_ID.toString(16)
        startAuthenticating(() => {
            // Prompt network switch
            // Inspired by https://stackoverflow.com/questions/68252365/how-to-trigger-change-blockchain-network-request-on-metamask
            // if (window.ethereum && window.ethereum.networkVersion != CHAIN_ID) {
                window.ethereum
                    .request({
                        method: 'wallet_switchEthereumChain',
                        params: [{ chainId: encodedChainId }],
                    })
                    .catch((err: any) => {
                        if (err.code === 4902) {
                            window.ethereum
                                .request({
                                    method: 'wallet_addEthereumChain',
                                    params: [
                                        {
                                            chainId: encodedChainId,
                                            rpcUrl: config.app.jsonRpcUri,
                                        },
                                    ],
                                })
                                .catch((error: any) => console.error(error));
                        } else {
                            console.error(err);
                        }
                    });
            // } else {
            //     console.error('Metamask is not installed');
            // }
            activateBrowserWallet();
        });
    };

    return {
        authenticating,
        authenticate,
    };
};
