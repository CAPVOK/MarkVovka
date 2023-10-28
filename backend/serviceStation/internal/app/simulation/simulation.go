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
		s.Latitude = 90
	}

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
		log.Println("dfgdfgdfg")
		// Обработка ошибки записи в файл
	}
}

func StartSimulation(location *ds.Location, done chan struct{}) {
    var speedBased float64 = 111
    var angle float64 = 60
    var height float64 = 0
	log.Println("sdfgsfdg")
    station := &Station{
        Latitude:  location.Latitude,
        Longitude: location.Longitude,
    }

    for {
        select {
        case <-done:
            // Если получен сигнал о завершении, выходим из цикла
            return
        default:
            // В противном случае обновляем координаты и ждем 1 секунду
            station.UpdateCoordinates(speedBased, angle, height)
			log.Println("sdgfsdgfsdg")
            time.Sleep(1 * time.Second)
        }
    }
}
