import { useEthers } from '@usedapp/core';
import React, { useEffect, useRef, useState } from 'react';
import 'styles/haptics.scss';
import { Spinner } from './Spinner';
import { observer } from 'mobx-react';
import store from '../store';
import Status from './Status';
import config from 'config'

export type HapticsProps = {
  children: React.ReactNode;
};

// TODO:
// - Get metamask working
// - Have inputs and text fade in
// - Add tooltips for addresses
// - Add a direct color palette to edit
// - Add countdown to editable
// - Make logs fade in/out

const Haptics = ({ children }: HapticsProps) => {
  const { activateBrowserWallet, account } = useEthers();
  const [isAuthenticating, setIsAuthenticating] = useState(false);
  const [isChangingColor, setIsChangingColor] = useState(false);
  const [contextActive, setContextActive] = useState(false);
  const [canEditColor, setCanEditColor] = useState(false);
  const [isPreview, setIsPreview] = useState(false);
  const [timeLeft, setTimeLeft] = useState<string | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (!store.activePixel) return;
    let diff = computeDifference()
    if (!diff) return 
    if (diff[0] > 0 || diff[1] > 0) {
      setCanEditColor(false)
    } else {
      setCanEditColor(true)
    }
  }, []);

  const computeDifference = (): number[] | undefined => {
    if (!store.activePixel || store.activePixel.updatedAt == 'Never') return undefined;
    let nextEditableTime = new Date(store.activePixel.updatedAt)
    nextEditableTime.setSeconds(nextEditableTime.getSeconds() + config.editTimeoutSeconds)
    let now = new Date()
    let numMins = nextEditableTime.getMinutes() - now.getMinutes()
    let numSecs = nextEditableTime.getSeconds() - now.getSeconds()
    return [numMins, numSecs]
  }

  useEffect(() => {
    let diff = computeDifference()
    if (!diff) return
    setTimeLeft(`${diff[0]}m${diff[1]}s`)
    if (diff[0] > 0 || diff[1] > 0) {
      setCanEditColor(false)
    } else {
      setCanEditColor(true)
    }
  }, [store.activePixel])

  useEffect(() => {
    const timer = setTimeout(() => store.popFromLogs(), 10000)
    return () => clearTimeout(timer)
  }, [store.logs])

  const changeZoom = (factor: number) => {
    store.setScale(store.scale + factor);
  };

  const web3SignIn = () => {
    if (isAuthenticating || account) return;
    setIsAuthenticating(true);
    activateBrowserWallet();
  };

  const resetPreviewColor = () => {
    store.pushToLogs({
      mood: 'success',
      message: 'Reset preview.',
    });
    store.setActiveColor(store.activePixelOriginalColor);
  };

  const changePreviewColor = (e: React.ChangeEvent<HTMLInputElement>) => {
    let newColor = e.target.value;
    if (!newColor) return;
    if (!newColor.match('^#')) {
      store.pushToLogs({
        mood: 'warning',
        message: 'Pixel color does not start with a pound sign.',
      });
    }

    setIsPreview(true);
    if (newColor == store.activePixelOriginalColor) {
      setIsPreview(false);
    }

    store.setActiveColor(newColor);
  };

  const changeColor = async () => {
    if (!canEditColor) return;
    setIsChangingColor(true);
    store.setActivePixelColor().then(() => {
      setIsChangingColor(false);
    });
  };

  return (
    <div className='haptics'>
      { store.logs.length ? (
        <div className='logs'>
          {store.logs.slice(0, 2).map((log, idx) => (
            <div className={"log" + (idx == 0 ? ' top' : '')} key={idx}>
              <Status mood={log.mood} height={20} width={20} />
              <p>{log.message}</p>
            </div>
          ))}
        </div>
      ) : null}
      <div className='extra-info'>
        <div className='coords'>
          {store.activePixel ? (
            <p>
              Active X: {store.activePixel.x} - Active Y: {store.activePixel.y}
            </p>
          ) : null}
          {store.hoverPixel ? (
            <p>
              Hover X: {store.hoverPixel.x} - Hover Y: {store.hoverPixel.y}
            </p>
          ) : null}
        </div>
        <div className='web3-status'>
          {account ? (
            <>
              <Status width={20} height={20} mood='success' />
              <p>Web3 connected</p>
            </>
          ) : (
            <>
              <Status width={20} height={20} mood='error' />
              <p>Web3 disconnected</p>
            </>
          )}
        </div>
      </div>
      <div className='zoom'>
        <div className='plus' onClick={() => changeZoom(0.2)}>
          <svg xmlns='http://www.w3.org/2000/svg' className='ionicon' viewBox='0 0 512 512'>
            <title>Add</title>
            <path fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M256 112v288M400 256H112' />
          </svg>
        </div>
        <div className='minus' onClick={() => changeZoom(-0.2)}>
          <svg xmlns='http://www.w3.org/2000/svg' className='ionicon' viewBox='0 0 512 512'>
            <title>Remove</title>
            <path fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M400 256H112' />
          </svg>
        </div>
      </div>
      <div
        className={'context' + (contextActive ? ' context-expanded' : '') + (store.activePixel ? '' : ' context-inactive')}
        onClick={!contextActive ? () => setContextActive(true) : () => null}
      >
        <div className='mini-header'>
          <div className='color'>
            <div className='color-circle-border-wrap'>
              <div
                className='color-circle'
                style={{
                  backgroundColor: store.activePixel?.color ? store.activePixel.color : '#FFFFFF',
                }}
              />
            </div>
          </div>
          <div className='color-description'>
            <div className='text'>
              <p>{store.activePixel ? store.activePixel.color : 'No pixel selected'}</p>
            </div>
          </div>
          <div onClick={() => setContextActive(false)} className={'close' + (contextActive ? ' fadeIn' : ' fadeOut')}>
            <svg xmlns='http://www.w3.org/2000/svg' className='ionicon' viewBox='0 0 512 512'>
              <title>Close</title>
              <path fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M368 368L144 144M368 144L144 368' />
            </svg>
          </div>
        </div>
        {contextActive && store.activePixel ? (
          <div className='information'>
            <div className='more-info'>
              <p className='caption'>Last editor: </p>
              <p className='value'>{store.activePixel.editor}</p>
              <p className='caption'>Last edited time: </p>
              <p className='value'>{store.activePixel.updatedAt}</p>
              <p className='caption'>Countdown: </p>
              <p className='value'>{timeLeft ? timeLeft : 'None'}</p>
            </div>
            <div className='editor'>
              <div className='input-container'>
                <input placeholder='New Color' ref={inputRef} onChange={changePreviewColor} value={store.activePixel.color} />
                {isPreview ? <button onClick={() => resetPreviewColor()}>Reset</button> : null}
              </div>
              {account ? (
                <button onClick={() => changeColor()} disabled={canEditColor}>
                  {isChangingColor ? <Spinner width={20} height={20} /> : null}
                  Change Color
                </button>
              ) : (
                <button onClick={() => web3SignIn()}>
                  <div className='metamask-button'>
                    {isAuthenticating ? (
                      <>
                        <Spinner
                          width={20}
                          height={20}
                          extraStyle={{
                            marginRight: '10px',
                          }}
                        />
                        Signing in...
                      </>
                    ) : (
                      <>
                        <svg xmlns='http://www.w3.org/2000/svg' width='212' height='189' viewBox='0 0 212 189'>
                          <g fill='none' fill-rule='evenodd'>
                            <polygon
                              fill='#CDBDB2'
                              points='60.75 173.25 88.313 180.563 88.313 171 90.563 168.75 106.313 168.75 106.313 180 106.313 187.875 89.438 187.875 68.625 178.875'
                            />
                            <polygon
                              fill='#CDBDB2'
                              points='105.75 173.25 132.75 180.563 132.75 171 135 168.75 150.75 168.75 150.75 180 150.75 187.875 133.875 187.875 113.063 178.875'
                              transform='matrix(-1 0 0 1 256.5 0)'
                            />
                            <polygon
                              fill='#393939'
                              points='90.563 152.438 88.313 171 91.125 168.75 120.375 168.75 123.75 171 121.5 152.438 117 149.625 94.5 150.188'
                            />
                            <polygon fill='#F89C35' points='75.375 27 88.875 58.5 95.063 150.188 117 150.188 123.75 58.5 136.125 27' />
                            <polygon fill='#F89D35' points='16.313 96.188 .563 141.75 39.938 139.5 65.25 139.5 65.25 119.813 64.125 79.313 58.5 83.813' />
                            <polygon fill='#D87C30' points='46.125 101.25 92.25 102.375 87.188 126 65.25 120.375' />
                            <polygon fill='#EA8D3A' points='46.125 101.813 65.25 119.813 65.25 137.813' />
                            <polygon fill='#F89D35' points='65.25 120.375 87.75 126 95.063 150.188 90 153 65.25 138.375' />
                            <polygon fill='#EB8F35' points='65.25 138.375 60.75 173.25 90.563 152.438' />
                            <polygon fill='#EA8E3A' points='92.25 102.375 95.063 150.188 86.625 125.719' />
                            <polygon fill='#D87C30' points='39.375 138.938 65.25 138.375 60.75 173.25' />
                            <polygon fill='#EB8F35' points='12.938 188.438 60.75 173.25 39.375 138.938 .563 141.75' />
                            <polygon fill='#E8821E' points='88.875 58.5 64.688 78.75 46.125 101.25 92.25 102.938' />
                            <polygon fill='#DFCEC3' points='60.75 173.25 90.563 152.438 88.313 170.438 88.313 180.563 68.063 176.625' />
                            <polygon
                              fill='#DFCEC3'
                              points='121.5 173.25 150.75 152.438 148.5 170.438 148.5 180.563 128.25 176.625'
                              transform='matrix(-1 0 0 1 272.25 0)'
                            />
                            <polygon fill='#393939' points='70.313 112.5 64.125 125.438 86.063 119.813' transform='matrix(-1 0 0 1 150.188 0)' />
                            <polygon fill='#E88F35' points='12.375 .563 88.875 58.5 75.938 27' />
                            <path
                              fill='#8E5A30'
                              d='M12.3750002,0.562500008 L2.25000003,31.5000005 L7.87500012,65.250001 L3.93750006,67.500001 L9.56250014,72.5625 L5.06250008,76.5000011 L11.25,82.1250012 L7.31250011,85.5000013 L16.3125002,96.7500014 L58.5000009,83.8125012 C79.1250012,67.3125004 89.2500013,58.8750003 88.8750013,58.5000009 C88.5000013,58.1250009 63.0000009,38.8125006 12.3750002,0.562500008 Z'
                            />
                            <g transform='matrix(-1 0 0 1 211.5 0)'>
                              <polygon fill='#F89D35' points='16.313 96.188 .563 141.75 39.938 139.5 65.25 139.5 65.25 119.813 64.125 79.313 58.5 83.813' />
                              <polygon fill='#D87C30' points='46.125 101.25 92.25 102.375 87.188 126 65.25 120.375' />
                              <polygon fill='#EA8D3A' points='46.125 101.813 65.25 119.813 65.25 137.813' />
                              <polygon fill='#F89D35' points='65.25 120.375 87.75 126 95.063 150.188 90 153 65.25 138.375' />
                              <polygon fill='#EB8F35' points='65.25 138.375 60.75 173.25 90 153' />
                              <polygon fill='#EA8E3A' points='92.25 102.375 95.063 150.188 86.625 125.719' />
                              <polygon fill='#D87C30' points='39.375 138.938 65.25 138.375 60.75 173.25' />
                              <polygon fill='#EB8F35' points='12.938 188.438 60.75 173.25 39.375 138.938 .563 141.75' />
                              <polygon fill='#E8821E' points='88.875 58.5 64.688 78.75 46.125 101.25 92.25 102.938' />
                              <polygon fill='#393939' points='70.313 112.5 64.125 125.438 86.063 119.813' transform='matrix(-1 0 0 1 150.188 0)' />
                              <polygon fill='#E88F35' points='12.375 .563 88.875 58.5 75.938 27' />
                              <path
                                fill='#8E5A30'
                                d='M12.3750002,0.562500008 L2.25000003,31.5000005 L7.87500012,65.250001 L3.93750006,67.500001 L9.56250014,72.5625 L5.06250008,76.5000011 L11.25,82.1250012 L7.31250011,85.5000013 L16.3125002,96.7500014 L58.5000009,83.8125012 C79.1250012,67.3125004 89.2500013,58.8750003 88.8750013,58.5000009 C88.5000013,58.1250009 63.0000009,38.8125006 12.3750002,0.562500008 Z'
                              />
                            </g>
                          </g>
                        </svg>
                        Sign in with Metamask
                      </>
                    )}
                  </div>
                </button>
              )}
            </div>
          </div>
        ) : null}
      </div>
      {children}
    </div>
  );
};

export default observer(Haptics);
