import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { RootState } from "..";

export interface IConsoleMessage {
  log: string;
  msg: string;
}
export interface IConsole {
  messages: IConsoleMessage[];
}

const initialState: IConsole = {
  /* messages: [
    {
      log: "[ИНФО] 02.11.2023 12:00:",
      msg: "Инициализация систем. Подготовка к запуску.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:05:",
      msg: "Система навигации активирована. Определение текущих координат.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:10:",
      msg: "Запуск двигателя. Переход в режим орбитального движения.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:15:",
      msg: "Мониторинг высоты. Проверка стабильности орбиты.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:20:",
      msg: "Коррекция ориентации. Выравнивание по широте и долготе.",
    },
    {
      log: "[ВАЖНО] 02.11.2023 12:25:",
      msg: "ПОЛОМКА: Аварийное отключение двигателя. Активация системы безопасности.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:30:",
      msg: "Автоматическое восстановление. Проверка всех систем.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:35:",
      msg: "Получение данных о солнечной активности. Адаптация к условиям космоса.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:40:",
      msg: "Активация научных инструментов. Начало сбора данных.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:45:",
      msg: "Контроль энергопотребления. Оптимизация работы систем.",
    },
    {
      log: "[ИНФО] 02.11.2023 12:50:",
      msg: "Обновление программного обеспечения. Повышение эффективности миссии.",
    },
    {
      log: "[ВАЖНО] 02.11.2023 12:55:",
      msg: "Сбой в связи с солнечным излучением. Автоматическое переключение на резервный источник энергии.",
    },
    {
      log: "[ИНФО] 02.11.2023 13:00:",
      msg: "Связь с земным контролем установлена. Передача собранных данных.",
    },
  ], */
  messages: [],
};

const slice = createSlice({
  name: "consoleSlice",
  initialState,
  reducers: {
    updateMessages: (state, action: PayloadAction<IConsoleMessage>) => {
      state.messages.push({
        log: action.payload.log || "",
        msg: action.payload.msg || "",
      });
    },
  },
});

// eslint-disable-next-line no-empty-pattern
export const { updateMessages } = slice.actions;

export const selectConsoleMessages = (state: RootState) =>
  state.console.messages;

export const { reducer } = slice;
