import React, { useEffect, useState } from "react";
import "./MainPage.scss";
import { Photo } from "../../componets";
import { useSelector } from "../../core";
import { selectStationInfo } from "../../core/slices/station";

export const MainPage: React.FC = () => {
  const info = useSelector(selectStationInfo);
  const [position, setPosition] = useState<{
    latitude: number;
    longitude: number;
    altitude: number;
  }>({ latitude: 0, longitude: 0, altitude: 0 });
  useEffect(() => {
    if (info.longitude && !position.longitude) {
      setPosition({
        latitude: info.latitude,
        longitude: info.longitude,
        altitude: info.altitude,
      });
    }
  }, [info]);

  function convertCoordinatesToDMS(
    lat: number,
    lon: number
  ): { lat: string; lon: string } {
    const formatCoordinate = (
      value: number,
      positiveSuffix: string,
      negativeSuffix: string
    ): string => {
      const absValue = Math.abs(value);
      const degrees = Math.floor(absValue);
      const minutes = Math.floor((absValue - degrees) * 60);

      const direction = value >= 0 ? positiveSuffix : negativeSuffix;

      return `${degrees}° ${minutes}' ${direction}`;
    };
    const lonInfo = formatCoordinate(lon, "в. д.", "з. д.");
    const latInfo = formatCoordinate(lat, "с. ш.", "ю. ш.");

    return { lat: latInfo, lon: lonInfo };
  }
  return (
    <div className="mainPage">
      <div className="photo_block">
        <div className="header">
          <h1>Россия-1</h1>
          <p>Россия-1: Вместе в космическом полете будущего!</p>
        </div>
        <Photo />
        <div className="info_block">
          <div className="info">
            <p className="label">Широтa</p>
            <p className="data">
              {
                convertCoordinatesToDMS(position.latitude, position.longitude)
                  .lat
              }
            </p>
          </div>
          <div className="info">
            <p className="label">Долгота</p>
            <p className="data">
              {
                convertCoordinatesToDMS(position.latitude, position.longitude)
                  .lon
              }
            </p>
          </div>
          <div className="info">
            <p className="label">Высота</p>
            <p className="data">{info.altitude || "неизвестно"} км</p>
          </div>
        </div>
      </div>
    </div>
  );
};
