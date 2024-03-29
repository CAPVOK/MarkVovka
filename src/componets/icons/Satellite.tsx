import { FC } from "react";
import { IconType } from "../../App.typig";

export const SatelliteIcon: FC<IconType> = (props) => {
  const { fill } = props;

  return (
    <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
      <g
        id="SVGRepo_tracerCarrier"
        strokeLinecap="round"
        strokeLinejoin="round"
      ></g>
      <g id="SVGRepo_iconCarrier">
        {" "}
        <path
          d="M12 3C16.9706 3 21 7.02944 21 12M12 7C14.7614 7 17 9.23858 17 12M10 14L12 12M10 21C6.13401 21 3 17.866 3 14C3 12.067 3.7835 10.317 5.05025 9.05029L14.9512 18.9512C13.6844 20.2179 11.933 21 10 21Z"
          stroke={fill || "#000000"}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        ></path>{" "}
      </g>
    </svg>
  );
};
