import React from "react";
import "./LoadingConnection.scss";

export const LoadingConnection: React.FC = () => {
  return (
    <div className="loading_connection">
      <div className="content">
        <div className="planet">
          <div className="ring"></div>
          <div className="cover-ring"></div>
          <div className="spots">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>
        <p>Подключение к станции</p>
      </div>
    </div>
  );
};
