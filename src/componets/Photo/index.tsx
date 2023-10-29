import React, { useState, useEffect } from "react";
import "./Photo.scss";

export const Photo: React.FC = () => {
  const [imageSrc, setImageSrc] = useState(null);

  useEffect(() => {
    const fetchImage = async () => {
      try {
        const response = await fetch('http://localhost:8080/sector-image');
        const data = await response.json();

        // Используй data.log для получения строки base64
        const base64data = data.log;

        // Обновляем состояние, чтобы отобразить изображение
        setImageSrc(base64data);
      } catch (error) {
        console.error('Ошибка при загрузке изображения:', error);
      }
    };

    fetchImage();
  }, []);

  return (
    <div className="photo">
      {imageSrc && <img src={`data:image/png;base64,${imageSrc}`} alt="Изображение" />}
    </div>
  );
};
