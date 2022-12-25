package cmd

import (
	"github.com/Pkittipat/cashier-service/config"
	"github.com/Pkittipat/cashier-service/internal/http/handlers"
	adminHandler "github.com/Pkittipat/cashier-service/internal/http/handlers/admin"
	"github.com/Pkittipat/cashier-service/internal/usecases"
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
	_ = container.Provide(adminHandler.NewInventory)

	// usecases
	_ = container.Provide(usecases.NewCashierUsecase)

	return container
}
