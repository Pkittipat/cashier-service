package responses

import "github.com/Pkittipat/cashier-service/pkg/inventory"

type Inventory struct {
	Value  float64 `json:"value"`
	Amount int     `json:"amount"`
}

func NewInvetory(inventory []*inventory.Node) (out []*Inventory) {
	for _, node := range inventory {
		out = append(out, &Inventory{
			Value:  node.Value,
			Amount: node.Amount,
		})
	}
	return out
}

// Purcahse
// Returns a JSON response with the following parameters:
// change: the amount of change to be given to the customer (float)
// breakdown: a list of dictionaries representing the bank notes and coins to be given to the customer, with the following keys:
type Purchase struct {
	Change    float64      `json:"change"`
	Breakdown []*Breakdown `json:"breakdown"`
}

// value: the value of the bank note or coin (float)
// amount: the number of bank notes or coins of this value to be given (int)
type Breakdown struct {
	Value  float64 `json:"value"`
	Amount int     `json:"amount"`
}
