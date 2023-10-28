package api

import (
	"MarkVovka/backend/serviceStation/internal/app/config"
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"MarkVovka/backend/serviceStation/internal/app/simulation"
)

type Handler struct {
	Cfg             *config.Config
	LastLocationData string
	LocationData    *ds.Location
	StationData     *simulation.Station // Добавлено поле StationData
}

func (h *Handler) StartSimulation() {
    done := make(chan struct{})
    defer close(done)

    go simulation.StartSimulation(done)
}

func NewHandler(cfg *config.Config) *Handler {
	// Реальная логика получения данных о местоположении станции здесь.
	locationData := &ds.Location{
		Latitude:     50.123,
		Longitude:    30.456,
		Speed:        200,
		Altitude:     300,
		PlanetRadius: 6371,
		Angle:        45,
		PlanetName:   "Earth",
		Status:       "active",
	}

	stationData := &simulation.Station{Latitude: locationData.Latitude, Longitude: locationData.Longitude}

	return &Handler{
		Cfg:          cfg,
		LocationData: locationData,
		StationData:  stationData,
	}
}
