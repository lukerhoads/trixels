import { observer } from 'mobx-react';
import React, { ReactNode, useEffect, useState } from 'react';
import { HexColorPicker } from 'react-colorful';
import store from '../store';
import '../styles/overlay.scss';
import AuthModal from './AuthModal';
import Close from './Close';

const Overlay = () => {
    const [object, setObject] = useState<ReactNode | undefined>(undefined);
    const [color, setColor] = useState(store.activePixel?.color);

    useEffect(() => {
        if (!color) return;
        store.setActiveColor(color.toUpperCase());
    }, [color]);

    useEffect(() => {
        if (!store.activePixel) return;
        setColor(store.activePixel.color);
    }, [store.activePixel?.color]);

    const onExit = (e: React.MouseEvent<HTMLDivElement>) => {
        e.stopPropagation();
        store.setOverlayInfo(undefined);
    };

    useEffect(() => {
        if (store.overlayInfo) {
            if (store.overlayInfo) {
                switch (store.overlayInfo.type.valueOf()) {
                    case 'color-picker':
                        setObject(
                            <div className='color-picker'>
                                <Close onExit={onExit} />
                                <HexColorPicker color={color} onChange={setColor} />
                            </div>
                        );
                        return;
                    case 'auth-modal':
                        setObject(<AuthModal onExit={onExit} />);
                        return;
                    default:
                        setObject(undefined);
                        return;
                }
            }
        }
    }, [store.overlayInfo]);

    return (
        <div
            className='overlay'
            style={{
                display: store.overlayInfo ? 'block' : 'none',
                backdropFilter: store.overlayInfo ? 'blur(5px)' : '',
            }}
            onClick={onExit}
        >
            {object}
        </div>
    );
};

export default observer(Overlay);
