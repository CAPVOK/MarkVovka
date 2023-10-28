package api

import (
	"MarkVovka/backend/internal/app/ds"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

const jwtPrefix = "Bearer "

// WithAuthCheck - Middleware функция для проверки авторизации с использованием JWT токена.
func (h *Handler) WithAuthCheck(c *gin.Context) {																							
	// Получаем JWT токен из заголовка Authorization.
	jwtStr := c.GetHeader("Authorization")
	if !strings.HasPrefix(jwtStr, jwtPrefix) {
		// Если токен не содержит префикс, возвращаем ошибку 403 Forbidden и завершаем обработку.
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// Отрезаем префикс "Bearer " от JWT токена.
	jwtStr = jwtStr[len(jwtPrefix):]

	// проверяем jwt в блеклист редиса
	err := h.Redis.CheckJWTInBlacklist(c.Request.Context(), jwtStr)
	if err == nil { // значит что токен в блеклисте
	   c.AbortWithStatus(http.StatusForbidden)

	   return
	}
	if !errors.Is(err, redis.Nil) { // значит что это не ошибка отсуствия - внутренняя ошибка
		c.AbortWithError(http.StatusInternalServerError, err)

	   return
	}


	// Парсим JWT токен с использованием заданных claims (параметров) и проверяем его подпись.
	claims, err := parseJWTToken(jwtStr, h.Cfg.JWT.Token)
	if err != nil {
		// Если произошла ошибка при проверке токена, возвращаем ошибку 403 Forbidden и логируем ошибку.
		c.AbortWithStatus(http.StatusForbidden)
		log.Println(err)
		return
	}
	// Извлекаем UUID пользователя из claims.
	userUUID := claims.UserUUID

	// Передаем данные пользователя в контекст, чтобы они были доступны в следующих обработчиках.
	c.Set("user_uuid", userUUID)

	// Продолжаем выполнение цепочки запросов.
	c.Next()
}
func parseJWTToken(tokenString, secretKey string) (*ds.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // Используем секретный ключ для проверки подписи токена.
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ds.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT token")
	}

	return claims, nil
}
