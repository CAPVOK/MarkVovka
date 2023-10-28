package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost string
	ServicePort int
	JWT         JWTConfig // Новое поле для конфигурации JWT
	Redis 		RedisConfig
}

type JWTConfig struct {
	SigningMethod jwt.SigningMethod
	Token         string
	ExpiresIn     time.Duration
}

type RedisConfig struct {
	Host        string
	Password    string
	Port        int
	User        string
	DialTimeout time.Duration
	ReadTimeout time.Duration
}
const (
	envRedisHost = "REDIS_HOST"
	envRedisPort = "REDIS_PORT"
	envRedisUser = "REDIS_USER"
	envRedisPass = "REDIS_PASSWORD"
 )

func NewConfig(ctx context.Context) (*Config, error) {
	var err error
	envFilePath := "../.env"
	configName := "config"
	_ = godotenv.Load(envFilePath)
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	// Установка значений напрямую
	cfg.JWT.SigningMethod = jwt.GetSigningMethod("HS256")
	cfg.JWT.Token = "markgregr"
	expiresIn, _ := time.ParseDuration("1h")
	cfg.JWT.ExpiresIn = expiresIn

	cfg.Redis.Host = os.Getenv(envRedisHost)
	cfg.Redis.Port, err = strconv.Atoi(os.Getenv(envRedisPort))
	if err != nil {
	return nil, fmt.Errorf("redis port must be int value: %w", err)
	}
	cfg.Redis.Password = os.Getenv(envRedisPass)
	cfg.Redis.User = os.Getenv(envRedisUser)


	log.Info("config parsed")
	log.Println(cfg)
	return cfg, nil
}


