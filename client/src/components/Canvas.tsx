import config from '../config';
import React, { useEffect, useRef, useState } from 'react';
import { SimplePixel } from 'types/pixel';
import '../styles/canvas.scss';
import Haptics from './Haptics';
import store from '../store';
import { observer } from 'mobx-react';
import { validateHexCode } from '../util';

const PIXEL_LENGTH = 10;
const CANVAS_LENGTH = PIXEL_LENGTH * config.params.imageDimensions;
const SCROLL_SENSITIVITY = 0.0005;
const MAX_ZOOM = 5;
const MIN_ZOOM = 0.1;

const Canvas = () => {
  let canvasRef = useRef<HTMLCanvasElement | null>(null);
  let [ctx, setCtx] = useState<CanvasRenderingContext2D | null>();

  useEffect(() => {
    store.fetchPixels();
    setCtx(canvasRef.current?.getContext('2d'));
    canvasRef.current?.addEventListener('mousedown', onPointerDown);
    canvasRef.current?.addEventListener('mouseup', onPointerUp);
    canvasRef.current?.addEventListener('mousemove', onPointerMove);
    canvasRef.current?.addEventListener('wheel', (e) => adjustScale(e.deltaY * SCROLL_SENSITIVITY));
    canvasRef.current?.addEventListener('click', onClick);
    draw();
  }, [canvasRef.current]);

  useEffect(() => {
    if (store.activePixel && validateHexCode(store.activePixel.color)) {
      redrawActivePixel();
    }
  }, [store.activePixel?.color]);

  // Debug this idk why not working
  const redrawActivePixel = () => {
    if (!canvasRef || !canvasRef.current || !ctx || !store.activePixel) return;

    ctx.translate(window.innerWidth / 2, window.innerHeight / 2);
    ctx.scale(store.scale, store.scale);
    ctx.translate(-window.innerWidth / 2 + store.xOffset, -window.innerHeight / 2 + store.yOffset);

    let middle: SimplePixel = {
      x: store.xOffset,
      y: store.yOffset,
    };

    let scaledCanvasLength = CANVAS_LENGTH * store.scale;
    let scaledPixelLength = PIXEL_LENGTH * store.scale;
    let topLeft: SimplePixel = {
      x: Math.floor(middle.x - scaledCanvasLength / 2),
      y: Math.floor(middle.y - scaledCanvasLength / 2),
    };

    const x = topLeft.x + store.activePixel.x * scaledPixelLength;
    const y = topLeft.y + store.activePixel.y * scaledPixelLength;
    ctx.fillStyle = store.activePixel.color;
    ctx.fillRect(x, y, scaledPixelLength, scaledPixelLength);
  };

  const draw = () => {
    if (!canvasRef || !canvasRef.current || !ctx) return;

    canvasRef.current.width = window.innerWidth;
    canvasRef.current.height = window.innerHeight;

    ctx.translate(window.innerWidth / 2, window.innerHeight / 2);
    ctx.scale(store.scale, store.scale);
    ctx.translate(-window.innerWidth / 2 + store.xOffset, -window.innerHeight / 2 + store.yOffset);

    ctx.fillStyle = '#FFF';
    ctx.fillRect(-CANVAS_LENGTH / 2, -CANVAS_LENGTH / 2, CANVAS_LENGTH, CANVAS_LENGTH);

    const zeroX = -CANVAS_LENGTH / 2;
    const zeroY = -CANVAS_LENGTH / 2;
    if (store.pixels) {
      store.pixels.forEach((pixel) => {
        if (!ctx) return;
        const x = zeroX + pixel.x * PIXEL_LENGTH;
        const y = zeroY + pixel.y * PIXEL_LENGTH;
        ctx.fillStyle = pixel.color;
        ctx.fillRect(x, y, PIXEL_LENGTH, PIXEL_LENGTH);
      });
    }

    requestAnimationFrame(draw);
  };

  const getEventLocation = (e: MouseEvent): SimplePixel | undefined => {
    if (e.clientX && e.clientY) {
      return { x: e.clientX, y: e.clientY };
    }
  };

  const onPointerDown = (e: MouseEvent) => {
    store.setIsDragging(true);
    let eventLocation = getEventLocation(e);
    if (!eventLocation || !eventLocation.x || !eventLocation.y) return;
    store.setDragStartX(eventLocation.x / store.scale - store.xOffset);
    store.setDragStartY(eventLocation.y / store.scale - store.yOffset);
  };

  const onPointerUp = (e: MouseEvent) => {
    store.setIsDragging(false);
  };

  const onPointerMove = (e: MouseEvent) => {
    if (store.isDragging) {
      let eventLocation = getEventLocation(e);
      if (!eventLocation || !eventLocation.x || !eventLocation.y) return;
      store.setXOffset(eventLocation.x / store.scale - store.dragStartX);
      store.setYOffset(eventLocation.y / store.scale - store.dragStartY);
    }
  };

  const adjustScale = (deltaScale: number) => {
    if (!store.isDragging) {
      let newScale = store.scale + deltaScale;
      newScale = Math.min(newScale, MAX_ZOOM);
      newScale = Math.max(newScale, MIN_ZOOM);
      store.setScale(newScale);
    }
  };

  const getPixelFromCoordinate = (coord: SimplePixel): SimplePixel | undefined => {
    let middle: SimplePixel = {
      x: store.xOffset,
      y: store.yOffset,
    };

    let scaledCanvasLength = CANVAS_LENGTH * store.scale;
    let topLeft: SimplePixel = {
      x: middle.x - scaledCanvasLength / 2,
      y: middle.y - scaledCanvasLength / 2,
    };

    let bottomRight: SimplePixel = {
      x: middle.x + scaledCanvasLength / 2,
      y: middle.y + scaledCanvasLength / 2,
    };

    if (coord.x < topLeft.x || coord.y < topLeft.y || coord.x > bottomRight.x || coord.y > bottomRight.y) {
      return undefined;
    }

    let scaledPixelLength = PIXEL_LENGTH * store.scale;
    return {
      x: Math.floor((coord.x - topLeft.x) / scaledPixelLength),
      y: Math.floor((coord.y - topLeft.y) / scaledPixelLength),
    };
  };

  const onClick = async (e: MouseEvent) => {
    if (!store.isDragging) {
      let eventLocation = getEventLocation(e);
      if (!eventLocation) return;
      let pixel = getPixelFromCoordinate(eventLocation);
      if (!pixel) {
        store.setActivePixel(undefined);
        return;
      }
      let fullPixel = await store.fetchPixel(pixel.x, pixel.y);
      if (fullPixel.color != '') {
        store.setActivePixel(fullPixel);
      } else {
        store.setActivePixel({
          x: pixel.x,
          y: pixel.y,
          color: '#FFFFFF',
        });
      }
    }
  };

  return (
    <Haptics>
      <canvas className='canvas' ref={canvasRef}></canvas>
    </Haptics>
  );
};

export default observer(Canvas);
