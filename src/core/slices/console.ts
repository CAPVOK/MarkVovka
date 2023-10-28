import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "..";

export interface IConsoleMessages {
  messages: string[];
}

const initialState: IConsoleMessages = {
  messages: [
    "[ИНФО] 02.11.2023 12:00: Инициализация систем. Подготовка к запуску.",
    "[ИНФО] 02.11.2023 12:05: Система навигации активирована. Определение текущих координат.",
    "[ИНФО] 02.11.2023 12:10: Запуск двигателя. Переход в режим орбитального движения.",
    "[ИНФО] 02.11.2023 12:15: Мониторинг высоты. Проверка стабильности орбиты.",
    "[ИНФО] 02.11.2023 12:20: Коррекция ориентации. Выравнивание по широте и долготе.",
    "[ВАЖНО] 02.11.2023 12:25: ПОЛОМКА: Аварийное отключение двигателя. Активация системы безопасности.",
    "[ИНФО] 02.11.2023 12:30: Автоматическое восстановление. Проверка всех систем.",
    "[ИНФО] 02.11.2023 12:35: Получение данных о солнечной активности. Адаптация к условиям космоса.",
    "[ИНФО] 02.11.2023 12:40: Активация научных инструментов. Начало сбора данных.",
    "[ИНФО] 02.11.2023 12:45: Контроль энергопотребления. Оптимизация работы систем.",
    "[ИНФО] 02.11.2023 12:50: Обновление программного обеспечения. Повышение эффективности миссии.",
    "[ВАЖНО] 02.11.2023 12:55: Сбой в связи с солнечным излучением. Автоматическое переключение на резервный источник энергии.",
    "[ИНФО] 02.11.2023 13:00: Связь с земным контролем установлена. Передача собранных данных.",
  ],
};

const slice = createSlice({
  name: "consoleSlice",
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

export const selectConsoleMessages = (state: RootState) => state.console.messages;

export const { reducer } = slice;
