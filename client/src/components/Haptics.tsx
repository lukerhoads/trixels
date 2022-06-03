import { useEthers } from '@usedapp/core'
import React, { ReactChild, useEffect, useRef, useState } from 'react'
import 'styles/haptics.scss'
import { Spinner } from './Spinner'

export type HapticsProps = {
    children: ReactChild
}

// TODO:
// - Have inputs and text fade in
// - Add metamask icon to button
// - Add tooltips for addresses
// - Add a direct color palette to edit
// - Add loading to metamask sign in button

export const Haptics = ({ children }: HapticsProps) => {
    const { activateBrowserWallet, account } = useEthers()
    const [isAuthenticating, setIsAuthenticating] = useState(false)
    const [isChangingColor, setIsChangingColor] = useState(false)
    const [contextActive, setContextActive] = useState(false)
    const [canEditColor, setCanEditColor] = useState(false)
    const [lastEditor, setLastEditor] = useState("")
    const [lastEditedTime, setLastEditedTime] = useState("")
    const inputRef = useRef<HTMLInputElement>(null)

    const changeZoom = (factor: number) => {
        console.log("Changing zoom by ", factor)
    }

    const web3SignIn = () => {
        if (isAuthenticating || account) {
            return
        }

        setIsAuthenticating(true)
        activateBrowserWallet()
    }

    const changeColor = () => {
        if (!canEditColor) {
            return 
        }
        
        setIsChangingColor(true)
        let newColor = inputRef.current?.value;
        console.log("New color: ", newColor)
    }

    return (
        <div className="haptics">
            <div className="zoom">
                <div className="plus" onClick={() => changeZoom(0.2)}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="ionicon" viewBox="0 0 512 512"><title>Add</title><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M256 112v288M400 256H112"/></svg>
                </div>
                <div className="minus" onClick={() => changeZoom(-0.2)}>
                    <svg xmlns="http://www.w3.org/2000/svg" className="ionicon" viewBox="0 0 512 512"><title>Remove</title><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M400 256H112"/></svg>
                </div>
            </div>
            <div className={"context" + (contextActive ? " context-active" : "")} onClick={() => setContextActive(true)}>
                <div className="mini-header">
                    <div className="color">
                        <div className="color-circle"></div>
                    </div>
                    <div className="color-description">
                        <div className="text">
                            <p>#000000</p>
                        </div>
                    </div>
                    {contextActive ? (
                        <div className="close" onClick={() => setContextActive(false)}>
                            <svg xmlns="http://www.w3.org/2000/svg" className="ionicon" viewBox="0 0 512 512"><title>Close</title><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M368 368L144 144M368 144L144 368"/></svg>
                        </div>
                    ) : null}
                </div>
                {contextActive ? (
                    <div className="information">
                        <div className="more-info">
                            <p>Last editor: {lastEditor}</p>
                            <p>Last edited time: {lastEditedTime}</p>
                        </div>
                        <div className="editor">
                            <input placeholder="New Color" disabled={canEditColor} ref={inputRef} />
                            {account ? (
                                <button onClick={() => changeColor()}>
                                    {isChangingColor ? (
                                        <Spinner width={20} height={20} />
                                    ) : null}
                                    Change Color
                                </button>
                            ) : (
                                <button onClick={() => web3SignIn()}>
                                    <div className="metamask-button">
                                        { isAuthenticating ? (
                                            <>
                                                <Spinner width={20} height={20} extraStyle={{
                                                    marginRight: '10px'
                                                }} />
                                                Signing in...
                                            </>
                                        ) : (
                                            <>
                                                <svg xmlns="http://www.w3.org/2000/svg" width="212" height="189" viewBox="0 0 212 189"><g fill="none" fill-rule="evenodd"><polygon fill="#CDBDB2" points="60.75 173.25 88.313 180.563 88.313 171 90.563 168.75 106.313 168.75 106.313 180 106.313 187.875 89.438 187.875 68.625 178.875"/><polygon fill="#CDBDB2" points="105.75 173.25 132.75 180.563 132.75 171 135 168.75 150.75 168.75 150.75 180 150.75 187.875 133.875 187.875 113.063 178.875" transform="matrix(-1 0 0 1 256.5 0)"/><polygon fill="#393939" points="90.563 152.438 88.313 171 91.125 168.75 120.375 168.75 123.75 171 121.5 152.438 117 149.625 94.5 150.188"/><polygon fill="#F89C35" points="75.375 27 88.875 58.5 95.063 150.188 117 150.188 123.75 58.5 136.125 27"/><polygon fill="#F89D35" points="16.313 96.188 .563 141.75 39.938 139.5 65.25 139.5 65.25 119.813 64.125 79.313 58.5 83.813"/><polygon fill="#D87C30" points="46.125 101.25 92.25 102.375 87.188 126 65.25 120.375"/><polygon fill="#EA8D3A" points="46.125 101.813 65.25 119.813 65.25 137.813"/><polygon fill="#F89D35" points="65.25 120.375 87.75 126 95.063 150.188 90 153 65.25 138.375"/><polygon fill="#EB8F35" points="65.25 138.375 60.75 173.25 90.563 152.438"/><polygon fill="#EA8E3A" points="92.25 102.375 95.063 150.188 86.625 125.719"/><polygon fill="#D87C30" points="39.375 138.938 65.25 138.375 60.75 173.25"/><polygon fill="#EB8F35" points="12.938 188.438 60.75 173.25 39.375 138.938 .563 141.75"/><polygon fill="#E8821E" points="88.875 58.5 64.688 78.75 46.125 101.25 92.25 102.938"/><polygon fill="#DFCEC3" points="60.75 173.25 90.563 152.438 88.313 170.438 88.313 180.563 68.063 176.625"/><polygon fill="#DFCEC3" points="121.5 173.25 150.75 152.438 148.5 170.438 148.5 180.563 128.25 176.625" transform="matrix(-1 0 0 1 272.25 0)"/><polygon fill="#393939" points="70.313 112.5 64.125 125.438 86.063 119.813" transform="matrix(-1 0 0 1 150.188 0)"/><polygon fill="#E88F35" points="12.375 .563 88.875 58.5 75.938 27"/><path fill="#8E5A30" d="M12.3750002,0.562500008 L2.25000003,31.5000005 L7.87500012,65.250001 L3.93750006,67.500001 L9.56250014,72.5625 L5.06250008,76.5000011 L11.25,82.1250012 L7.31250011,85.5000013 L16.3125002,96.7500014 L58.5000009,83.8125012 C79.1250012,67.3125004 89.2500013,58.8750003 88.8750013,58.5000009 C88.5000013,58.1250009 63.0000009,38.8125006 12.3750002,0.562500008 Z"/><g transform="matrix(-1 0 0 1 211.5 0)"><polygon fill="#F89D35" points="16.313 96.188 .563 141.75 39.938 139.5 65.25 139.5 65.25 119.813 64.125 79.313 58.5 83.813"/><polygon fill="#D87C30" points="46.125 101.25 92.25 102.375 87.188 126 65.25 120.375"/><polygon fill="#EA8D3A" points="46.125 101.813 65.25 119.813 65.25 137.813"/><polygon fill="#F89D35" points="65.25 120.375 87.75 126 95.063 150.188 90 153 65.25 138.375"/><polygon fill="#EB8F35" points="65.25 138.375 60.75 173.25 90 153"/><polygon fill="#EA8E3A" points="92.25 102.375 95.063 150.188 86.625 125.719"/><polygon fill="#D87C30" points="39.375 138.938 65.25 138.375 60.75 173.25"/><polygon fill="#EB8F35" points="12.938 188.438 60.75 173.25 39.375 138.938 .563 141.75"/><polygon fill="#E8821E" points="88.875 58.5 64.688 78.75 46.125 101.25 92.25 102.938"/><polygon fill="#393939" points="70.313 112.5 64.125 125.438 86.063 119.813" transform="matrix(-1 0 0 1 150.188 0)"/><polygon fill="#E88F35" points="12.375 .563 88.875 58.5 75.938 27"/><path fill="#8E5A30" d="M12.3750002,0.562500008 L2.25000003,31.5000005 L7.87500012,65.250001 L3.93750006,67.500001 L9.56250014,72.5625 L5.06250008,76.5000011 L11.25,82.1250012 L7.31250011,85.5000013 L16.3125002,96.7500014 L58.5000009,83.8125012 C79.1250012,67.3125004 89.2500013,58.8750003 88.8750013,58.5000009 C88.5000013,58.1250009 63.0000009,38.8125006 12.3750002,0.562500008 Z"/></g></g></svg>
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
    )
}