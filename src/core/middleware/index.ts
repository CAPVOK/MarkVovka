import { Middleware } from "@reduxjs/toolkit";
import "regenerator-runtime/runtime";

export const BASE_URL = "ws://localhost:8081"; // Укажите свой адрес сервера

const initialization = (socket: WebSocket) => {
  socket.addEventListener("open", () => {
    console.log("Connection established :)");
  });
  socket.addEventListener("close", () => {
    console.log("Socket disconnected!");
  });
  socket.addEventListener("message", (event) => {
    const data = JSON.parse(event.data);
    console.log("Received message:", data);
  });
  socket.addEventListener("error", (error) => {
    console.error("WebSocket error:", error);
  });
};

const socket = new WebSocket(BASE_URL);

initialization(socket);

export const socketMiddleware: Middleware =
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  (store) => (next) => (action) => {
    switch (action.type) {
      default:
        return next(action);
    }
  };
