import React from 'react';

export type CloseProps = {
    text?: string;
    onExit: (e: React.MouseEvent<HTMLDivElement>) => void;
};

const Close = ({ text, onExit }: CloseProps) => {
    return (
        <div className='overlay-header'>
            <div>{text && <p>{text}</p>}</div>
            <div className='close close-overlay' onClick={onExit}>
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
    );
};

export default Close;
