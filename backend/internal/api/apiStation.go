package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	// Создаем JSON объект для отправки
	requestBody, err := json.Marshal(map[string]float64{
		"speed":    requestData.Speed,
		"altitude": requestData.Altitude,
		"angle":    requestData.Angle,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JSON request"})
		return
	}

	// Отправить POST-запрос на другой сервер (порт 8081) с данными из запроса
	response, err := http.Post("http://localhost:8081/update", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to port 8081"})
		return
	}
	defer response.Body.Close()

	// Прочитать ответ от другого сервера
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from port 8081"})
		return
	}

	// Отправить полученные данные обратно клиенту
	c.Data(http.StatusOK, "application/json", responseData)
}

func (h *Handler) GetSectorImage(c *gin.Context) {
	// Отправить GET-запрос на другой сервер (порт 8081)
	response, err := http.Get("http://localhost:8081/sector-image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to send GET request to port 8081"})
		return
	}
	defer response.Body.Close()

	// Прочитать ответ от другого сервера
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to read response from port 8081"})
		return
	}

	// Отправить полученные данные обратно клиенту в виде JSON
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(responseData, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}







