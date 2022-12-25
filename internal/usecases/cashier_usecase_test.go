package usecases

import (
	"testing"

	"github.com/Pkittipat/cashier-service/internal/http/responses"
	"github.com/Pkittipat/cashier-service/pkg/inventory"
)

func TestCalculateBrealdown(t *testing.T) {
	inventory := inventory.NewInventory()
	type args struct {
		price   float64
		payment float64
	}
	testcases := []struct {
		name string
		args args
		want map[float64]*responses.Breakdown
	}{
		{
			name: "case_change_245",
			args: args{price: 255, payment: 500},
			want: map[float64]*responses.Breakdown{
				100: &responses.Breakdown{
					Value:  100,
					Amount: 2,
				},
				20: &responses.Breakdown{
					Value:  20,
					Amount: 2,
				},
				5: &responses.Breakdown{
					Value:  5,
					Amount: 1,
				},
			},
		},
		{
			name: "case_change_245.50",
			args: args{price: 255.50, payment: 500},
			want: map[float64]*responses.Breakdown{
				100: &responses.Breakdown{
					Value:  100,
					Amount: 2,
				},
				20: &responses.Breakdown{
					Value:  20,
					Amount: 2,
				},
				1: &responses.Breakdown{
					Value:  1,
					Amount: 4,
				},
				0.25: &responses.Breakdown{
					Value:  0.25,
					Amount: 2,
				},
			},
		},
		{
			name: "case_change_gather_than_total_amount",
			args: args{price: 255.50, payment: 10000000},
			want: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := calculateBrealdown(inventory, tc.args.price, tc.args.payment)
			for key, elm := range got {
				val, ok := tc.want[key]
				if !ok {
					t.Errorf("ecpected: %+v, got: %+v", tc.want[key], key)
				}
				if elm.Amount != val.Amount {
					t.Errorf("ecpected: %+v, got: %+v", tc.want, got)
				}
			}
		})
	}
}
