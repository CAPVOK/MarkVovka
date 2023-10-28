package main

import (
	"MarkVovka/backend/internal/app/ds"
	"MarkVovka/backend/internal/app/dsn"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    // Путь к файлу .env относительно main.go
	envFilePath := "../../.env"
    _ = godotenv.Load(envFilePath)
    db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Явно мигрировать только нужные таблицы
    err = db.AutoMigrate(&ds.User{})
    if err != nil {
        panic("cant migrate db")
    }
}
