import config from '../config'
import React, { useEffect, useRef, useState } from 'react'
import { Pixel, SimplePixel } from 'types/pixel'
import "../styles/canvas.scss"
import Haptics from './Haptics'
import store from '../store'

// TODO:
// - On load, need a gradient zoom into the canvas
// - Listen to page resize, redraw

const PIXEL_LENGTH = 10
const CANVAS_LENGTH = PIXEL_LENGTH * config.imageDimensions
const SCROLL_SENSITIVITY = 0.0005
const MAX_ZOOM = 5
const MIN_ZOOM = 0.1

export const Canvas = () => {
    let canvasRef = useRef<HTMLCanvasElement | null>(null)
    let [ctx, setCtx] = useState<CanvasRenderingContext2D | null>()
    let [isDragging, setIsDragging] = useState(false)
    let [xOffset, setXOffset] = useState(window.innerWidth / 2)
    let [yOffset, setYOffset] = useState(window.innerHeight / 2)
    let [dragStartX, setDragStartX] = useState(0)
    let [dragStartY, setDragStartY] = useState(0)

    useEffect(() => {
        store.fetchPixels()
        setCtx(canvasRef.current?.getContext("2d"))
        draw()
        canvasRef.current?.addEventListener('mousedown', onPointerDown)
        canvasRef.current?.addEventListener('mouseup', onPointerUp)
        canvasRef.current?.addEventListener('mousemove', onPointerMove)
        canvasRef.current?.addEventListener('wheel', (e) => adjustScale(e.deltaY * SCROLL_SENSITIVITY))
    }, [canvasRef.current])

    const draw = () => {
        if (!canvasRef || !canvasRef.current || !ctx) return
    
        canvasRef.current.width = window.innerWidth
        canvasRef.current.height = window.innerHeight
    
        ctx.translate(window.innerWidth / 2, window.innerHeight / 2)
        ctx.scale(store.scale, store.scale)
        ctx.translate(-window.innerWidth / 2 + xOffset, -window.innerHeight / 2 + yOffset)
    
        ctx.fillStyle = "#FFF"
        ctx.fillRect(-CANVAS_LENGTH / 2, -CANVAS_LENGTH / 2, CANVAS_LENGTH, CANVAS_LENGTH)

        const zeroX = -CANVAS_LENGTH / 2
        const zeroY = -CANVAS_LENGTH / 2
        store.pixels.forEach(pixel => {
            if (!ctx) return
            const x = zeroX + pixel.x * PIXEL_LENGTH
            const y = zeroY + pixel.y * PIXEL_LENGTH
            ctx.fillStyle = pixel.color
            ctx.fillRect(x, y, PIXEL_LENGTH, PIXEL_LENGTH)
        })  

        requestAnimationFrame(draw)
    }

    const getEventLocation = (e: MouseEvent): SimplePixel | undefined => {
        if (e.clientX && e.clientY) {
            return { x: e.clientX, y: e.clientY }
        }
    }

    const onPointerDown = (e: MouseEvent) => {
        setIsDragging(true)
        let eventLocation = getEventLocation(e)
        if (!eventLocation || !eventLocation.x || !eventLocation.y) return 
        setDragStartX(eventLocation.x / store.scale - yOffset)
        setDragStartY(eventLocation.y / store.scale - yOffset)
    }

    const onPointerUp = (e: MouseEvent) => {
        setIsDragging(false)
    }
    

    const onPointerMove = (e: MouseEvent) => {
        if (isDragging) {
            let eventLocation = getEventLocation(e)
            if (!eventLocation || !eventLocation.x || !eventLocation.y) return 
            setXOffset(eventLocation.x / store.scale - dragStartX)
            setYOffset(eventLocation.y / store.scale - dragStartY)
        }
    }

    const adjustScale = (deltaScale: number) => {
        if (!isDragging) {
            let newScale = store.scale + deltaScale
            newScale = Math.min(newScale, MAX_ZOOM)
            newScale = Math.max(newScale, MIN_ZOOM)
            store.setScale(newScale)
        }
    }

    const getPixelFromCoordinate = (coord: SimplePixel): SimplePixel => {
        return {
            x: 0,
            y: 0
        }
    }

    const onClick = (e: MouseEvent) => {
        let eventLocation = getEventLocation(e)
        if (!eventLocation) return
        let pixel = getPixelFromCoordinate(eventLocation)
        let fullPixel = store.fetchPixel(pixel.x, pixel.y)
        store.setActivePixel(fullPixel)
    }

    return (
        <Haptics>
            <canvas className="canvas" ref={canvasRef}></canvas>
        </Haptics>
    )
}