import * as ReactDOM from 'react-dom';
import { DAppProvider } from '@usedapp/core';

import { App } from './App';
import config from './config';

ReactDOM.render(
    // <DAppProvider config={{
    //     readOnlyChainId: config.chainId,
    //     readOnlyUrls: {
    //         [config.chainId]: config.rpcUrl,
    //     },
    // }}>
      <App />,
    // </DAppProvider>, 
    document.getElementById('root')
);
