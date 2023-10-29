import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { BASE_URL } from "../middleware";

export interface IStationData {
  speed: number;
  altitude: number;
  angle: number;
}

export interface IConsoleResponse {
  log: string;
  msg: string;
}

export const stationApi = createApi({
  reducerPath: "stationAPI",
  baseQuery: fetchBaseQuery({
    baseUrl: BASE_URL,
    credentials: "same-origin",
    mode: "no-cors",
  }),
  endpoints: (build) => ({
    updateStationData: build.mutation<unknown, IStationData>({
      query: (body) => ({
        url: "/update",
        method: "POST",
        body,
      }),
    }),
    sendConsoleCommand: build.mutation<IConsoleResponse, string>({
      query: (body) => ({
        url: "/console",
        method: "POST",
        body: { message: body },
      }),
    }),
    getPhoto: build.query<unknown, void>({
      query: () => ({
        url: "/sector-image",
      }),
    }),
  }),
});

export const { useUpdateStationDataMutation, useSendConsoleCommandMutation, useGetPhotoQuery } =
  stationApi;
