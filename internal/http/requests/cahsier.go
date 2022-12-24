package requests

// Purchase
// Takes a JSON request body with the following parameters:
// price: the price of the product being purchased (float)
// payment: the total payment made by the customer (float)
type PurchaseForm struct {
	Price   float64 `json:"price"`
	Payment float64 `json:"payment"`
}
