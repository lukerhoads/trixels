import React from 'react'
import "../styles/header.scss"

export const Header = () => {
    return (
        <div className="header">
            <div className="header-content">
                <div className="header-logo">
                    Trixels
                </div>
                <div className="header-spacer" />
                <div className="header-links">
                    <div className="header-link">
                        <a href="/auction">Auction</a>
                    </div>
                </div>
            </div>
        </div>
    )
}