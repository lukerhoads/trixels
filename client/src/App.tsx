import React from 'react';
import Auction from 'pages/Auction';
import Home from 'pages/Home';
import NotFound from 'pages/NotFound';
import { Routes, Route, BrowserRouter as Router } from 'react-router-dom';
import './styles/app.scss';

export const App = () => {
    return (
        <div className='grid-1'>
            <div className='grid-2'>
                <Router>
                    <Routes>
                        <Route path='/' element={<Home />} />
                        <Route path='/auction/:tokenID' element={<Auction />} />
                        <Route path='/*' element={<NotFound />} />
                    </Routes>
                </Router>
            </div>
        </div>
    );
};
