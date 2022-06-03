import config from '../config'
import React, { useEffect, useRef, useState } from 'react'
import { Pixel } from 'types/pixel'
import "../styles/canvas.scss"
import { Haptics } from './Haptics'

// TODO:
// - On load, need a gradient zoom into the canvas

let scale = 1
let pixelLength = 10
let canvasLength = pixelLength * config.imageDimensions

export const Canvas = () => {
    let canvasRef = useRef<HTMLCanvasElement | null>(null)
    let [ctx, setCtx] = useState<CanvasRenderingContext2D | null>()
    let [pixels, setPixels] = useState<Pixel[]>([{
        x: 0,
        y: 0,
        color: "#000",
        editor: "Hello",
        editedTime: "Hello"
    }])

    useEffect(() => {
        setCtx(canvasRef.current?.getContext("2d"))
        draw()
    }, [canvasRef.current])

    const draw = () => {
        if (!canvasRef || !canvasRef.current || !ctx) {
            return
        }
    
        canvasRef.current.width = window.innerWidth
        canvasRef.current.height = window.innerHeight
    
        ctx.translate(window.innerWidth / 2, window.innerHeight / 2)
        ctx.scale(scale, scale)
    
        ctx.fillStyle = "#FFF"
        ctx.fillRect(-canvasLength / 2, -canvasLength / 2, canvasLength, canvasLength)

        const zeroX = -canvasLength / 2
        const zeroY = -canvasLength / 2
        pixels.forEach(pixel => {
            const x = zeroX + pixel.x * pixelLength
            const y = zeroY + pixel.y * pixelLength
            ctx!.fillStyle = pixel.color
            ctx?.fillRect(x, y, pixelLength, pixelLength)
        })  
    }

    return (
        <Haptics>
            <canvas className="canvas" ref={canvasRef}></canvas>
        </Haptics>
    )
}