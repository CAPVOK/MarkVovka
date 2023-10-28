package api

import (
	"MarkVovka/backend/serviceStation/internal/app/config"
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"MarkVovka/backend/serviceStation/internal/app/simulation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cfg             *config.Config
	LastLocationData string
	LocationData    *ds.Location
	StationData     *simulation.Station // Добавлено поле StationData
}

func (h *Handler) StartSimulation() {
	simulation.StartSimulation()
}
func NewHandler(cfg *config.Config) *Handler {
	locationData := &ds.Location{
		Latitude:                  50.123,
		Longitude:                 30.456,
		Speed:                     200,
		Altitude:                  300,
		PlanetRadius:              6371,
		Angle:                     45,
		PlanetName:                "Earth",
		SolarPanelStatus:          true,
		FuelLevel:                 75.5,
		HullStatus:                "normal",
		Temperature:               25.5,
		ScientificInstrumentsStatus: "active",
		NavigationSystemStatus:     "enabled",
	}

    stationData := &simulation.Station{
        Latitude:  50.123,
        Longitude: 30.456,
    }

    return &Handler{
        Cfg:          cfg,
        LocationData: locationData,
        StationData:  stationData,
    }
}

func (h *Handler) UpdateStationData(c *gin.Context) {
	var requestData struct {
		Speed    float64 `json:"speed"`
		Altitude float64 `json:"altitude"`
		Angle    float64 `json:"angle"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Проверить непустые поля в requestData и обновить соответствующие поля в locationData
	if requestData.Speed != 0 {
		h.LocationData.Speed = requestData.Speed
	}
	if requestData.Altitude != 0 {
		h.LocationData.Altitude = requestData.Altitude
	}
	if requestData.Angle != 0 {
		h.LocationData.Angle = requestData.Angle
	}

	// Создать копию обновленных данных
	// Создать копию обновленных данных
	updatedData := &ds.Location{
		Latitude:                    h.StationData.Latitude,
		Longitude:                   h.StationData.Longitude,
		Speed:                       requestData.Speed,
		Altitude:                    requestData.Altitude,
		PlanetRadius:                h.LocationData.PlanetRadius,
		Angle:                       requestData.Angle,
		PlanetName:                  h.LocationData.PlanetName,
		SolarPanelStatus:            h.LocationData.SolarPanelStatus,  
		FuelLevel:                   h.LocationData.FuelLevel,  
		HullStatus:                  h.LocationData.HullStatus,  
		Temperature:                 h.LocationData.Temperature,  
		ScientificInstrumentsStatus: h.LocationData.ScientificInstrumentsStatus,  
		NavigationSystemStatus:      h.LocationData.NavigationSystemStatus,  
	}


	// Отправить обновленные данные в горутину для асинхронного обновления координат
	go func() {
		simulation.ParamsCh <- simulation.SimulationParams{
			Speed:  requestData.Speed,
			Angle:  requestData.Angle,
			Height: requestData.Altitude,
		}
	}()

	// Отправить обновленные данные в виде JSON
	c.JSON(http.StatusOK, gin.H{"message": "Station data updated successfully", "updatedData": updatedData})
}

func (h *Handler) ToggleSolarPanelsStatus(c *gin.Context) {
	// Парсим параметр activated из запроса
	activated, err := strconv.ParseBool(c.DefaultQuery("solarPanelStatus", "false"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for activated parameter"})
		return
	}

	// Обновить статус солнечных панелей в locationData
	h.LocationData.SolarPanelStatus = activated

	// Отправить статус солнечных панелей в функцию симуляции
	go func(activated bool) {
		simulation.ParamsCh <- simulation.SimulationParams{
			Speed:           h.LocationData.Speed,
			Angle:           h.LocationData.Angle,
			Height:          h.LocationData.Altitude,
			SolarPanelStatus: activated,
		}
	}(activated)

	c.JSON(http.StatusOK, gin.H{"message": "Solar panels status updated successfully", "data": h.LocationData})
}

func (h *Handler) ToggleScientificInstrumentsStatus(c *gin.Context) {
	// Парсим параметр scientificInstrumentsStatus из запроса
	status := c.DefaultQuery("scientificInstrumentsStatus", "inactive")

	// Проверяем допустимые значения
	switch status {
	case "active", "inactive", "maintenance":
		// Обновить статус научных инструментов в locationData
		h.LocationData.ScientificInstrumentsStatus = status

		// Отправить статус научных инструментов в функцию симуляции
		go func(status string) {
			simulation.ParamsCh <- simulation.SimulationParams{
				Speed:                    h.LocationData.Speed,
				Angle:                    h.LocationData.Angle,
				Height:                   h.LocationData.Altitude,
				SolarPanelStatus:         h.LocationData.SolarPanelStatus,
				ScientificInstrumentsStatus: status,
			}
		}(status)

		c.JSON(http.StatusOK, gin.H{"message": "Scientific instruments status updated successfully", "data": h.LocationData})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for scientificInstrumentsStatus parameter"})
	}
}

func (h *Handler) ToggleNavigationSystemStatus(c *gin.Context) {
	// Парсим параметр navigationSystemStatus из запроса
	status := c.DefaultQuery("navigationSystemStatus", "disabled")

	// Проверяем допустимые значения
	switch status {
	case "enabled", "disabled", "calibrating":
		// Обновить статус системы навигации в locationData
		h.LocationData.NavigationSystemStatus = status

		// Отправить статус системы навигации в функцию симуляции
		go func(status string) {
			simulation.ParamsCh <- simulation.SimulationParams{
				Speed:                  h.LocationData.Speed,
				Angle:                  h.LocationData.Angle,
				Height:                 h.LocationData.Altitude,
				SolarPanelStatus:       h.LocationData.SolarPanelStatus,
				ScientificInstrumentsStatus: h.LocationData.ScientificInstrumentsStatus,
				NavigationSystemStatus: status,
			}
		}(status)

		c.JSON(http.StatusOK, gin.H{"message": "Navigation system status updated successfully", "data": h.LocationData})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for navigationSystemStatus parameter"})
	}
}






