package main

import (
	"os"
	"poly_news/cmd/core-server/flag"
	"poly_news/internal/config"
	"poly_news/internal/handler"
	"poly_news/internal/repository/gorm"
	"poly_news/internal/repository/gorm/model"
	"poly_news/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func main() {
	coreServerFlag := lo.Must1(flag.ParseFlags(os.Args[1:]))
	coreServerConfig := lo.Must1(config.ParseConfig[config.CoreServerConfig](coreServerFlag.ConfigPath))

	// Create a new GORM database
	dbConn := lo.Must1(gorm.New(coreServerConfig.Database))
	lo.Must0(dbConn.AutoMigrate(
		&model.Subscribe{},
	))

	// Create repositories
	subscribeRepository := gorm.NewSubscribe()

	// Create usecases
	subscribeUsecase := usecase.NewSubscribe(dbConn, subscribeRepository)

	// Create handlers
	subscribeHandler := handler.NewSubscribe(subscribeUsecase)

	// Set middlewares
	ginMiddlewares := []gin.HandlerFunc{
		gin.Logger(),
	}

	// Initialize gin
	engine := initializeGinEngine(coreServerConfig.Server, ginMiddlewares...)

	handler.Initialize(
		engine,
		subscribeHandler,
	)

	if err := engine.Run(coreServerConfig.Server.Listen); err != nil {
		panic(err)
	}
}

func initializeGinEngine(cfg config.ServerConfig, middlewares ...gin.HandlerFunc) *gin.Engine {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(middlewares...)

	return engine
}
