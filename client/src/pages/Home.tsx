import 'styles/home.scss';

import Canvas from 'components/Canvas';
import Overlay from 'components/Overlay';
import Header from 'layout/header';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
        <div className='home'>
            <Overlay />
            <Canvas />
            <Header>
                <Link to='/auction/current'>Auction</Link>
            </Header>
        </div>
    );
};

export default Home;
