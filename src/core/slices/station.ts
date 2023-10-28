import { createSlice, PayloadAction } from "@reduxjs/toolkit";
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
  latitude:  0,
  longitude: 0,
  speed: 0,
  altitude: 0,
  planetRadius: 0,
  angle: 0,
  planetName: "",
  status: ""
};

const slice = createSlice({
  name: "stationSlice",
  initialState,
  reducers: {
    updateInfo: (state, action: PayloadAction<IStation>) => {
      state.latitude = action.payload.latitude;
      state.longitude = action.payload.longitude;
      state.speed = action.payload.speed;
      state.altitude = action.payload.altitude;
      state.planetName = action.payload.planetName;
      state.angle = action.payload.angle;
      state.status = action.payload.status;
      console.log(action.payload);
    }
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
  updateInfo
} = slice.actions;

export const selectStationInfo = (state: RootState) =>
  state.station;

export const { reducer } = slice;
