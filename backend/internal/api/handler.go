package api

import (
	"MarkVovka/backend/internal/app/config"
	"MarkVovka/backend/internal/app/redis"
	"MarkVovka/backend/internal/app/repository"
	"sync"

	"github.com/gorilla/websocket"
)

type Handler struct {
	Repo *repository.Repository
	Cfg *config.Config
	Redis *redis.Client
}

type WebSocketHandler struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	mutex     sync.Mutex
}

func NewHandler(repo *repository.Repository, cfg *config.Config, redis *redis.Client) *Handler {
	return &Handler{Repo: repo, Cfg:cfg, Redis:redis}
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		mutex:     sync.Mutex{},
	}
}

type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:   ws,
		send: make(chan []byte),
	}
}



