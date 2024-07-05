export type ISignInPayload = {
  email: string;
  password: string;
};

export type ISignUpPayload = {
  name: string;
  email: string;
  password: string;
  repeat_password: string;
};

export type IRecoveryPasswordPayload = {
  email: string;
  password: string;
};

export type ISignInResponseResult = {
  id: string;
  email: string;
  name: string;
  created_at: number;
  updated_at: number;
};

export type ISignInResponse = {
  mensaje: string;
  result: ISignInResponseResult;
};

export default interface IUseAuthentication {
  signin(jwt: string): Promise<ISignInResponse>;
  singup(payload: ISignUpPayload): Promise<void>;
  recoveryPassword(payload: IRecoveryPasswordPayload): Promise<void>;
}
