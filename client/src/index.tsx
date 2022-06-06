import { createRoot } from 'react-dom/client';
import { DAppProvider } from '@usedapp/core';

import { App } from './App';
import config from './config';

const container = document.getElementById('root')!;
const root = createRoot(container);
root.render(
  <DAppProvider config={{
      readOnlyChainId: config.chainId,
      readOnlyUrls: {
          [config.chainId]: config.rpcUrl,
      },
  }}>
    <App />
  </DAppProvider>,
);
