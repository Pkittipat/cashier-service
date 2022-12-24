package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashierHandler interface {
	GetInventory(c *gin.Context)
	Purchase(c *gin.Context)
}

type cashierHandler struct {
}

func NewCashierHandler() CashierHandler {
	return &cashierHandler{}
}

func (h *cashierHandler) GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}

func (h *cashierHandler) Purchase(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}
