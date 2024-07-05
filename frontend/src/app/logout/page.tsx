"use client";

import { logout } from "@/app/lib/actions";
import { useEffect } from "react";
import { useFormState } from "react-dom";

export default function Page() {
    const [errorMessage, dispatch] = useFormState(logout, undefined);

    useEffect(() => {
        dispatch();
    }, []);

    return (
        <div className="flex flex-col items-center justify-center h-screen w-full">
            <h1 className="text-4xl">Logging out...</h1>
        </div>
    );
}
