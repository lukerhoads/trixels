import React, { CSSProperties } from 'react';
import GreenStatus from '../../assets/icons/GreenStatus.svg';
import RedStatus from '../../assets/icons/RedStatus.svg';
import 'styles/status-icon.scss';

export type StatusProps = {
  width: number;
  height: number;
  green: boolean;
};

const Status = ({ width, height, green }: StatusProps) => {
  const alt = green ? 'Green Status' : 'Red Status';
  const src = green ? GreenStatus : RedStatus;

  return (
    <img
      className='status-icon'
      src={src}
      alt={alt}
      style={{
        width: width,
        height: height,
      }}
    />
  );
};

export default Status;
