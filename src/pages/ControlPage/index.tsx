import React from "react";
import "./ControlPage.scss";
/* import { useUpdateStationDataMutation } from "../../core/api/stationApi"; */

import { Console, StationInfo } from "../../componets";
import { RotatingPlanet } from "../../componets/RotatingPlanet";

export const ControlPage: React.FC = () => {
  /* const [updateData] = useUpdateStationDataMutation(); */

  return (
    <div className="control_page">
      <div className="wrapper">
        <StationInfo />
        <div className="right_block">
          <RotatingPlanet/>
        </div>
      </div>
      <Console />
    </div>
  );
};
