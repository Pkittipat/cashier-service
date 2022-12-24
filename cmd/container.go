package cmd

import (
	"github.com/Pkittipat/cashier-service/config"
	"github.com/Pkittipat/cashier-service/internal/http/handlers"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(inventory.NewInventory)

	// Handlers
	_ = container.Provide(handlers.NewCashierHandler)

	return container
}
