package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ExecuteConsoleCommand(c *gin.Context) {
	// Получаем параметры из запроса
	command := c.Param("command")
	value := c.DefaultQuery("value", "false")

	// Определяем URL для запроса на основе команды
	var url string
	switch command {
	case "solar-panel-status":
		url = "http://localhost:8081/solar-panel-status?solarPanelStatus=" + value
	case "scientific-instruments-status":
		url = "http://localhost:8081/scientific-instruments-status?scientificInstrumentsStatus=" + value
	case "navigation-system-status":
		url = "http://localhost:8081/navigation-system-status?navigationSystemStatus=" + value
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid command"})
		return
	}

	
		// Создаем PUT запрос
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // Устанавливаем тип контента для запроса

	// Отправляем PUT запрос
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer response.Body.Close()

	// Читаем ответ от сервера
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Возвращаем ответ в виде JSON
	c.JSON(http.StatusOK, gin.H{"message": "Command executed successfully", "response": string(responseData)})
}

