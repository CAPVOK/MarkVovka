import { createSlice } from "@reduxjs/toolkit"
import { RootState } from "..";

const initialState = {
  isSocketConnected: false,
}

const slice = createSlice({
  name: "connectionState",
  initialState,
  reducers: {
    setSocketConnection: (state) => {
      state.isSocketConnected = true;
    },
  },
})

export const { setSocketConnection } = slice.actions

export const selectConnection = (state: RootState) => state.connection.isSocketConnected;

export const { reducer } = slice
