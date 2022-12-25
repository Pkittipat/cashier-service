package responses

import "github.com/Pkittipat/cashier-service/pkg/inventory"

type Inventory struct {
	TotalAmount float64      `json:"total_amount"`
	Breakdown   []*Breakdown `json:"breakdown"`
}

func NewInvetory(breakdown []*inventory.Node, totalAmount float64) *Inventory {
	inventory := &Inventory{
		TotalAmount: totalAmount,
	}
	for _, node := range breakdown {
		inventory.Breakdown = append(inventory.Breakdown, &Breakdown{
			Value:  node.Value,
			Amount: node.Amount,
		})
	}
	return inventory
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
