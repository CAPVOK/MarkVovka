/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useEffect, useRef, useState } from "react";
import "./Console.scss";
import { useSelector } from "../../core";
import { ChangeEvent } from "../../App.typig";
import { selectConsoleMessages } from "../../core/slices/console";

import { CloseIcon, TopArrowIcon } from "../icons";

export const Console: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [isConsoleActive, setIsConsoleActive] = useState(false);
  const [message, setMessage] = useState("");

  const consoleRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  const messages = useSelector(selectConsoleMessages);

  const handleInputChange = (event: ChangeEvent) => {
    setMessage(event.target.value);
  };

  const handleConsoleArrowClick = () => {
    setIsOpen(!isOpen);
    inputRef.current?.focus();
  }

  const handleDocumentClick = (event: MouseEvent) => {
    // Если клик произошел вне компонента, убираем фокус
    if (!event.target || !consoleRef.current || !consoleRef.current.contains(event.target as Node)) {
      inputRef.current?.blur();
    } else {
      inputRef.current?.focus();
    }
  };
  useEffect(() => {
    document.addEventListener('click', handleDocumentClick);

    return () => {
      document.removeEventListener('click', handleDocumentClick);
    };
  }, []);

  useEffect(() => {
    const messagesContainer = consoleRef.current?.querySelector('.main_block');

    if (messagesContainer) {
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
  }, [messages]);

  return (
    <div className="console bordert">
      <div className="header">
        <h2>Консоль</h2>
        <div className="manage_buttons">
          <button className="btn">
            <div className="icon close">
              <CloseIcon fill={"black"} />
            </div>
          </button>
          <button className="btn" onClick={handleConsoleArrowClick}>
            <div className={isOpen ? "icon reverse" : "icon"}>
              <TopArrowIcon />
            </div>
          </button>
        </div>
      </div>
      <div ref={consoleRef} className={isOpen ? "main_block" : "main_block_close"}>
        <div className="messages">
          {messages.length > 0 &&
            messages.map((message, id) => (
              <div key={id} className="console_message">
                {message}
              </div>
            ))}
        </div>
        <div className="cursor">
          <input
            ref={inputRef}
            id="console_input"
            type="text"
            value={message}
            onChange={handleInputChange}
            className="console_message console_input"
          />
          <i></i>
        </div>
      </div>
    </div>
  );
};
