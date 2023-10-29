import React from "react";
import "./ControlPage.scss";
import { useUpdateDataMutation } from "../../core/api/authApi";

import { Console, StationInfo } from "../../componets";

export const ControlPage: React.FC = () => {
  const [updateData] = useUpdateDataMutation();

  return (
    <div className="control_page">
      <StationInfo />
      <button
        onClick={() =>
          updateData({
            speed: 400,
            altitude: 0,
            angle: 0,
          })
        }
      >
        скорость 400
      </button>
      <button
        onClick={() =>
          updateData({
            speed: 40,
            altitude: 0,
            angle: 0,
          })
        }
      >
        скорость 40
      </button>
      <Console />
    </div>
  );
};
