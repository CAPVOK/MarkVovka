package app

import (
	"context"
	"fmt"
	"log"

	"MarkVovka/backend/internal/api"
	"MarkVovka/backend/internal/app/config"
	"MarkVovka/backend/internal/app/dsn"
	"MarkVovka/backend/internal/app/redis"
	"MarkVovka/backend/internal/app/repository"

	"github.com/gin-gonic/gin"
)

// Application представляет основное приложение.
type Application struct {
    Config       *config.Config
    Repository   *repository.Repository
    Redis        *redis.Client
    RequestLimit int
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context) (*Application, error) {
    // Инициализируйте конфигурацию
    cfg, err := config.NewConfig(ctx)
    if err != nil {
        return nil, err
    }

    // Инициализируйте подключение к базе данных (DB)
    repo, err := repository.New(dsn.FromEnv())
    if err != nil {
        return nil, err
    }

    redisClient, err := redis.New(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

    // Инициализируйте и настройте объект Application
    app := &Application{
        Config: cfg,
        Repository: repo,
        Redis: redisClient,
        // Установите другие параметры вашего приложения, если необходимо
    }

    return app, nil
}

// Run запускает приложение.
func (app *Application) Run() {
    handler := api.NewHandler(app.Repository,app.Config, app.Redis)
    r := gin.Default()
    // Разрешить запросы от любого источника
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
   // Создаем группу роутов, к которой хотим применить мидлваре
    authGroup := r.Group("/auth")
    {
        // POST /auth/signup
        // Регистрирует нового пользователя.
        // Принимает JSON с полями full_name, username, password, email.
        // Пример запроса:
        // {
        //     "full_name": "John Doe",
        //     "username": "johndoe",
        //     "password": "password123",
        //     "email": "johndoe@example.com"
        // }
        authGroup.POST("/signup", handler.Register)

        // POST /auth/login
        // Аутентифицирует пользователя.
        // Принимает JSON с полями username и password.
        // Пример запроса:
        // {
        //     "username": "johndoe",
        //     "password": "password123"
        // }
        // Возвращает JWT токен в ответе в формате JSON.
        authGroup.POST("/login", handler.Login)

        // POST /auth/logout
        // Завершает сеанс пользователя и добавляет токен в черный список.
        // Требует заголовка Authorization с токеном в формате Bearer.
        // Пример запроса:
        // {
        //     "Authorization": "Bearer <ваш JWT токен>"
        // }
        authGroup.POST("/logout", handler.Logout)
        //authGroup.POST("/logout", handlerLogoutHandler)
    }

    mainGroup := r.Group("/")
    {
        // Применяем мидлваре к этой группе
        mainGroup.Use(handler.WithAuthCheck)

        // GET /user
        // Возвращает данные пользователя, чей токен был передан в заголовке Authorization.
        // Требует действительный токен в заголовке Authorization в формате Bearer.
        // Пример запроса:
        // {
        //     "Authorization": "Bearer <ваш JWT токен>"
        // }
        mainGroup.GET("/user", handler.GetUser)
    }

    r.GET("/data", handler.GetStationData)
    r.POST("/update",handler.UpdateStationData)
    r.GET("/sector-image",handler.GetSectorImage)
    r.POST("/console",handler.ExecuteConsoleCommand)
    //r.PUT("/console/:command")
    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
}