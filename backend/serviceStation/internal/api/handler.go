package api

import (
	"MarkVovka/backend/serviceStation/internal/app/config"
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"MarkVovka/backend/serviceStation/internal/app/simulation"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
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
		Latitude:                    50.123,
		Longitude:                   0,
		Speed:                       7.685,
		Altitude:                    300,
		PlanetRadius:                6371,
		Angle:                       0,
		PlanetName:                  "Земля",
		SolarPanelStatus:            "true",
		FuelLevel:                   75.5,
		HullStatus:                  "нормально",
		Temperature:                 25.5,
		ScientificInstrumentsStatus: "активен",
		NavigationSystemStatus:      "включена",
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
	// Парсим параметр solarPanelStatus из запроса
	status := c.DefaultQuery("solarPanelStatus", "off")
	var solarPanelStatus string

	// Проверяем допустимые значения
	switch status {
	case "on":
		solarPanelStatus = "открыто"
	case "off":
		solarPanelStatus = "закрыто"
	default:
		c.JSON(http.StatusOK, gin.H{"message": "Invalid value for solarPanelStatus parameter"})
		return
	}

	// Обновить статус солнечных панелей в locationData
	h.LocationData.SolarPanelStatus = solarPanelStatus

	// Отправить статус солнечных панелей в функцию симуляции
	go func(status string) {
		simulation.ParamsCh <- simulation.SimulationParams{
			Speed:           h.LocationData.Speed,
			SolarPanelStatus: status,
		}
	}(solarPanelStatus)

	c.JSON(http.StatusOK, gin.H{"message": "Solar panels status updated successfully", "data": h.LocationData})
}


func (h *Handler) ToggleScientificInstrumentsStatus(c *gin.Context) {
	// Парсим параметр scientificInstrumentsStatus из запроса
	status := c.DefaultQuery("scientificInstrumentsStatus", "off")
	var instrumentsStatus string

	// Проверяем допустимые значения
	switch status {
	case "active":
		instrumentsStatus = "активен"
	case "inactive":
		instrumentsStatus = "не активный"
	case "maintenance":
		instrumentsStatus = "требуют обслуживания"
	default:
		c.JSON(http.StatusOK, gin.H{"message": "Invalid value for scientificInstrumentsStatus parameter"})
		return
	}

	// Обновить статус научных инструментов в locationData
	h.LocationData.ScientificInstrumentsStatus = instrumentsStatus

	// Отправить статус научных инструментов в функцию симуляции
	go func(status string) {
		simulation.ParamsCh <- simulation.SimulationParams{
			Speed:                      h.LocationData.Speed,
			SolarPanelStatus:           h.LocationData.SolarPanelStatus,
			ScientificInstrumentsStatus: status,
		}
	}(instrumentsStatus)

	c.JSON(http.StatusOK, gin.H{"message": "Scientific instruments status updated successfully", "data": h.LocationData})
}

func (h *Handler) ToggleNavigationSystemStatus(c *gin.Context) {
	// Парсим параметр navigationSystemStatus из запроса
	status := c.DefaultQuery("navigationSystemStatus", "disabled")

	// Проверяем допустимые значения
	var systemStatus string
	switch status {
	case "enabled":
		systemStatus = "включена"
	case "disabled":
		systemStatus = "выключена"
	case "calibrating":
		systemStatus = "калибруется"
	default:
		c.JSON(http.StatusOK, gin.H{"message": "Invalid value for navigationSystemStatus parameter"})
		return
	}

	// Обновить статус системы навигации в locationData
	h.LocationData.NavigationSystemStatus = systemStatus

	// Отправить статус системы навигации в функцию симуляции
	go func(status string) {
		simulation.ParamsCh <- simulation.SimulationParams{
			Speed:                      h.LocationData.Speed,
			SolarPanelStatus:           h.LocationData.SolarPanelStatus,
			ScientificInstrumentsStatus: h.LocationData.ScientificInstrumentsStatus,
			NavigationSystemStatus:     status,
		}
	}(systemStatus)

	c.JSON(http.StatusOK, gin.H{"message": "Navigation system status updated successfully", "data": h.LocationData})
}

func (h *Handler) GetSectorImageByLongitude(c *gin.Context) {
	// Определить сектор на основе долготы из LocationData
	sectorCount := 20
	sector := int((h.LocationData.Latitude + 180.0) / 360.0 * float64(sectorCount))

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

	// Декодирование изображения
	img, err := imaging.Decode(bytes.NewReader(imageData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode image"})
		return
	}

	// Определение коэффициента масштабирования на основе значения Longitude
	scaleFactor := 1.0 + math.Abs(h.LocationData.Altitude/1000.0) // Примерный расчет коэффициента масштабирования

	// Масштабирование изображения
	scaledImg := imaging.Resize(img, int(float64(img.Bounds().Dx())*scaleFactor), int(float64(img.Bounds().Dy())*scaleFactor), imaging.Lanczos)

	// Кодирование масштабированного изображения в Base64
	var encodedImage string
	if err := func() error {
		buffer := new(bytes.Buffer)
		if err := imaging.Encode(buffer, scaledImg, imaging.PNG); err != nil {
			return err
		}
		encodedImage = base64.StdEncoding.EncodeToString(buffer.Bytes())
		return nil
	}(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode image"})
		return
	}

	// Отправить изображение в JSON ответе
	c.JSON(http.StatusOK, gin.H{"log": encodedImage, "msg": "OKEY"})
}















