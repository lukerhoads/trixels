import { useEthers } from '@usedapp/core';
import React, { useEffect, useRef, useState } from 'react';
import 'styles/haptics.scss';
import { Spinner } from './Spinner';
import { observer } from 'mobx-react';
import store from '../store';
import Status from './Status';
import config from 'config'
import { validateHexCode } from '../util'

export type HapticsProps = {
  children: React.ReactNode;
};

// TODO:
// - Get metamask working
// - Have inputs and text fade in
// - Add tooltips for addresses
// - Add a direct color palette to edit
// - Make logs fade in/out
// - Fix why logs aren't autoremoving
// - Find out why we are fetching data twice
// - convert times to local zone

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
    applyTimer()
  }, []);

  useEffect(() => {
    applyTimer()
  }, [store.activePixel])

  const applyTimer = () => {
    if (!store.activePixel) return;
    if (!store.activePixel.updatedAt) {
      setCanEditColor(true)
      return
    }

    let nextEditableTime = new Date(store.activePixel.updatedAt)
    nextEditableTime.setSeconds(nextEditableTime.getSeconds() + config.editTimeoutSeconds)
    let now = new Date()
    let numMins = nextEditableTime.getMinutes() - now.getMinutes()
    let numSecs = nextEditableTime.getSeconds() - now.getSeconds()
    if (numMins >= 0 && numSecs >= 0) {
      setTimeLeft(`${numMins}m${numSecs}s`)
      setCanEditColor(false)
    } else {
      setTimeLeft(`0m0s, can edit`)
      setCanEditColor(true)
    }
  }

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
      return
    }

    setIsPreview(true);
    if (newColor == store.activePixelOriginalColor) {
      setIsPreview(false);
    }

    store.setActiveColor(newColor);
  };

  const changeColor = async () => {
    if (!canEditColor || !account) return;
    setIsChangingColor(true);
    store.setActivePixelColor(account).then(() => {
      setIsChangingColor(false);
      store.fetchPixels()
    });
  };

  const parseDate = (dateStr: string) => {
    let date = new Date(dateStr)
    return `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`
  }

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
      <div className="socials">
        <div className="social">
          <a href={`https://instagram.com/${config.instagramUser}`} target="_blank">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
          </a>
        </div>
        <div className="social">
          <a href={`https://facebook.com/${config.facebookUser}`} target="_blank">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M22.675 0h-21.35c-.732 0-1.325.593-1.325 1.325v21.351c0 .731.593 1.324 1.325 1.324h11.495v-9.294h-3.128v-3.622h3.128v-2.671c0-3.1 1.893-4.788 4.659-4.788 1.325 0 2.463.099 2.795.143v3.24l-1.918.001c-1.504 0-1.795.715-1.795 1.763v2.313h3.587l-.467 3.622h-3.12v9.293h6.116c.73 0 1.323-.593 1.323-1.325v-21.35c0-.732-.593-1.325-1.325-1.325z"/></svg>
          </a>
        </div>
      </div>
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
              <p>{store.activePixel ? validateHexCode(store.activePixel.color) ? store.activePixel.color : 'Invalid hex' : 'No pixel selected'}</p>
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
              <p className='value'>{store.activePixel.editor ? store.activePixel.editor : 'None'}</p>
              <p className='caption'>Last edited time: </p>
              <p className='value'>{store.activePixel.updatedAt ? parseDate(store.activePixel.updatedAt) : 'None'}</p>
              <p className='caption'>Countdown: </p>
              <p className='value'>{timeLeft ? timeLeft : 'None'}</p>
            </div>
            <div className='editor'>
              <div className='input-container'>
                <input placeholder='New Color' ref={inputRef} onChange={changePreviewColor} value={store.activePixel.color} />
                {isPreview ? <button onClick={() => resetPreviewColor()}>Reset</button> : null}
              </div>
              {account ? (
                <button onClick={() => changeColor()} disabled={!canEditColor}>
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
