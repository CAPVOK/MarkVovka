package app

import (
	"MarkVovka/backend/serviceStation/internal/api"
	"MarkVovka/backend/serviceStation/internal/app/config"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Application представляет основное приложение.
type Application struct {
    Config       *config.Config
}

// New создает новый объект Application и настраивает его.
func New(ctx context.Context) (*Application, error) {
    // Инициализируйте конфигурацию
    cfg, err := config.NewConfig(ctx)
    if err != nil {
        return nil, err
    }
    // Инициализируйте и настройте объект Application
    app := &Application{
        Config: cfg,
        // Установите другие параметры вашего приложения, если необходимо
    }

    return app, nil
}

// SpaceStationService описывает методы, которые может вызывать бэкенд.
type SpaceStationService struct{}


// Run запускает приложение.
func (app *Application) Run() {
    handler := api.NewHandler(app.Config)
    handler.StartSimulation()
    r := gin.Default()
    r.GET("/location", handler.Location)


    r.PUT("/update-station-speed", handler.UpdateSpeedStation)
    r.PUT("/solar-panel-status", handler.ToggleSolarPanelsStatus)
    r.PUT("scientific-instruments-status", handler.ToggleScientificInstrumentsStatus)
    r.PUT("navigation-system-status", handler.ToggleNavigationSystemStatus)
    
    r.GET("/sector-image", handler.GetSectorImageByLongitude)

    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
    

}