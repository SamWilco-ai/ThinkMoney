package main

import (
	"fmt"
	"log"

	"thinkmoney.com/supermarketsimulator/supermarket"
)

func main() {

	//initalise our defined SKU's and Pricing Structs
	prices := map[string]supermarket.Pricing{
		"A": {
			UnitPrice:     10,
			MagicQuantity: 3,
			MagicPrice:    25,
		},
		"B": {
			UnitPrice:     5,
			MagicQuantity: 4,
			MagicPrice:    15,
		},
		"C": {
			UnitPrice:     20,
			MagicQuantity: 2,
			MagicPrice:    25,
		},
		"D": {
			UnitPrice:     100000,
			MagicQuantity: 2,
			MagicPrice:    10,
		},
	}

	// create the pointer if our pricing rules and empty cart
	checkout := supermarket.NewCheckout(prices)

	// the items we're purchasing on this fine day
	purchasedGoods := []string{"A", "A", "A", "E", "A", "B", "B"}

	// loop over the items in our basket to scan them
	for _, item := range purchasedGoods {
		if err := checkout.Scan(item); err != nil {
			log.Println(err)
		}
	}

	// happy shopping done now for the worst bit
	total, err := checkout.GetTotalPrice()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("total price comes to: %d", total)

}
