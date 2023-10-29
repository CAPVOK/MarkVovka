package api

import (
	"MarkVovka/backend/serviceStation/internal/app/config"
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"MarkVovka/backend/serviceStation/internal/app/simulation"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fogleman/gg"
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
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Проверить непустые поля в requestData и обновить соответствующие поля в locationData
	if requestData.Speed != 0 {
		h.LocationData.Speed = requestData.Speed
	}

	// Создать копию обновленных данных
	// Создать копию обновленных данных
	updatedData := &ds.Location{
		Latitude:                    h.StationData.Latitude,
		Longitude:                   h.StationData.Longitude,
		Speed:                       requestData.Speed,
		Altitude:                    h.LocationData.Altitude,
		PlanetRadius:                h.LocationData.PlanetRadius,
		Angle:                       h.LocationData.Angle,
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

func (h *Handler) GetSectorImageByLongitude(c *gin.Context) {
	// Определить сектор на основе долготы из LocationData
	sectorCount := 20
	sector := int((h.LocationData.Longitude + 180.0) / 360.0 * float64(sectorCount))

	// Проверить, что сектор находится в допустимых пределах (0-19)
	if sector < 0 || sector >= sectorCount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sector"})
		return
	}

	// Создать новое изображение
	const imageSize = 200 // Размер изображения (ширина и высота)
	dc := gg.NewContext(imageSize, imageSize)

	// Рисовать изображение сектора (просто для примера)
	dc.SetRGB(0, float64(sector)/float64(sectorCount), 0)
	dc.DrawRectangle(50, 50, 100, 100)
	dc.Fill()

	// Сохранить изображение в формате PNG
	imagePath := fmt.Sprintf("/sector%d.png", sector) // Укажите путь, куда сохранить изображение
	if err := dc.SavePNG(imagePath); err != nil {
		return 
	}


	// Создать буфер для сохранения изображения
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	// Кодировать изображение в Base64
	encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Отправить изображение в JSON ответе
	c.JSON(http.StatusOK, gin.H{"image": encodedImage})
}








