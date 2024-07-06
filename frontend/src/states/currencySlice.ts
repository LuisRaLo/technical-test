import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "./store";

export enum CurrencyEnum {
  MXN = "MXN",
  USD = "USD",
}

interface ICurrencyState {
  currency: CurrencyEnum;
}

const initialState: ICurrencyState = {
  currency: CurrencyEnum.MXN,
};

export const currencySlice = createSlice({
  name: "currency",
  initialState,
  reducers: {
    setCurrency: (state, action: PayloadAction<CurrencyEnum>) => {
      state.currency = action.payload;
    },
  },
});

export const { setCurrency } = currencySlice.actions;
export const selectCurrency = (state: RootState) => state.currency;
export default currencySlice.reducer;
