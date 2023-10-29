import React from "react";
import "./RotatingPlanet.scss";

export const RotatingPlanet: React.FC = () => {
  return (
    <div className="rotating_planet">
      <div className="content">
        <div className="planet">
          <div className="ring"></div>
          <div className="cover-ring"></div>
          {/* <div className="spots">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
          </div> */}
        </div>
      </div>
    </div>
  );
};
