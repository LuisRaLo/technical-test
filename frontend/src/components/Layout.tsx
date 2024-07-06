import { ReactNode } from 'react';
import { HeaderBar } from './HeaderBar';

interface LayoutProps {
    children: ReactNode; // Declaración del tipo de children como ReactNode
}

export function Layout({ children }: LayoutProps): JSX.Element {
    return (
        <>
            <HeaderBar />
            <div>{children}</div>
        </>
    );
}
