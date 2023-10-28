package api

import (
	"MarkVovka/backend/internal/app/ds"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// loginReq представляет структуру для декодирования JSON-запроса аутентификации.
type loginReq struct {
	Username    string `json:"username"`    // Login - поле для имени пользователя.
	Password string `json:"password"` // Password - поле для пароля пользователя.
}

// loginResp представляет структуру для кодирования JSON-ответа аутентификации.
type loginResp struct {
	ExpiresIn   int    `json:"expires_in"`  // ExpiresIn - продолжительность действия токена в секундах.
	AccessToken string `json:"access_token"` // AccessToken - сгенерированный JWT токен.
	TokenType   string `json:"token_type"`   // TokenType - тип токена (Bearer).
}

// Login обрабатывает POST-запрос для аутентификации пользователей.
func (h *Handler) Login(c *gin.Context) {
	cfg := h.Cfg // Получаем конфигурацию из обработчика.

	var req loginReq

	// Попытка декодировать JSON из тела запроса и преобразовать его в структуру loginReq.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Поиск пользователя в базе данных по логину.
	user, err := h.Repo.FindUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid login credentials"})
		return
	}

	// Проверка пароля.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid login credentials"})
		return
	}

	// Если аутентификация прошла успешно, генерируем JWT токен.
	token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bitop-admin",
		},
		UserUUID: user.UserUUID, // Используем UUID пользователя из базы данных.
		Scopes:   []string{},    // Пользовательские области видимости (пустой массив в данном случае).
	})

	// Подписываем токен, используя секретный ключ.
	tokenString, err := token.SignedString([]byte(cfg.JWT.Token))
	if err != nil {
		// Если не удалось подписать токен, возвращаем ошибку 500 Internal Server Error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}
	// Устанавливаем токен в куку
	c.SetCookie("jwt", tokenString, int(cfg.JWT.ExpiresIn), "/", "localhost:5173", false, true)	
	
	// Возвращаем успешный ответ с токеном в формате JSON.
	c.JSON(http.StatusOK, loginResp{
		ExpiresIn:   int(cfg.JWT.ExpiresIn),
		AccessToken: tokenString,
		TokenType:   "Bearer",
	})
}


// GetUser обрабатывает GET-запрос для получения данных пользователя по его UUID.
func (h *Handler) GetUser(c *gin.Context) {
	// Извлекаем UUID пользователя из контекста запроса.
	userUUID, exists := c.Get("user_uuid")
	if !exists {
		// Если UUID пользователя отсутствует в контексте, возвращаем ошибку 400 Bad Request.
		c.JSON(http.StatusBadRequest, gin.H{"error": "User UUID not found in context"})
		return
	}

	// Преобразуем UUID в нужный формат (предполагается, что userUUID имеет тип uuid.UUID).
	convertedUUID, ok := userUUID.(uuid.UUID)
	if !ok {
		// Если произошла ошибка при преобразовании UUID, возвращаем ошибку 400 Bad Request.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User UUID format"})
		return
	}

	// Запрос пользователя из базы данных по его UUID.
	user, err := h.Repo.FindUserByUUID(convertedUUID)
	if err != nil {
		// Если произошла ошибка при запросе пользователя из базы данных, возвращаем ошибку 500 Internal Server Error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}

	// Возвращаем данные пользователя в формате JSON.
	c.JSON(http.StatusOK, user)
}



// registerReq представляет структуру для декодирования JSON-запроса регистрации.
type registerReq struct {
	FullName    string `json:"full_name"` // Fullname - поле для имени пользователя.
	Username    string `json:"username"` // Username - поле для имени пользователя.
	Password 	string `json:"password"` // Password - поле для пароля пользователя.
	Email    	string `json:"email"`    // Email - поле для адреса электронной почты пользователя.
	// Другие поля, которые вы хотите получить от клиента при регистрации.
}

// Register обрабатывает POST-запрос для регистрации пользователей.
func (h *Handler) Register(c *gin.Context) {
	var req registerReq

	// Попытка декодировать JSON из тела запроса и преобразовать его в структуру registerReq.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Проверка полей в запросе.
	if req.Username == "" || req.Password == "" || req.Email == "" || req.FullName==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Хеширование пароля с использованием bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}


	// Создание нового пользователя с использованием данных из запроса.
	newUser := ds.User{
		UserUUID: uuid.New(),
		FullName: req.FullName,
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		// Другие поля, которые вы хотите сохранить в базе данных.
	}

	erro := h.Repo.Register(&newUser)
    if erro != nil {
        fmt.Println("Failed to register user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }
	// Если пользователь успешно зарегистрирован, отправьте успешный ответ.
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}


func (h *Handler) Logout(c *gin.Context) {
	// получаем заголовок
	jwtStr := c.GetHeader("Authorization")
	if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
		c.AbortWithStatus(http.StatusBadRequest) // отдаем что нет доступа

		return // завершаем обработку
	}

	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	_, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.Cfg.JWT.Token), nil
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)

		return
	}

	// сохраняем в блеклист редиса
	err = h.Redis.WriteJWTToBlacklist(c.Request.Context(), jwtStr, h.Cfg.JWT.ExpiresIn)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	c.Status(http.StatusOK)
}