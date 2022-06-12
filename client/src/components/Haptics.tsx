import { useEthers } from '@usedapp/core';
import config from 'config';
import { useWeb3Auth } from 'hooks/useWeb3Auth';
import { observer } from 'mobx-react';
import React, { useEffect, useRef, useState } from 'react';
import 'styles/haptics.scss';
import store from '../store';
import { addDays, getDateDiff, validateHexCode } from '../util';
import Address from './Address';
import { Spinner } from './Spinner';
import Status from './Status';

export type HapticsType = 'logs' | 'full'

export type HapticsProps = {
    type: HapticsType
    children: React.ReactNode;
};

const Haptics = ({ children, type }: HapticsProps) => {
    const { account, deactivate } = useEthers();
    const { authenticating, authenticate } = useWeb3Auth();
    const [isChangingColor, setIsChangingColor] = useState(false);
    const [contextActive, setContextActive] = useState(false);
    const [canEditColor, setCanEditColor] = useState(false);
    const [isPreview, setIsPreview] = useState(false);
    const [timeLeft, setTimeLeft] = useState<string | null>(null);
    const inputRef = useRef<HTMLInputElement>(null);

    useEffect(() => {
        const timer = setInterval(applyTimer, 1000);
        return () => clearInterval(timer);
    }, [store.activePixel]);

    const applyTimer = () => {
        if (!store.activePixel) return;
        if (!store.activePixel.updatedAt) {
            setCanEditColor(true);
            return;
        }

        let editedAt = new Date(store.activePixel.updatedAt);
        let nextEditableTime = addDays(editedAt, 1);
        let now = new Date();
        if (now < nextEditableTime) {
            let diff = getDateDiff(nextEditableTime, now);
            setTimeLeft(`${diff.getMinutes()}m${diff.getSeconds()}s`);
            setCanEditColor(false);
        } else {
            setTimeLeft(`Can edit`);
            setCanEditColor(true);
        }
    };

    useEffect(() => {
        const timer = setTimeout(() => {
            store.popFromLogs()
        }, 5000);
        return () => clearTimeout(timer);
    }, [store.logs]);

    const changeZoom = (factor: number) => {
        store.setScale(store.scale + factor);
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
            return;
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
        await store.setActivePixelColor(account);
        await store.fetchPixels();
        store.setActivePixel(await store.fetchPixel(store.activePixel!.x, store.activePixel!.y));
        setIsChangingColor(false);
    };

    const parseDate = (dateStr: string) => {
        let date = new Date(dateStr);
        return `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()} Central`;
    };

    const colorClick = () => {
        store.setOverlayInfo({
            type: 'color-picker',
        });
    };

    const authenticateClick = () => {
        store.setOverlayInfo({
            type: 'auth-modal',
        });
    };

    

    return (
        <div className='haptics'>
            {store.logs.length && (
                <div className='logs'>
                    {store.logs.slice(0, 2).map((log, idx) => (
                        <div className={'log' + (idx == 0 ? ' top' : '')} key={idx}>
                            <Status mood={log.mood} height={20} width={20} />
                            <p>{log.message}</p>
                        </div>
                    ))}
                </div>
            )}
            { type == "full" && (

                <>
                    <div className='socials'>
                        <div className='social'>
                            <a href={`https://instagram.com/${config.alt.instagramUser}`} target='_blank'>
                                <svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24'>
                                    <path d='M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z' />
                                </svg>
                            </a>
                        </div>
                        <div className='social'>
                            <a href={`https://facebook.com/${config.alt.facebookUser}`} target='_blank'>
                                <svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24'>
                                    <path d='M22.675 0h-21.35c-.732 0-1.325.593-1.325 1.325v21.351c0 .731.593 1.324 1.325 1.324h11.495v-9.294h-3.128v-3.622h3.128v-2.671c0-3.1 1.893-4.788 4.659-4.788 1.325 0 2.463.099 2.795.143v3.24l-1.918.001c-1.504 0-1.795.715-1.795 1.763v2.313h3.587l-.467 3.622h-3.12v9.293h6.116c.73 0 1.323-.593 1.323-1.325v-21.35c0-.732-.593-1.325-1.325-1.325z' />
                                </svg>
                            </a>
                        </div>
                    </div>
                    <div className='extra-info'>
                        <div className='coords'>
                            {store.activePixel && (
                                <p>
                                    Active X: {store.activePixel.x} - Active Y: {store.activePixel.y}
                                </p>
                            )}
                            {store.hoverPixel && (
                                <p>
                                    Hover X: {store.hoverPixel.x} - Hover Y: {store.hoverPixel.y}
                                </p>
                            )}
                        </div>
                        <div className='web3-status'>
                            {account ? (
                                <>
                                    <Status width={20} height={20} mood='success' />
                                    <p onClick={() => deactivate()}>Web3 connected</p>
                                </>
                            ) : (
                                <>
                                    <Status width={20} height={20} mood='error' />
                                    <p onClick={() => authenticateClick()}>Web3 disconnected</p>
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
                            <div className='color' onClick={() => colorClick()}>
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
                                    <p>
                                        {store.activePixel ? (validateHexCode(store.activePixel.color) ? store.activePixel.color : 'Invalid hex') : 'No pixel selected'}
                                    </p>
                                </div>
                            </div>
                            <div onClick={() => setContextActive(false)} className={'close' + (contextActive ? ' fadeIn' : ' fadeOut')}>
                                <svg xmlns='http://www.w3.org/2000/svg' className='ionicon' viewBox='0 0 512 512'>
                                    <title>Close</title>
                                    <path
                                        fill='none'
                                        stroke='currentColor'
                                        stroke-linecap='round'
                                        stroke-linejoin='round'
                                        stroke-width='32'
                                        d='M368 368L144 144M368 144L144 368'
                                    />
                                </svg>
                            </div>
                        </div>
                        {contextActive && store.activePixel && (
                            <div className='information'>
                                <div className='more-info'>
                                    <p className='caption'>Last editor: </p>
                                    <Address className='value'>{store.activePixel.editor ? store.activePixel.editor : 'None'}</Address>
                                    <p className='caption'>Last edited time: </p>
                                    <p className='value'>{store.activePixel.updatedAt ? parseDate(store.activePixel.updatedAt) : 'None'}</p>
                                    <p className='caption'>Countdown: </p>
                                    <p className='value'>{timeLeft ? timeLeft : 'None'}</p>
                                </div>
                                <div className='editor'>
                                    <div className='input-container'>
                                        <input placeholder='New Color' ref={inputRef} onChange={changePreviewColor} value={store.activePixel.color} />
                                        {isPreview && <button onClick={() => resetPreviewColor()}>Reset</button>}
                                    </div>
                                    {account ? (
                                        <button onClick={() => changeColor()} disabled={!canEditColor}>
                                            {isChangingColor && <Spinner width={20} height={20} />}
                                            Change Color
                                        </button>
                                    ) : (
                                        <button onClick={() => authenticateClick()}>
                                            <div className='auth-button'>
                                                {authenticating ? (
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
                                                    <>Sign in</>
                                                )}
                                            </div>
                                        </button>
                                    )}
                                </div>
                            </div>
                        )}
                    </div>
                </>
            )}
            {children}
        </div>
    );
};

export default observer(Haptics);
