import { useEffect } from "react";
import { useSelector } from "../../core";
import { selectStationInfo } from "../../core/slices/station";
import "./StationInfo.scss";

function StationInfo() {
  const info = useSelector(selectStationInfo);

  useEffect(() => {

  }, [info]);

  function convertCoordinatesToDMS(lat: number, lon: number): string {
    const formatCoordinate = (value: number, positiveSuffix: string, negativeSuffix: string): string => {
      const absValue = Math.abs(value);
      const degrees = Math.floor(absValue);
      const minutes = Math.floor((absValue - degrees) * 60);
  
      const direction = value >= 0 ? positiveSuffix : negativeSuffix;
  
      return `${degrees}° ${minutes}' ${direction}`;
    };
    const lonInfo = formatCoordinate(lon, 'в.д.', 'з.д.');
    const latInfo = formatCoordinate(lat, 'с.ш.', 'ю.ш.');
  
    return `${latInfo}\n${lonInfo}`;
  }
  // Пример использования
  const result = convertCoordinatesToDMS(info.latitude, info.longitude);
  console.log(result);
  

  return (
    <div className="station_info">

    </div>
  );
}

export { StationInfo };
