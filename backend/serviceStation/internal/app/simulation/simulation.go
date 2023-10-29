package simulation

import (
	"log"
	"math"
	"time"

	"MarkVovka/backend/serviceStation/internal/app/ds"
)

const R = 6371
const update_time = 1
const updateInterval = 1 * time.Second // Интервал времени для обновления координат

type Station struct {
	Latitude  float64
	Longitude float64
}
type SimulationParams struct {
	Speed  						float64
	Height  					float64
	SolarPanelStatus 			bool
	ScientificInstrumentsStatus string
	NavigationSystemStatus      string
}
var ParamsCh = make(chan SimulationParams) // Канал для передачи параметров симуляции


var lastLocation *ds.Location

func StartSimulation() {
	lastLocation = &ds.Location{
		Latitude:                    50.123,
		Longitude:                   30.456,
		Speed:                       7.685,
		Altitude:                    300,
		PlanetRadius:                6371,
		Angle:                       0,
		PlanetName:                  "Earth",
		SolarPanelStatus:            true,
		FuelLevel:                   75.5,
		HullStatus:                  "normal",
		Temperature:                 25.5,
		ScientificInstrumentsStatus: "active",
		NavigationSystemStatus:      "enabled",
	}
	
	go func() {
		for {
			select {
			case params := <-ParamsCh:
				// Обновляем координаты симуляции на основе полученных параметров
				updateCoordinates(params.Speed, params.SolarPanelStatus, params.ScientificInstrumentsStatus, params.NavigationSystemStatus)
			default:
				// Если нет новых параметров, продолжаем обновлять координаты с последними известными параметрами
				updateCoordinates(lastLocation.Speed, lastLocation.SolarPanelStatus,lastLocation.ScientificInstrumentsStatus, lastLocation.NavigationSystemStatus)
			}

			// Ждем заданный интервал времени перед следующим обновлением
			//time.Sleep(updateInterval)
		}
	}()
}


// updateCoordinates обновляет координаты станции на основе переданных параметров.
func updateCoordinates(speed float64, solar_panel_status bool,scientific_instruments_status, navigation_system_status string) {

	var iterations = 10

	// Рассчитываем разницу между текущей скоростью и желаемой скоростью
	speedDifference := speed - lastLocation.Speed

	// Вычисляем шаг интерполяции для скорости
	speedStep := speedDifference / float64(iterations)

	// Интерполируем скорость по заданному количеству итераций
	for i := 0; i < iterations; i++ {
		// Обновляем координаты на основе текущей скорости
		speedBased := lastLocation.Speed + speedStep
		angleRad := lastLocation.Angle * math.Pi / 180
		newHeight := ((((9.81 * R * R) / (speedBased * speedBased)) - R*1000) / 1000)
		// Обновляем широту и долготу на основе скорости и угла
		newLatitude := lastLocation.Latitude + speedBased*math.Cos(angleRad)*updateInterval.Seconds()*360/((newHeight+R)*2*math.Pi)
		newLongitude := lastLocation.Longitude + speedBased*math.Sin(angleRad)*updateInterval.Seconds()*360/((newHeight+R)*2*math.Pi)
	
		// Проверяем и корректируем координаты, чтобы они оставались в пределах [-90, 90] для широты и [-180, 180] для долготы
		if newLongitude < -180 {
			newLongitude += 360
		} else if newLongitude > 180 {
			newLongitude -= 360
		}
	
		if newLatitude < -90 {
			newLatitude += -90
		} else if newLatitude > 90 {
			newLatitude -= 90
		}

		// Обновляем последние координаты
		lastLocation.Latitude = newLatitude
		lastLocation.Longitude = newLongitude
		lastLocation.Speed = speedBased
		lastLocation.Altitude = newHeight
		lastLocation.SolarPanelStatus = solar_panel_status
		lastLocation.ScientificInstrumentsStatus = scientific_instruments_status
		lastLocation.NavigationSystemStatus = navigation_system_status

		// Записываем новые координаты в файл JSON
		err := ds.WriteLocationToFile(lastLocation)
		if err != nil {
			log.Println("Error writing location data to file:", err)
		}

		// Ждем заданный интервал времени перед следующим обновлением
		time.Sleep(updateInterval)
	}
}



