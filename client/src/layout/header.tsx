import React, { Children } from 'react';
import { Link } from 'react-router-dom';
import '../styles/header.scss';

export type HeaderProps = {
    children: React.ReactNode;
};

const Header = ({ children }: HeaderProps) => {
    return (
        <div className='header'>
            <div className='header-content'>
                <div className='header-logo'>
                    <h2>
                        <Link to='/'>Trixels</Link>
                    </h2>
                </div>
                <div className='header-spacer' />
                <div className='header-links'>
                    {Children.map(children, (child) => (
                        <div className='header-link'>{child}</div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Header;
