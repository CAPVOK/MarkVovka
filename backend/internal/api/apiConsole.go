package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)
func isValidCommand(command string) bool {
	// Здесь можно проверить, существует ли команда в вашем списке допустимых команд
	// Вернуть true, если команда допустима, и false в противном случае
	return command == "solar-panel-status" || command == "scientific-instruments-status" || command == "navigation-system-status" || command == "help" || command == "speed"
}

func (h *Handler) ExecuteConsoleCommand(c *gin.Context) {
	// Читаем текстовую команду из тела запроса
	var commandMessage struct {
		Message string `json:"message"`
	}
	// Логирование с текущим временем
	logMessage := fmt.Sprintf("[%s] Command:", time.Now().Format(time.RFC3339))
	
	if err := c.ShouldBindJSON(&commandMessage); err != nil {
		c.JSON(http.StatusOK, gin.H{"masg": "Invalid JSON data",
											"log": logMessage})
		return
	}

		// Разделяем команду и значение
	parts := strings.Fields(commandMessage.Message)

	// Проверяем, что ввод содержит только одну часть (команду)
	if len(parts) != 1 && len(parts) != 2 {
		c.JSON(http.StatusOK, gin.H{"msg": "Invalid command format",
											"log":logMessage})
		return
	}

	command := parts[0]
	var value string
	if len(parts) == 2 {
		value = parts[1]
	}

	// Проверяем, существует ли команда
	if !isValidCommand(command) {
		c.JSON(http.StatusOK, gin.H{"msg": "Invalid command. Type 'help' for available commands.",
											"log":logMessage})
		return
	}

	
	// Если команда - "help", отправляем сообщение со списком доступных команд
	if command == "help" {
		helpMessage := "Available commands:\n" +
			"- solar-panel-status <true/false>: Toggle solar panel status\n" +
			"- scientific-instruments-status <active/inactive/maintenance>: Set scientific instruments status\n" +
			"- navigation-system-status <enabled/disabled/calibrating>: Set navigation system status\n" +
			"- help: Show available commands"
		c.JSON(http.StatusOK, gin.H{"msg": helpMessage,
									"log":logMessage})
		return
	}

	// Определяем URL для запроса на основе команды
	var url string
	switch command {
	case "solar-panel-status":
		url = "http://localhost:8081/solar-panel-status?solarPanelStatus=" + value
	case "scientific-instruments-status":
		url = "http://localhost:8081/scientific-instruments-status?scientificInstrumentsStatus=" + value
	case "navigation-system-status":
		url = "http://localhost:8081/navigation-system-status?navigationSystemStatus=" + value
	case "speed":
		url = "http://localhost:8081/update-station-speed?speed=" + value
	}

	// Создаем PUT запрос
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to create request",
		"log":logMessage})
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // Устанавливаем тип контента для запроса

	// Отправляем PUT запрос
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Failed to send request","log":logMessage})
		return
	}
	defer response.Body.Close()

	// Возвращаем ответ в виде JSON
	c.JSON(http.StatusOK, gin.H{"msg": "Command executed successfully", "log":logMessage})
}


