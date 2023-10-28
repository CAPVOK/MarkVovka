package ds

import "github.com/google/uuid"

type User struct {
	UserUUID uuid.UUID `gorm:"autoIncrement;primarykey" json:"user_id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Username string    `json:"username"` // Поле для логина (или "логин")
	Password string    `json:"password"`
}
