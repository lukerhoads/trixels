import 'styles/status-icon.scss';
import { Mood } from 'types/log';
import GreenStatus from '../../assets/icons/GreenStatus.svg';
import OrangeStatus from '../../assets/icons/OrangeStatus.svg';
import RedStatus from '../../assets/icons/RedStatus.svg';

export type StatusProps = {
    width: number;
    height: number;
    mood: Mood;
};

const Status = ({ width, height, mood }: StatusProps) => {
    let alt;
    let src;
    if (mood == 'error') {
        alt = 'Red Status';
        src = RedStatus;
    } else if (mood == 'success') {
        alt = 'Green Status';
        src = GreenStatus;
    } else {
        alt = 'Orange Status';
        src = OrangeStatus;
    }

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
