package supermarket

import (
	"testing"
)

func TestScan(t *testing.T) {

	type scanTestCase struct {
		name         string
		prices       map[string]Pricing
		scanned      []string
		expectedCart map[string]int
		expectedErr  error
	}

	tests := []scanTestCase{
		{
			name: "Valid SKU added to cart",
			prices: map[string]Pricing{
				"A": {UnitPrice: 10},
			},
			scanned:      []string{"A"},
			expectedCart: map[string]int{"A": 1},
		},
		{
			name: "Multiple valid SKUs",
			prices: map[string]Pricing{
				"A": {UnitPrice: 10},
				"B": {UnitPrice: 5},
			},
			scanned:      []string{"A", "B", "A"},
			expectedCart: map[string]int{"A": 2, "B": 1},
		},
		{
			name: "Invalid SKUs",
			prices: map[string]Pricing{
				"A": {UnitPrice: 10},
			},
			scanned:      []string{"A", "Z"},
			expectedCart: map[string]int{"A": 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testCheckout := NewCheckout(tc.prices)

			for _, sku := range tc.scanned {
				err := testCheckout.Scan(sku)
				if tc.expectedErr != nil {
					if err != tc.expectedErr {
						t.Errorf("expected error %v got %v", tc.expectedErr, err)
					}
				}
			}

			for sku, quantity := range tc.expectedCart {
				if testCheckout.Cart[sku] != quantity {
					t.Errorf("expected %d of product %s, got %d", quantity, sku, testCheckout.Cart[sku])
				}
			}
		})
	}
}

func TestTotal(t *testing.T) {

	type scanTestCase struct {
		name          string
		prices        map[string]Pricing
		scanned       []string
		expectedCart  map[string]int
		expectedTotal int
		expectedErr   error
	}

	tests := []scanTestCase{
		{
			name: "no items scanned - no price",
			prices: map[string]Pricing{
				"A": {
					UnitPrice: 10,
				},
			},
			scanned:       []string{},
			expectedCart:  map[string]int{},
			expectedTotal: 0,
		},
		{
			name: "sinlge item scanned - no discount",
			prices: map[string]Pricing{
				"A": {
					UnitPrice: 10,
				},
			},
			scanned:       []string{"A", "A"},
			expectedCart:  map[string]int{"A": 2},
			expectedTotal: 20,
		},
		{
			name: "two items scanned - no discount",
			prices: map[string]Pricing{
				"A": {
					UnitPrice: 10,
				},
				"B": {
					UnitPrice: 40,
				},
			},
			scanned:       []string{"A", "A", "B"},
			expectedCart:  map[string]int{"A": 2, "B": 1},
			expectedTotal: 60,
		},
		{
			name: "two items scanned - one discount - one no discount",
			prices: map[string]Pricing{
				"A": {
					UnitPrice:     10,
					MagicQuantity: 2,
					MagicPrice:    15,
				},
				"B": {
					UnitPrice: 40,
				},
			},
			scanned:       []string{"A", "A", "B"},
			expectedCart:  map[string]int{"A": 2, "B": 1},
			expectedTotal: 55,
		},
		{
			name: "one item scanned - one discount with remainder",
			prices: map[string]Pricing{
				"A": {
					UnitPrice:     10,
					MagicQuantity: 2,
					MagicPrice:    15,
				},
			},
			scanned:       []string{"A", "A", "A"},
			expectedCart:  map[string]int{"A": 3},
			expectedTotal: 25,
		},
		{
			name: "two items scanned - two discounts",
			prices: map[string]Pricing{
				"A": {
					UnitPrice:     10,
					MagicQuantity: 2,
					MagicPrice:    15,
				},
				"B": {
					UnitPrice:     20,
					MagicQuantity: 2,
					MagicPrice:    30,
				},
			},
			scanned:       []string{"A", "A", "B", "B"},
			expectedCart:  map[string]int{"A": 2, "B": 2},
			expectedTotal: 45,
		},
		{
			name: "two items scanned - two discounts - one remainder",
			prices: map[string]Pricing{
				"A": {
					UnitPrice:     10,
					MagicQuantity: 2,
					MagicPrice:    15,
				},
				"B": {
					UnitPrice:     20,
					MagicQuantity: 2,
					MagicPrice:    30,
				},
			},
			scanned:       []string{"A", "A", "B", "B", "A"},
			expectedCart:  map[string]int{"A": 3, "B": 2},
			expectedTotal: 55,
		},
		{
			name: "two items scanned - two discounts - two remainder",
			prices: map[string]Pricing{
				"A": {
					UnitPrice:     10,
					MagicQuantity: 2,
					MagicPrice:    15,
				},
				"B": {
					UnitPrice:     20,
					MagicQuantity: 2,
					MagicPrice:    30,
				},
			},
			scanned:       []string{"A", "A", "B", "B", "A", "B"},
			expectedCart:  map[string]int{"A": 3, "B": 3},
			expectedTotal: 75,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testCheckout := NewCheckout(tc.prices)

			for _, sku := range tc.scanned {
				err := testCheckout.Scan(sku)
				if tc.expectedErr != nil {
					if err != tc.expectedErr {
						t.Errorf("expected error %v got %v", tc.expectedErr, err)
					}
				}
			}

			for sku, quantity := range tc.expectedCart {
				if testCheckout.Cart[sku] != quantity {
					t.Errorf("expected %d of product %s, got %d", quantity, sku, testCheckout.Cart[sku])
				}
			}

			total, err := testCheckout.GetTotalPrice()
			if tc.expectedErr != nil {
				if err != tc.expectedErr {
					t.Errorf("expected error %v got %v", tc.expectedErr, err)
				}
			}

			if total != tc.expectedTotal {
				t.Errorf("expected total %d, got %d", tc.expectedTotal, total)
			}
		})
	}
}
