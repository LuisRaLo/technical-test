import IUseAuthentication, {
  IRecoveryPasswordPayload,
  ISignInResponse,
  ISignUpPayload,
} from "@/domain/authentication";

export default function useAuthentication(): IUseAuthentication {
  async function signin(jwt: string): Promise<ISignInResponse> {
    const url: string = process.env.URL_SERVICES + "/api/v1/users/byJWT";

    const headers: HeadersInit = {
      authorization: "Bearer " + jwt,
      "Content-Type": "application/json",
    };

    const req: Response = await fetch(url, {
      method: "GET",
      headers: headers,
    });

    const res = await req.json();

    if (req.status > 200) {
      throw new Error(res.message);
    }

    return res;
  }

  async function singup(payload: ISignUpPayload): Promise<void> {}

  async function recoveryPassword(
    payload: IRecoveryPasswordPayload
  ): Promise<void> {
    throw "not implemented";
  }

  return {
    signin,
    singup,
    recoveryPassword,
  };
}
