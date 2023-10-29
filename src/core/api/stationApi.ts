import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { BASE_URL } from "../middleware";

export interface IStationData {
  speed: number;
  altitude: number;
  angle: number;
}

export const authApi = createApi({
  reducerPath: "userAPI",
  baseQuery: fetchBaseQuery({
    baseUrl: BASE_URL,
    credentials: "same-origin",
    mode: "no-cors",
  }),
  tagTypes: ["Station"],
  endpoints: (build) => ({
    updateStationData: build.mutation<unknown, IStationData>({
      query: (body) => ({
        url: "/update",
        method: "POST",
        body
      }),
      invalidatesTags: [{ type: "Station" }],
    }),
  }),
});

export const { useUpdateStationDataMutation } = authApi;
