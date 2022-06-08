import { createRoot } from 'react-dom/client';
import { DAppProvider } from '@usedapp/core';

import { App } from './App';
import config, { CHAIN_ID, supportedChainUrls } from './config';

export const useDappConfig = {
  readOnlyChainId: CHAIN_ID,
  readOnlyUrls: {
    [CHAIN_ID]: supportedChainUrls[CHAIN_ID],
  },
};

const container = document.getElementById('root')!;
const root = createRoot(container);
root.render(
  <DAppProvider config={useDappConfig}>
    <App />
  </DAppProvider>
);
