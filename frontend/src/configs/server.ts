interface ServerConfig {
  serviceAccount: {
    projectId: string;
    clientEmail: string;
    privateKey: string;
  };
}

export const serverConfig: ServerConfig = {
  serviceAccount: {
    projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID!,
    clientEmail: process.env.FIREBASE_ADMIN_CLIENT_EMAIL!,
    privateKey: process.env.FIREBASE_ADMIN_PRIVATE_KEY?.replace(/\\n/g, "\n")!,
  },
};

interface ClientConfig {
  apiKey: string;
  authDomain: string;
  projectId: string;
  messagingSenderId: string;
  cookieName: string;
  cookieSignatureKeys: string[];
  cookieSerializeOptions: {
    path: string;
    httpOnly: boolean;
    secure: boolean;
    sameSite: "strict" | "lax" | "none";
    maxAge: number;
  };
}

export const clientConfig: ClientConfig = {
  apiKey: import.meta.env.VITE_PUBLIC_FIREBASE_API_KEY as string,
  authDomain: import.meta.env.VITE_PUBLIC_FIREBASE_AUTH_DOMAIN as string,
  projectId: import.meta.env.VITE_PUBLIC_FIREBASE_PROJECT_ID as string,
  messagingSenderId: import.meta.env
    .VITE_PUBLIC_FIREBASE_MESSAGING_SENDER_ID as string,

  cookieName: import.meta.env.VITE_AUTH_COOKIE_NAME!,
  cookieSignatureKeys: [
    import.meta.env.VITE_AUTH_COOKIE_SIGNATURE_KEY_CURRENT!,
    import.meta.env.AUTH_COOKIE_SIGNATURE_KEY_PREVIOUS!,
  ],
  cookieSerializeOptions: {
    path: "/",
    httpOnly: true,
    secure: import.meta.env.VITE_USE_SECURE_COOKIES === "true",
    sameSite: "lax" as const,
    maxAge: 12 * 60 * 60 * 24,
  },
};
