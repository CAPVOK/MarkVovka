package simulation

import (
	"log"
	"math"
	"time"

	"MarkVovka/backend/serviceStation/internal/app/ds"
)

const R = 6371
const update_time = 1

type Station struct {
	Latitude  float64
	Longitude float64
}
type simulationParams struct {
	Speed  float64
	Angle float64
	Height  float64
}



var lastLocation *ds.Location

func StartSimulation(done chan struct{}) {
	// Инициализация последних координат из файла JSON
	location, err := ds.ReadLocationFromFile()
	if err != nil {
		log.Println("Error reading location data from file:", err)
		return
	}
	lastLocation = location

	for {
		select {
		default:
			// Обновляем координаты симуляции с использованием последних значений из файла JSON
			station := &Station{
				Latitude:  lastLocation.Latitude,
				Longitude: lastLocation.Longitude,
			}

			// В противном случае обновляем координаты и ждем 1 секунду
			params := &simulationParams{
				Speed:  lastLocation.Speed,
				Angle:  lastLocation.Angle,
				Height: lastLocation.Altitude,
			}

			station.UpdateCoordinates(params.Speed, params.Angle, params.Height)
			log.Println("sdgfsdgfsdg")
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *Station) UpdateCoordinates(speedBased, angle, height float64) {
	angleRad := angle * math.Pi / 180
	s.Latitude = s.Latitude + speedBased*math.Cos(angleRad)*update_time*360/((height+R)*2*math.Pi)
	s.Longitude = s.Longitude + speedBased*math.Sin(angleRad)*update_time*360/((height+R)*2*math.Pi)

	if s.Longitude < -180 {
		s.Longitude += 360
	} else if s.Longitude > 180 {
		s.Longitude -= 360
	}

	if s.Latitude < -90 {
		s.Latitude = -90
	} else if s.Latitude > 90 {
		s.Latitude -= 180
	}

	// Обновляем последние координаты
	lastLocation.Latitude = s.Latitude
	lastLocation.Longitude = s.Longitude

	// Создайте новый объект Location с обновленными координатами и запишите его в файл
	newLocation := &ds.Location{
		Latitude:     s.Latitude,
		Longitude:    s.Longitude,
		Speed:        speedBased,  // Укажите желаемое значение для скорости
		Altitude:     height,  // Укажите желаемое значение для высоты
		PlanetRadius: 6371, // Укажите желаемое значение для радиуса планеты
		Angle:        angle,   // Укажите желаемый угол
		PlanetName:   "Earth", // Укажите желаемое имя планеты
		Status:       "active", // Укажите желаемый статус
	}

	// Запишите новые координаты в файл JSON
	err := ds.WriteLocationToFile(newLocation)
	if err != nil {
		log.Println("Error writing location data to file:", err)
	}
}
