package usecases

import (
	"github.com/Pkittipat/cashier-service/internal/http/responses"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
)

type CashierUsecase interface {
	CalculateChange(
		price float64,
		payment float64,
	) (*responses.Purchase, error)
}

type cashierUsecase struct {
	inventoryNode *inventory.Inventory
}

func NewCashierUsecase(
	inventoryNode *inventory.Inventory,
) CashierUsecase {
	return &cashierUsecase{
		inventoryNode: inventoryNode,
	}
}

func (c *cashierUsecase) CalculateChange(
	price float64,
	payment float64,
) (*responses.Purchase, error) {
	change := payment - price
	node := c.inventoryNode.Root
	breakdownMap := make(map[float64]*responses.Breakdown)
	for change > 0 && node != nil {
		if change >= node.Value && node.Amount > 0 {
			change -= node.Value
			node.Amount -= 1
			_, ok := breakdownMap[node.Value]
			if !ok {
				breakdownMap[node.Value] = &responses.Breakdown{
					Value:  node.Value,
					Amount: 0,
				}

			}
			breakdownMap[node.Value].Amount += 1
		} else {
			node = node.Left
		}
	}

	breakdown := make([]*responses.Breakdown, 0)
	for _, val := range breakdownMap {
		breakdown = append(breakdown, val)
	}
	return &responses.Purchase{
		Change:    (payment - price),
		Breakdown: breakdown,
	}, nil
}
