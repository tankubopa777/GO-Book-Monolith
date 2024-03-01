package server

import (
	"fmt"

	"tansan/config"
	userHandlers "tansan/user/handlers"
	userRepositories "tansan/user/repositories"
	userUsecases "tansan/user/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
 app *echo.Echo
 db  *gorm.DB
 cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
 return &echoServer{
  app: echo.New(),
  db:  db,
  cfg: cfg,
 }
}

func (s *echoServer) Start() {
 s.initializeUserHttpHandler() // <- add this function here

 s.app.Use(middleware.Logger())

 serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
 s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeUserHttpHandler() {
 // Initialize all layers
 userPostgresRepository := userRepositories.NewUserPostgresRepository(s.db)
 userFCMMessaging := userRepositories.NewUserFCMMessaging()

 userUsecase := userUsecases.NewUserUsecaseImpl(
  userPostgresRepository,
  userFCMMessaging,
 )

 userHttpHandler := userHandlers.NewUserHttpHandler(userUsecase)

 // Routers
 userRouters := s.app.Group("v1/user")
 userRouters.POST("", userHttpHandler.DetectUser)
}