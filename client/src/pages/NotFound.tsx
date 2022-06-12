import Header from 'layout/header';
import { Link } from 'react-router-dom';
import '../styles/notfound.scss';

const NotFound = () => {
    return (
        <div className='not-found'>
            <Header>
                <Link to='/'>Home</Link>
            </Header>
            <h1>
                Page not found! Return <Link to='/'>home</Link>.
            </h1>
        </div>
    );
};

export default NotFound;
