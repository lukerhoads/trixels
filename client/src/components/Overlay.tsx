import { useEthers } from '@usedapp/core';
import { CHAIN_ID } from 'config';
import { observer } from 'mobx-react';
import { ReactNode, useEffect, useState } from 'react';
import { HexColorPicker } from 'react-colorful';
import store from '../store';
import '../styles/overlay.scss';
import AuthModal from './AuthModal';
import Close from './Close';
import ExecuteProposalModal from './dao/ExecuteProposalModal';
import ProposalModal from './dao/ProposalModal';

const Overlay = () => {
    const { chainId } = useEthers();
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

    useEffect(() => {
        if (chainId != CHAIN_ID) {
            store.setOverlayInfo({
                type: 'wrong-network',
            });
        }
    }, [chainId])

    const onExit = () => {
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
                    case 'proposal-modal':
                        setObject(<ProposalModal onExit={onExit} />);
                        return;
                    case 'execute-proposal-modal':
                        setObject(<ExecuteProposalModal onExit={onExit} />);
                        return;
                    case 'wrong-network':
                        setObject(<div>Wrong network</div>);
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
