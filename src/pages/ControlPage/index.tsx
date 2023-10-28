import React from "react";
import "./ControlPage.scss";

import { Console, StationInfo } from "../../componets"

export const ControlPage: React.FC = () => {
  return <div className="control_page">
    <StationInfo />
    <Console/>
  </div>;
}
