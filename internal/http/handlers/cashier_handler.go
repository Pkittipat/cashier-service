package handlers

import (
	"net/http"

	"github.com/Pkittipat/cashier-service/internal/http/requests"
	"github.com/Pkittipat/cashier-service/internal/http/responses"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
	"github.com/gin-gonic/gin"
)

type CashierHandler interface {
	GetInventory(c *gin.Context)
	Purchase(c *gin.Context)
}

type cashierHandler struct {
	inventoryNode *inventory.Inventory
}

func NewCashierHandler(
	inventoryNode *inventory.Inventory,
) CashierHandler {
	return &cashierHandler{
		inventoryNode: inventoryNode,
	}
}

func (h *cashierHandler) GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, responses.NewInvetory(h.inventoryNode.GetInventory()))
	return
}

func (h *cashierHandler) Purchase(c *gin.Context) {
	var request requests.PurchaseForm
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := CalculateChange(request.Price, request.Payment, h.inventoryNode.Root)
	c.JSON(http.StatusOK, result)
	return
}

// CalculateChange calculates the change to be given to the customer and determines the optimal combination of bank notes and coins to give as change.
func CalculateChange(price float64, payment float64, root *inventory.Node) *responses.Purchase {
	change := payment - price
	node := root
	breakdown := make([]*responses.Breakdown, 0)
	for change > 0 && node != nil {
		if change >= node.Value && node.Amount > 0 {
			change -= node.Value
			node.Amount -= 1
			breakdown = append(breakdown, &responses.Breakdown{
				Value:  node.Value,
				Amount: 1,
			})
		} else {
			node = node.Left
		}
	}
	return &responses.Purchase{
		Change:    change,
		Breakdown: breakdown,
	}
}
