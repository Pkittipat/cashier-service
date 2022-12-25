package handlers

import (
	"errors"
	"net/http"

	"github.com/Pkittipat/cashier-service/internal/http/requests"
	"github.com/Pkittipat/cashier-service/internal/http/responses"
	"github.com/Pkittipat/cashier-service/internal/usecases"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
	"github.com/gin-gonic/gin"
)

type CashierHandler interface {
	GetInventory(c *gin.Context)
	Purchase(c *gin.Context)
}

type cashierHandler struct {
	inventoryNode *inventory.Inventory
	usecase       usecases.CashierUsecase
}

func NewCashierHandler(
	inventoryNode *inventory.Inventory,
	usecase usecases.CashierUsecase,
) CashierHandler {
	return &cashierHandler{
		inventoryNode: inventoryNode,
		usecase:       usecase,
	}
}

func (h *cashierHandler) GetInventory(c *gin.Context) {
	totalAmount := h.inventoryNode.TotalAmount()
	data := responses.NewInvetory(h.inventoryNode.GetInventory(), totalAmount)
	responses.NewResponse(data).Response(c, http.StatusOK)
	return
}

func (h *cashierHandler) Purchase(c *gin.Context) {
	var request requests.PurchaseForm
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.NewErrorResponse(err).Response(c, http.StatusBadRequest)
		return
	}
	if err := request.Validate(); err != nil {
		responses.NewErrorResponse(err).Response(c, http.StatusBadRequest)
		return
	}

	result, err := h.usecase.CalculateChange(request.Price, request.Payment)
	if err != nil {
		if errors.Is(err, inventory.ErrTotalAmountNotEnough) {
			responses.NewErrorResponse(err).Response(c, http.StatusBadRequest)
			return
		}
		responses.NewErrorResponse(err).Response(c, http.StatusInternalServerError)
		return
	}
	responses.NewResponse(result).Response(c, http.StatusOK)
	return
}
