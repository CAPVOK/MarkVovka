import { createSlice } from "@reduxjs/toolkit"

const initialState = {
  isSocketConnected: false,
}

const slice = createSlice({
  name: "clusterState",
  initialState,
  reducers: {
    setSocketConnection: (state) => {
      state.isSocketConnected = true;
    },
  },
})

export const { setSocketConnection } = slice.actions

export const { reducer } = slice
