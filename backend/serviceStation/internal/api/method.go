package api

import (
	"MarkVovka/backend/serviceStation/internal/app/ds"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) Location(c *gin.Context) {
	// Читаем данные из JSON файла
	data, err := ds.ReadLocationFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Отправляем данные в формате JSON
	c.JSON(http.StatusOK, data)
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

	// Сохранить обновленные данные в файл JSON
	err := ds.WriteLocationToFile(h.LocationData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data to JSON file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Station data updated and saved successfully"})
}
