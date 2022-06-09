import React, { useEffect, useRef, useState } from 'react';
import 'styles/home.scss';

import Header from 'layout/header';
import Canvas from 'components/Canvas';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
        <div className='home'>
            <Canvas />
            <Header>
                <Link to='/auction/current'>Auction</Link>
            </Header>
        </div>
    );
};

export default Home;
