"use client";

import { authenticate } from "@/app/lib/actions";

import { useEffect } from "react";
import { useFormState, useFormStatus } from "react-dom";

export default function Page() {
    const [errorMessage, dispatch] = useFormState(authenticate, undefined);
    const { pending } = useFormStatus();

    const handleClick = (event: any) => {
        if (pending) {
            event.preventDefault();
        }
    };

    useEffect(() => {
        if (errorMessage === "success") {
            window.location.href = "/";
        }
    }, [errorMessage]);

    return (
        <form
            action={dispatch}
            className="flex flex-col items-center justify-center h-screen w-full"
        >
            <input
                type="email"
                name="email"
                placeholder="Email"
                className="mb-4 w-64 p-2 rounded border"
                required
            />
            <input
                type="password"
                name="password"
                placeholder="Password"
                className="mb-4 w-64 p-2 rounded border"
                required
            />
            <div className="mb-4 w-64 p-2 rounded" aria-live="polite" role="alert">
                {errorMessage && <p>{errorMessage}</p>}
            </div>
            <button
                aria-disabled={pending}
                type="submit"
                onClick={handleClick}
                className="bg-blue-500 text-white p-2 rounded w-64"
            >
                Login
            </button>
        </form>
    );
}
