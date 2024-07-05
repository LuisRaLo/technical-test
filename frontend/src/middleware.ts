import { NextRequest } from "next/server";

import { serverConfig } from "./configs/server";
import AES from "./utils/helpers/aesHelper";

const PUBLIC_PATHS = ["/login", "/register", "/forgot-password"];

export async function middleware(request: NextRequest) {
  const { cookieName } = serverConfig;
  const { pathname } = request.nextUrl;

  const currentCookie = request.cookies.get(cookieName)?.value;

  if (currentCookie) {
    if (pathname.startsWith("/login")) {
      return Response.redirect(new URL("/", request.url));
    }
  } else {
    if (!PUBLIC_PATHS.includes(pathname)) {
      return Response.redirect(new URL("/login", request.url));
    }
  }
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)"],
};
