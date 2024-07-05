import CryptoJS from "crypto-js";

export default function AES() {
  function encrypt(data: string, password?: string) {
    if (password === undefined) {
      password = process.env.AES_PASSWORD as string;
    }

    const ciphertext = CryptoJS.AES.encrypt(data, password).toString();
    return ciphertext;
  }

  //cifrar objeto
  function encryptObject(data: any, password?: string) {
    if (password === undefined) {
      password = process.env.AES_PASSWORD as string;
    }

    const ciphertext = CryptoJS.AES.encrypt(
      JSON.stringify(data),
      password
    ).toString();
    return ciphertext;
  }

  function decrypt(data: string, password?: string) {
    if (password === undefined) {
      password = process.env.AES_PASSWORD as string;
    }

    const bytes = CryptoJS.AES.decrypt(data, password);
    const originalText = bytes.toString(CryptoJS.enc.Utf8);
    return originalText;
  }

  //descifrar objeto
  function decryptObject(data: string, password?: string) {
    if (password === undefined) {
      password = process.env.AES_PASSWORD as string;
    }

    const bytes = CryptoJS.AES.decrypt(data, password);
    const originalText = bytes.toString(CryptoJS.enc.Utf8);
    return JSON.parse(originalText);
  }

  return {
    encrypt,
    encryptObject,
    decrypt,
    decryptObject,
  };
}
