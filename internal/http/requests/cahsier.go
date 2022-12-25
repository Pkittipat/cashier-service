package requests

import (
	"github.com/Pkittipat/cashier-service/internal/http/responses"
)

// Purchase
// Takes a JSON request body with the following parameters:
// price: the price of the product being purchased (float)
// payment: the total payment made by the customer (float)
type PurchaseForm struct {
	Price   float64 `json:"price" binding:"required,gte=1"`
	Payment float64 `json:"payment" binding:"required,gte=1"`
}

func (p *PurchaseForm) Validate() error {
	if p.Price > p.Payment {
		return responses.ErrInvalidPaymentLessThanPrice
	}
	return nil
}
