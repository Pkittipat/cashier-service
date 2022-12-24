package main

import (
	"net/http"
	"strconv"

	"github.com/Pkittipat/cashier-service/cmd"
	"github.com/Pkittipat/cashier-service/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	err := Run()
	if err != nil {
		panic(err)
	}
}

func Run() error {
	container := cmd.BuildContainer()

	var configInstance *config.Config
	if err := container.Invoke(func(_config *config.Config) {
		configInstance = _config
	}); err != nil {
		panic(err)
	}

	app := gin.New()
	setupMiddlewares(app, container)
	setupRoutes(app, container)

	err := app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))
	return err
}

func setupMiddlewares(app *gin.Engine, container *dig.Container) {
}

func setupRoutes(app *gin.Engine, container *dig.Container) {
	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
}
