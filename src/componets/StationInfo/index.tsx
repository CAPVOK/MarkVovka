import { useSelector } from "../../core";
import { selectStationInfo } from "../../core/slices/station";
import "./StationInfo.scss";
import { FlySatelliteIcon } from "../icons/SatelliteIcon";
import { useTheme } from "../../ThemeProvider";

function StationInfo() {
  const theme = useTheme();
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
  // Пример использования
  const result = convertCoordinatesToDMS(info.latitude, info.longitude);
  console.log(result);

  return (
    <div className="station_data">
      <h1>Россия-1</h1>
      <div className="station_info">
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
            <p className="data">{info.altitude} км</p>
          </div>
          <div className="info">
            <p className="label">Скорость</p>
            <p className="data">{info.speed} км/час</p>
          </div>
          <div className="info">
            <p className="label">Угол</p>
            <p className="data">{info.angle}°</p>
          </div>
          <div className="info">
            <p className="label">Солнечные панели</p>
            <p className="data">{info.status ? "активны" : "убраны"}</p>
          </div>
        </div>
        <div className="planet_block">
          <div
            className="planet"
            style={{ transform: `rotate(${180 - info.angle || 0}deg)` }}
          >
            <div className="orbit">
              <div className="sputnic">
                <FlySatelliteIcon fill={theme?.accentColor} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export { StationInfo };
