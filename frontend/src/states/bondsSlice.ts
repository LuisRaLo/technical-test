import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "./store";

// Define a type for the slice state
interface IBondState {
  myBonds: any;
}

// Define the initial state using that type
const initialState: IBondState = {
  myBonds: [],
};

export const bondSlice = createSlice({
  name: "bond",
  initialState,
  reducers: {
    addBond: (state, action: PayloadAction<any>) => {
      state.myBonds.push(action.payload);
    },
    removeBond: (state, action: PayloadAction<number>) => {
      state.myBonds = state.myBonds.filter(
        (bond: any) => bond.id !== action.payload
      );
    },
  },
});

export const { addBond, removeBond } = bondSlice.actions;
export const selectBonds = (state: RootState) => state.bond;
export default bondSlice.reducer;
