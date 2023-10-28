import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "..";

export interface IStation {
  latitude: number;
  longitude: number;
  speed: number;
  altitude: number;
  planetRadius: number;
  angle: number;
  planetName: string;
  status: string;
}

const initialState: IStation = {
  latitude:  50.123,
  longitude: 30.456,
  speed: 200,
  altitude: 300,
  planetRadius: 6371,
  angle: 45,
  planetName: "Earth",
  status: "active"
};

const slice = createSlice({
  name: "stationSlice",
  initialState,
  reducers: {
    /* logInUser: (state, action: PayloadAction<string>) => {
      state.isUserLogIn = true
      state.userName = action.payload
    },
    logOutUser: (state) => {
      state.isUserLogIn = false
      state.userName = ""
    }, */
  },
});

// eslint-disable-next-line no-empty-pattern
export const {
  /* logInUser, logOutUser */
} = slice.actions;

export const selectStationInfo = (state: RootState) =>
  state.station;

export const { reducer } = slice;
