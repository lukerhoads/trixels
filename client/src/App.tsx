import { DAppProvider } from '@usedapp/core';
import { useWeb3Auth } from 'hooks/useWeb3Auth';
import Auction from 'pages/Auction';
import Dao from 'pages/DAO';
import Home from 'pages/Home';
import NotFound from 'pages/NotFound';
import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { CHAIN_ID, supportedChainUrls } from './config';
import './styles/app.scss';

export const useDappConfig = {
    readOnlyChainId: CHAIN_ID,
    readOnlyUrls: {
        [CHAIN_ID]: supportedChainUrls[CHAIN_ID],
    },
};

export const App = () => {
    const { connectChain } = useWeb3Auth();

    useEffect(() => {
        connectChain();
    }, []);

    return (
        <DAppProvider config={useDappConfig}>
            <div className='grid-1'>
                <div className='grid-2'>
                    <Router>
                        <Routes>
                            <Route path='/' element={<Home />} />
                            <Route path='/dao' element={<Dao />} />
                            <Route path='/auction/:tokenID' element={<Auction />} />
                            <Route path='/dao/:proposalID' element={<Dao />} />
                            <Route path='/*' element={<NotFound />} />
                        </Routes>
                    </Router>
                </div>
            </div>
        </DAppProvider>
    );
};
