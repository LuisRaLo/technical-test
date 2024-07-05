"use server";

import useAuthentication from "@/api/useAuthentication";
import { ISignInResponse } from "@/domain/authentication";

import { getAuth, signInWithEmailAndPassword, signOut } from "firebase/auth";
import { cookies } from "next/headers";
import { app } from "@/configs/firebase";
import AES from "@/utils/helpers/aesHelper";
import { RedirectType, redirect } from "next/navigation";
import { serverConfig } from "@/configs/server";

export async function authenticate(_currentState: unknown, formData: FormData) {
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;

  const { signin } = useAuthentication();

  try {
    const auth = getAuth(app);
    const { user } = await signInWithEmailAndPassword(auth, email, password);
    const token = await user.getIdToken();
    const response = await signin(token);

    const { result }: ISignInResponse = response;
    const encryptedSessionData = AES().encrypt(
      JSON.stringify({
        data: result,
        token,
      })
    );

    cookies().set(
      serverConfig.cookieName,
      encryptedSessionData,
      serverConfig.cookieSerializeOptions
    );

    return "success";
  } catch (error: any) {
    console.error(error);
    return "Invalid email or password.";
  }
}

export async function logout() {
  const auth = getAuth(app);
  await signOut(auth);

  cookies().delete(serverConfig.cookieName);

  redirect("/login", RedirectType.replace);
}
