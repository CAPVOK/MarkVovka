package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) GetStationData(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка при установке веб-сокет соединения:", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second) // Таймер с интервалом 5 секунд

	for {
		select {
		case <-c.Request.Context().Done():
			log.Println("Соединение закрыто по запросу клиента")
			return
		case <-ticker.C:
			// Отправить запрос на другой сервис (на примере GET запроса)
			resp, err := http.Get("http://localhost:8081/location")
			if err != nil {
				log.Println("Ошибка при выполнении HTTP запроса:", err)
				continue
			}
			defer resp.Body.Close()

			// Прочитать данные из ответа
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Ошибка при чтении данных из HTTP ответа:", err)
				continue
			}

			// Отправить полученные данные через веб-сокет
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Ошибка при отправке данных в веб-сокет:", err)
				return
			}
		}
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

	// Отправить POST-запрос на другой сервер (порт 8081) с данными из запроса
	response, err := http.Post("http://localhost:8081/update", "application/json", bytes.NewBuffer([]byte(`{
		"speed": ` + strconv.FormatFloat(requestData.Speed, 'f', -1, 64) + `,
		"altitude": ` + strconv.FormatFloat(requestData.Altitude, 'f', -1, 64) + `,
		"angle": ` + strconv.FormatFloat(requestData.Angle, 'f', -1, 64) + `
	}`)))
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
	c.JSON(http.StatusOK, gin.H{"message": "Station data updated successfully", "data": string(responseData)})
}

