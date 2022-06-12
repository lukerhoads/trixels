import { observer } from "mobx-react";
import store from "../store";

export type AddressProps = React.HTMLAttributes<HTMLParagraphElement> & {
    children: React.ReactNode
}

const Address = ({ children, className }: AddressProps) => {
    const addressClick = (e: React.MouseEvent<HTMLParagraphElement>) => {
        navigator.clipboard.writeText(e.currentTarget.innerText);
        store.pushToLogs({
            mood: 'success',
            message: 'Copied address.',
        });
    };

    return (
        <p className={className + " clicktext"} onClick={addressClick}>
            {children}
        </p>
    )
}

export default observer(Address)