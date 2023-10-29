/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useEffect, useRef, useState } from "react";
import "./Console.scss";
import { useDispatch, useSelector } from "../../core";
import { ChangeEvent } from "../../App.typig";
import {
  selectConsoleMessages,
  updateMessages,
} from "../../core/slices/console";

import { TopArrowIcon } from "../icons";
import { useTheme } from "../../ThemeProvider";
import { IConsoleResponse } from "../../core/api/stationApi";

export const Console: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [message, setMessage] = useState("");
  const theme = useTheme();

  const consoleRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  const dispatch = useDispatch();
  const messages = useSelector(selectConsoleMessages);

  const sendConsoleCommand = async (message: string) => {
    try {
      const response = await fetch("http://localhost:8080/console", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ message }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data: IConsoleResponse = await response.json();
      return data;
    } catch (error) {
      console.error("Error:", error);
      throw error;
    }
  };

  const handleSendCommand = () => {
    console.log(message);
    sendConsoleCommand(message).then((data) => {
      dispatch(updateMessages(data));
    });
    setMessage("");
  };

  const handleInputChange = (event: ChangeEvent) => {
    setMessage(event.target.value);
  };

  const handleConsoleArrowClick = () => {
    setIsOpen(!isOpen);
    inputRef.current?.focus();
  };

  const handleConsoleClick = (event: MouseEvent) => {
    if (
      !event.target ||
      !consoleRef.current ||
      !consoleRef.current.contains(event.target as Node)
    ) {
      inputRef.current?.blur();
    } else {
      inputRef.current?.focus();
    }
  };

  /* const handleEnterPress = (event: KeyboardEvent) => {
    if (
      event.key === "Enter" &&
      event.target &&
      consoleRef.current &&
      consoleRef.current.contains(event.target as Node)
    ) {
      handleSendCommand();
    }
  }; */
  const handleEnterPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      event.preventDefault();
      handleSendCommand();
    }
  };
  useEffect(() => {
    document.addEventListener("click", handleConsoleClick);
    /* document.addEventListener("keydown", handleEnterPress); */
    return () => {
      document.removeEventListener("click", handleConsoleClick);
      /* document.removeEventListener("keydown", handleEnterPress); */
    };
  }, []);

  useEffect(() => {
    const messagesContainer = consoleRef.current?.querySelector(".main_block");
    if (messagesContainer) {
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
  }, [messages]);

  return (
    <div className="console bordert">
      <div className="header">
        <h2>Консоль</h2>
        <div className="manage_buttons">
          <button className="btn" onClick={handleConsoleArrowClick}>
            <div className={isOpen ? "icon reverse" : "icon"}>
              <TopArrowIcon fill={theme?.textColor} />
            </div>
          </button>
        </div>
      </div>
      <div
        ref={consoleRef}
        className={isOpen ? "main_block" : "main_block_close"}
      >
        <div className="messages">
          {messages.length > 0 &&
            messages.map((message, id) => (
              <div key={id} className="console_message">
                <div className="log">{message.log}</div>
                <div className="msg">{message.msg}</div>
              </div>
            ))}
        </div>
        <div className="cursor">
          <input
            ref={inputRef}
            id="console_input"
            placeholder="$"
            type="text"
            value={message}
            onChange={handleInputChange}
            onKeyDown={handleEnterPress}
            className="console_input"
          />
          <i></i>
        </div>
      </div>
    </div>
  );
};
