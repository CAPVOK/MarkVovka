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
  solarPanelStatus: string;
  scientificInstrumentsStatus: string;
  navigationSystemStatus: string;
  temperature: number;
  hullStatus: string;
  fuelLevel: number;
}

const initialState: IStation = {
  latitude: 0,
  longitude: 0,
  speed: 0,
  altitude: 0,
  planetRadius: 0,
  angle: 0,
  planetName: "",
  scientificInstrumentsStatus: "",
  navigationSystemStatus: "",
  solarPanelStatus: "",
  temperature: 0,
  hullStatus: "",
  fuelLevel: 0,
};

const slice = createSlice({
  name: "stationSlice",
  initialState,
  reducers: {
    updateInfo: (state, action: PayloadAction<IStation>) => {
      state.latitude = action.payload.latitude || 0;
      state.longitude = action.payload.longitude || 0;
      state.speed = action.payload.speed || 0;
      state.altitude = action.payload.altitude || 0;
      state.planetName = action.payload.planetName || "";
      state.angle = action.payload.angle || 0;
      state.solarPanelStatus = action.payload.solarPanelStatus || "";
      state.planetName = action.payload.planetName || "";
      state.scientificInstrumentsStatus = action.payload.scientificInstrumentsStatus || "";
      state.navigationSystemStatus = action.payload.navigationSystemStatus || "";
      state.hullStatus = action.payload.hullStatus || "";
      state.temperature = action.payload.temperature || 0;
      state.fuelLevel = action.payload.fuelLevel || 0;
    },
  },
});

export const { updateInfo } = slice.actions;

export const selectStationInfo = (state: RootState) => state.station;

export const { reducer } = slice;
