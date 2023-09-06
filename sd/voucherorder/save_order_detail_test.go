package voucherorder

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetTaxUnitPrice(t *testing.T) {
	tests := []struct {
		givenDetail   SaveOrderDetail
		givenPrice    float64
		givenExchRate float64
		want          SaveOrderDetail
	}{
		{
			givenDetail:   SaveOrderDetail{Qty: 2, TaxRate: 9.0},
			givenPrice:    1000.0,
			givenExchRate: 0.6,
			want: SaveOrderDetail{
				Qty:             2,
				OriTaxUnitPrice: 1000.0,
				OriSum:          2000.0,
				TaxRate:         9.0,
				OriTax:          165.14,
				OriUnitPrice:    917.43,
				OriMoney:        1834.86,
				NatTaxUnitPrice: 1000.0,
				NatSum:          2000.0,
				NatTax:          165.14,
				NatUnitPrice:    917.43,
				NatMoney:        1834.86,
				OrderDetailPrices: SaveOrderDetailPrice{
					OriTaxUnitPrice: 1000.0,
					OriSum:          2000.0,
					OriTax:          165.14,
					OriUnitPrice:    917.43,
					OriMoney:        1834.86,
					NatTaxUnitPrice: 600.0,
					NatSum:          1200.00,
					NatTax:          99.08,
					NatUnitPrice:    550.46,
					NatMoney:        1100.92,
				},
			},
		},
		{
			givenDetail:   SaveOrderDetail{Qty: 1022, TaxRate: 13.0},
			givenPrice:    76.41,
			givenExchRate: 6.4316,
			want: SaveOrderDetail{
				Qty:             1022,
				OriTaxUnitPrice: 76.41,
				OriSum:          78091.02,
				TaxRate:         13.0,
				OriTax:          8983.92,
				OriUnitPrice:    67.62,
				OriMoney:        69107.10,
				NatTaxUnitPrice: 491.43855,
				NatSum:          502250.20,
				NatTax:          57781,
				NatUnitPrice:    434.90139,
				NatMoney:        444469.2,
				OrderDetailPrices: SaveOrderDetailPrice{
					OriTaxUnitPrice: 76.41,
					OriSum:          78091.02,
					OriTax:          8983.92,
					OriUnitPrice:    67.62,
					OriMoney:        69107.10,
					NatTaxUnitPrice: 491.43855,
					NatSum:          502250.20,
					NatTax:          57781,
					NatUnitPrice:    434.90139,
					NatMoney:        444469.2,
				},
			},
		},
	}

	for _, tc := range tests {
		tc.givenDetail.SetTaxUnitPrice(tc.givenPrice, tc.givenExchRate)
		b, _ := json.Marshal(tc.givenDetail)
		log.Printf("detail: %v", string(b))
		assert.Equal(t, tc.want.OriTaxUnitPrice, tc.givenDetail.OriTaxUnitPrice)
		assert.Equal(t, tc.want.OriSum, tc.givenDetail.OriSum)
		assert.Equal(t, tc.want.TaxRate, tc.givenDetail.TaxRate)
		assert.Equal(t, tc.want.OriUnitPrice, tc.givenDetail.OriUnitPrice)
		assert.Equal(t, tc.want.OriMoney, tc.givenDetail.OriMoney)
		assert.Equal(t, tc.want.OrderDetailPrices.OriMoney, tc.givenDetail.OrderDetailPrices.OriMoney)
		assert.Equal(t, tc.want.OrderDetailPrices.OriTax, tc.givenDetail.OrderDetailPrices.OriTax)
		assert.Equal(t, tc.want.OrderDetailPrices.NatSum, tc.givenDetail.OrderDetailPrices.NatSum)
		assert.Equal(t, tc.want.OrderDetailPrices.NatMoney, tc.givenDetail.OrderDetailPrices.NatMoney)
		assert.Equal(t, tc.want.OrderDetailPrices.NatTax, tc.givenDetail.OrderDetailPrices.NatTax)
	}
}
