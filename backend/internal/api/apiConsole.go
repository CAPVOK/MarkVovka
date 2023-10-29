package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommandRequest представляет структуру для JSON-запроса с командой.
type CommandRequest struct {
	Command string `json:"command"`
	Value   string `json:"value"`
}

func (h *Handler) ExecuteConsoleCommand(c *gin.Context) {
	// Читаем JSON-запрос из тела запроса
	var commandRequest CommandRequest
	if err := c.ShouldBindJSON(&commandRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Определяем URL для запроса на основе команды
	var url string
	switch commandRequest.Command {
	case "solar-panel-status":
		url = "http://localhost:8081/solar-panel-status?solarPanelStatus=" + commandRequest.Value
	case "scientific-instruments-status":
		url = "http://localhost:8081/scientific-instruments-status?scientificInstrumentsStatus=" + commandRequest.Value
	case "navigation-system-status":
		url = "http://localhost:8081/navigation-system-status?navigationSystemStatus=" + commandRequest.Value
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
