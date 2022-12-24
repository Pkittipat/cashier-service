package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pkittipat/cashier-service/cmd"
	"github.com/Pkittipat/cashier-service/config"
	"github.com/Pkittipat/cashier-service/internal/http/handlers"
	adminHandler "github.com/Pkittipat/cashier-service/internal/http/handlers/admin"
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

	fmt.Println("HERE: ", configInstance.App.APP.DebugMode)
	// if configInstance.App.APP.DebugMode {
	// 	fmt.Println("INNN")
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	app := gin.New()
	setupMiddlewares(app, container)
	setupRoutes(app, container)

	err := app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))
	return err
}

func setupMiddlewares(app *gin.Engine, container *dig.Container) {
}

func setupRoutes(app *gin.Engine, container *dig.Container) {
	var (
		cashierHandler        handlers.CashierHandler
		adminInventoryHandler adminHandler.InventoryHandler
	)

	if err := container.Invoke(func(
		_cashierHandler handlers.CashierHandler,
		_adminInventoryHandler adminHandler.InventoryHandler,
	) {
		cashierHandler = _cashierHandler
		adminInventoryHandler = _adminInventoryHandler
	}); err != nil {
		panic(err)
	}

	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	adminRoute := app.Group("/admin")
	{
		inventoryRoute := adminRoute.Group("/inventory")
		{
			inventoryRoute.PUT(":value/:amount", adminInventoryHandler.UpdateInventory)
		}
	}

	cashierRoute := app.Group("/cashier")
	{
		cashierRoute.GET("/inventory", cashierHandler.GetInventory)
		cashierRoute.POST("/purchase", cashierHandler.Purchase)
	}
}
