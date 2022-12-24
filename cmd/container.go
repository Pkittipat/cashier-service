package cmd

import (
	"github.com/Pkittipat/cashier-service/config"
	"github.com/Pkittipat/cashier-service/internal/http/handlers"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)

	// Handlers
	_ = container.Provide(handlers.NewCashierHandler)

	return container
}
