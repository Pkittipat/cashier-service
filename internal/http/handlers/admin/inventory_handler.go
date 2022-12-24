package admin

import (
	"net/http"
	"strconv"

	"github.com/Pkittipat/cashier-service/internal/http/responses"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
	"github.com/gin-gonic/gin"
)

type InventoryHandler interface {
	UpdateInventory(c *gin.Context)
}

type inventoryHandler struct {
	inventory *inventory.Inventory
}

func NewInventory(
	inventory *inventory.Inventory,
) InventoryHandler {
	return &inventoryHandler{
		inventory: inventory,
	}
}

func (h *inventoryHandler) UpdateInventory(c *gin.Context) {
	value, err := strconv.ParseFloat(c.Param("value"), 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount, err := strconv.Atoi(c.Param("amount"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.inventory.UpdateInventory(value, amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := gin.H{"value": value, "amount": amount}
	responses.NewResponse(data).Response(c, http.StatusOK)
	return
}
