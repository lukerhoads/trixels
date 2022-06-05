import React, { CSSProperties } from 'react';
import 'styles/spinner.scss';

export type SpinnerProps = {
  width: number;
  height: number;
  extraStyle?: CSSProperties;
};

export const Spinner = ({ width, height, extraStyle }: SpinnerProps) => {
  return (
    <div className='spinner-container'>
      <div
        className='spinner'
        style={{
          width: width,
          height: height,
          ...extraStyle,
        }}
      />
    </div>
  );
};
