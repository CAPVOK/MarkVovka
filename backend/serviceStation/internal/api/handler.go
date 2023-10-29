package api

import (
	"MarkVovka/backend/serviceStation/internal/app/config"
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"MarkVovka/backend/serviceStation/internal/app/simulation"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
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

func (h *Handler) UpdateSpeedStation(c *gin.Context) {
	// Получить значение параметра "speed" из query parameters
	speedStr := c.Query("speed")

	// Преобразовать значение параметра в float64
	speed, err := strconv.ParseFloat(speedStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid speed parameter"})
		return
	}

	// Проверить непустые поля в requestData и обновить соответствующие поля в locationData
	if speed != 0 {
		h.LocationData.Speed = speed
	}

	// Создать копию обновленных данных
	updatedData := &ds.Location{
		Latitude:                    h.StationData.Latitude,
		Longitude:                   h.StationData.Longitude,
		Speed:                       speed,
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
			Speed: speed,
		}
	}()

	// Отправить обновленные данные в виде JSON
	c.JSON(http.StatusOK, gin.H{"msg": "Station data updated successfully", "log": updatedData})
}


func (h *Handler) ToggleSolarPanelsStatus(c *gin.Context) {
	// Парсим параметр activated из запроса
	activated, err := strconv.ParseBool(c.DefaultQuery("solarPanelStatus", "false"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid value for activated parameter"})
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

	c.JSON(http.StatusOK, gin.H{"msg": "Solar panels status updated successfully", "log": h.LocationData})
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

		c.JSON(http.StatusOK, gin.H{"msg": "Scientific instruments status updated successfully", "log": h.LocationData})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid value for scientificInstrumentsStatus parameter"})
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

	// Путь к папке, содержащей изображения секторов
	imageFolderPath := "../resources/data"

	// Формирование пути к изображению сектора
	imageFileName := fmt.Sprintf("%d.png", sector)
	imageFilePath := filepath.Join(imageFolderPath, imageFileName)

	// Чтение изображения из файла
	imageData, err := ioutil.ReadFile(imageFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	// Кодирование изображения в Base64
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Отправить изображение в JSON ответе
	c.JSON(http.StatusOK, gin.H{"log": encodedImage, "msg":"OKEY"})
}









