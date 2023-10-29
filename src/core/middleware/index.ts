import { Middleware } from "@reduxjs/toolkit";
import "regenerator-runtime/runtime";
import { store } from "..";
import { setSocketConnection } from "../slices/common";
import { IStation, updateInfo } from "../slices/station";

export const BASE_URL = "http://localhost:8080";
const Socket_URL = "ws://localhost:8080/data";

const initialization = (socket: WebSocket) => {
  socket.addEventListener("open", () => {
    console.log("Connection established :)");
    store.dispatch(setSocketConnection());
  });
  socket.addEventListener("close", () => {
    console.log("Socket disconnected!");
  });
  socket.addEventListener("message", (event) => {
    const data: IStation = JSON.parse(event.data);
    store.dispatch(updateInfo(data));
  });
  socket.addEventListener("error", (error) => {
    console.error("WebSocket error:", error);
  });
};

const socket = new WebSocket(Socket_URL);

initialization(socket);

export const socketMiddleware: Middleware =
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  (store) => (next) => (action) => {
    switch (action.type) {
      default:
        return next(action);
    }
  };
