import { ChainId } from '@usedapp/core';
import { getContractAddressesForChain } from './util';

type SupportedChains = ChainId.Hardhat;

export const CHAIN_ID: SupportedChains = ChainId.Hardhat;

export const supportedChainUrls = {
  [ChainId.Hardhat]: 'http://localhost:8545',
};

type AppConfig = {
  jsonRpcUri: string;
  wsRpcUri: string;
};

const app: Record<SupportedChains, AppConfig> = {
  [ChainId.Hardhat]: {
    jsonRpcUri: 'http://localhost:8545',
    wsRpcUri: 'ws://localhost:8545',
  },
};

type AltConfig = {
  apiUrl: string;
  facebookUser: string;
  instagramUser: string;
};

const altConfig: AltConfig = {
  apiUrl: 'http://localhost:8080',
  facebookUser: '',
  instagramUser: '',
};

type Params = {
  imageDimensions: number;
  editTimeoutSeconds: number;
};

const params: Params = {
  imageDimensions: 30,
  editTimeoutSeconds: 300,
};

const config = {
  app: app[CHAIN_ID],
  addresses: getContractAddressesForChain(CHAIN_ID),
  params: params,
  alt: altConfig,
};

export default config;
