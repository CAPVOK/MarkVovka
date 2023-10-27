package api

import (
	"MarkVovka/backend/internal/app/config"
	"MarkVovka/backend/internal/app/redis"
	"MarkVovka/backend/internal/app/repository"
)

type Handler struct {
	Repo *repository.Repository
	Cfg *config.Config
	Redis *redis.Client
}

func NewHandler(repo *repository.Repository, cfg *config.Config, redis *redis.Client) *Handler {
	return &Handler{Repo: repo, Cfg:cfg, Redis:redis}
}




