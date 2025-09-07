package supermarket

import (
	"errors"
)

type ICheckout interface {
	Scan(SKU string) error
	GetTotalPrice() (int, error)
}

// the struct holding our information of price, special price and the threshold to achieve the special
type Pricing struct {
	UnitPrice     int
	MagicQuantity int
	MagicPrice    int
}

// the struct to implement our interface
type Checkout struct {
	Prices map[string]Pricing
	Cart   map[string]int
}

// the function to create our pointer of the checkout containing known prices and the object we posess(?) at that point
func NewCheckout(prices map[string]Pricing) *Checkout {
	return &Checkout{
		Prices: prices,
		Cart:   make(map[string]int),
	}
}

// the function of our checkout type to satisfy the interface for scanning the items to see if they match what we know
func (c *Checkout) Scan(SKU string) error {
	// check to see if the SKU we're being provided exists in our known map of prices
	if _, ok := c.Prices[SKU]; !ok {
		return errors.New("unexpected item in the bagging area")
	}
	// if it does exist increment the quantity
	c.Cart[SKU]++
	return nil
}

// the function of our checkout type to satisfy the interface to calculate the total price
func (c *Checkout) GetTotalPrice() (int, error) {
	totalPrice := 0

	// range over every item in our SKU and the alligning quantity in our cart
	for SKU, count := range c.Cart {

		// get the prices, no need for an exists check as this is done in the scan function
		price := c.Prices[SKU]

		// if there is a "magic quantity" then we need to take it into consideration when calculating
		if price.MagicQuantity > 0 {
			// get num of full specials that could be applied
			discountedProductsPrice := price.MagicPrice * (count / price.MagicQuantity)

			//get the remainder to apply singular unit pricing to
			fullPriceProductsPrice := price.UnitPrice * (count % price.MagicQuantity)

			// add it all
			totalPrice += discountedProductsPrice + fullPriceProductsPrice
		} else {
			totalPrice += count * price.UnitPrice
		}

	}
	return totalPrice, nil
}
