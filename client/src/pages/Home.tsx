import React, { useEffect, useRef, useState } from 'react'
import "styles/home.scss"

import { Header } from 'layout/header'
import { Canvas } from 'components/Canvas'

export const Home = () => {
    return (
        <div className="home">
            <Header />
            <Canvas />
        </div>
    )
}