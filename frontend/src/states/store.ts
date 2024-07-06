import { configureStore } from "@reduxjs/toolkit";
import currencyReducer from "./currencySlice";
import bondReducer from "./bondsSlice";

export const AppStore = configureStore({
  reducer: {
    currency: currencyReducer,
    bond: bondReducer,
  },
});

export type RootState = ReturnType<typeof AppStore.getState>;
export type AppDispatch = typeof AppStore.dispatch;
