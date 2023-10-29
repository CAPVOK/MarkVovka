import { useSelector } from "../../core";
import { selectStationInfo } from "../../core/slices/station";
import "./StationInfo.scss";

function StationInfo() {
  const info = useSelector(selectStationInfo);

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
    <div className="station_data">
      <h1>Россия-1</h1>
      <div className="info_block">
        <div className="info">
          <p className="label">Широтa</p>
          <p className="data">
            {convertCoordinatesToDMS(info.latitude, info.longitude).lat}
          </p>
        </div>
        <div className="info">
          <p className="label">Долгота</p>
          <p className="data">
            {convertCoordinatesToDMS(info.latitude, info.longitude).lon}
          </p>
        </div>
        <div className="info">
          <p className="label">Высота</p>
          <p className="data">{info.altitude  || "неизвестно"} км</p>
        </div>
        <div className="info">
          <p className="label">Скорость</p>
          <p className="data">{info.speed  || "неизвестно"} км/c</p>
        </div>
        <div className="info">
          <p className="label">Солнечные панели</p>
          <p className="data">
            {info.solarPanelStatus || "неизвестно"}
          </p>
        </div>

        <div className="info">
          <p className="label">Научные инструменты</p>
          <p className="data">
            {info.scientificInstrumentsStatus || "неизвестно"}
          </p>
        </div>

        <div className="info">
          <p className="label">Система навигации Астра</p>
          <p className="data">
            {info.navigationSystemStatus || "неизвестно"}
          </p>
        </div>

        <div className="info">
          <p className="label">Температура</p>
          <p className="data">{info.temperature || "неизвестно"}</p>
        </div>

        <div className="info">
          <p className="label">Топливо</p>
          <p className="data">{info.fuelLevel || "неизвестно"}%</p>
        </div>

        <div className="info">
          <p className="label">Состояние корпуса</p>
          <p className="data">{info.hullStatus || "неизвестно"}</p>
        </div>
      </div>
    </div>
  );
}

export { StationInfo };
