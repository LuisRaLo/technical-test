export default class StringsHelper {
  public static isAlphaNumeric(str: string): boolean {
    const alphaNumericRegex = /^[a-zA-Z0-9]+$/;
    const isValid = alphaNumericRegex.test(str);
    return isValid;
  }

  public static isEmpty(str: string): boolean {
    return str === "" || str === null || str === undefined;
  }

  public static isSafeString(str: string): boolean {
    const safeStringRegex = "^[a-zA-Z0-9!@#$%^&*)(+=._-]+$";
    const isValid = new RegExp(safeStringRegex).test(str);

    return isValid;
  }
}
