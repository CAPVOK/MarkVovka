import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { BASE_URL } from "../middleware";

export interface ISignUpInfo {
  fullName: string;
  userName: string;
  password: string;
  email: string;
}

export interface ILogInInfo {
  userName: string;
  password: string;
}

export const authApi = createApi({
  reducerPath: "userAPI",
  baseQuery: fetchBaseQuery({
    baseUrl: BASE_URL,
    credentials: "same-origin",
    mode: "no-cors",
  }),
  tagTypes: ["Authorization"],
  endpoints: (build) => ({
    logIn: build.mutation<unknown, ILogInInfo>({
      query: (body) => ({
        url: "/auth/login",
        method: "POST",
        body: {
          username: body.userName,
          password: body.password,
        },
      }),
      invalidatesTags: [{ type: "Authorization" }],
    }),
    signUp: build.mutation<unknown, ISignUpInfo>({
      query: (body) => ({
        url: "/auth/signup",
        method: "POST",
        body: {
          full_name: body.fullName,
          email: body.email,
          username: body.userName,
          password: body.password,
        },
      }),
      invalidatesTags: [{ type: "Authorization" }],
    }),
  }),
});

export const { useSignUpMutation, useLogInMutation } = authApi;
