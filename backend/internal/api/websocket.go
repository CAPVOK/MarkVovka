package api

import (
	"io/ioutil"
	"log"
	"net/http"
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