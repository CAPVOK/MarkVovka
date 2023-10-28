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




