package repository

import (
	"MarkVovka/backend/internal/app/ds"

	"github.com/google/uuid"
)

func (r *Repository) Register(user *ds.User) error {
	if user.UserUUID == uuid.Nil {
		user.UserUUID = uuid.New()
	}
	// Создаем багаж
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByUsername ищет пользователя в базе данных по логину.
func (r *Repository) FindUserByUsername(username string) (*ds.User, error) {
	var user ds.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByUUID ищет пользователя в базе данных по его UUID.
func (r *Repository) FindUserByUUID(userUUID uuid.UUID) (*ds.User, error) {
	user := &ds.User{}
	if err := r.db.Where("user_uuid = ?", userUUID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
