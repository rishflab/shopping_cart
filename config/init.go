package config

import (
	"github.com/rishflab/shopping_cart/cart"
)

func ActivePromotions() []func(cart.Cart) float64 {

	return []func(cart.Cart) float64{cart.BeltAre15PercentOffIf2OrMoreTrousers, cart.ShirtsAre45DollarsIf2OrMoreShirts}
}
