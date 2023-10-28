import { FC } from "react";
import { IconType } from "../../App.typig";

export const FlySatelliteIcon: FC<IconType> = (props) => {
  const { fill } = props;

  return (
    <svg
      fill={fill || "#000000"}
      version="1.1"
      id="Icons"
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      viewBox="0 0 32 32"
      xmlSpace="preserve"
    >
      <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
      <g
        id="SVGRepo_tracerCarrier"
        strokeLinecap="round"
        strokeLinejoin="round"
      ></g>
      <g id="SVGRepo_iconCarrier">
        {" "}
        <g>
          {" "}
          <path d="M20.8,19.8c-2.4,2.4-2.4,6.4,0,8.9c0.2,0.2,0.4,0.3,0.7,0.3s0.5-0.1,0.7-0.3l3-3l2,2c0.2,0.2,0.5,0.3,0.7,0.3 s0.5-0.1,0.7-0.3c0.4-0.4,0.4-1,0-1.4l-2-2l3-3c0.4-0.4,0.4-1,0-1.4C27.3,17.4,23.3,17.4,20.8,19.8z"></path>{" "}
          <path d="M21.7,15.3l-2.8-2.8l1.6-1.6l2.8,2.8c0.2,0.2,0.5,0.3,0.7,0.3s0.5-0.1,0.7-0.3l1.8-1.8l-8.4-8.4l-1.8,1.8 c-0.4,0.4-0.4,1,0,1.4l2.8,2.8l-1.6,1.6L9.7,3.3c-0.4-0.4-1-0.4-1.4,0l-4,4c-0.4,0.4-0.4,1,0,1.4l7.8,7.8l-1.6,1.6l-2.8-2.8 c-0.4-0.4-1-0.4-1.4,0l-1.8,1.8l8.4,8.4l1.8-1.8c0.4-0.4,0.4-1,0-1.4l-2.8-2.8l1.6-1.6l2.8,2.8c0.2,0.2,0.5,0.3,0.7,0.3 s0.5-0.1,0.7-0.3l4-4C22.1,16.3,22.1,15.7,21.7,15.3z"></path>{" "}
          <path d="M1.3,20.3c-0.4,0.4-0.4,1,0,1.4l7,7C8.5,28.9,8.7,29,9,29s0.5-0.1,0.7-0.3l1.8-1.8l-8.4-8.4L1.3,20.3z"></path>{" "}
          <path d="M29.7,8.7c0.4-0.4,0.4-1,0-1.4l-7-7c-0.4-0.4-1-0.4-1.4,0l-1.8,1.8l8.4,8.4L29.7,8.7z"></path>{" "}
        </g>{" "}
      </g>
    </svg>
  );
};
