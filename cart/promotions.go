package cart

type PromotionDiscount func(Cart) float64

func ActivePromotions() []func(Cart) float64 {

	return[] func(Cart) float64 {
		BeltAre15PercentOffIf2OrMoreTrousers,
		ShirtsAre45DollarsIf2OrMoreShirts
	}
	
}

func ApplyPromotions(cart Cart, promotions []func(Cart) float64) float64 {

	total := PriceWithoutPromotions(cart)
	discount := 0.0

	for _, f := range promotions {
		discount += f(cart, total)
	}

	return total - discount
}

func BeltAre15PercentOffIf2OrMoreTrousers(cart Cart) float64 {

	discount := 0.0

	quantity, some := cart.quantity["belts"]

	if some != true {
		return discount
	}

	beltPrice, _ := cart.GetPrice("belts")

	discount = beltPrice * float64(quantity) * 0.15

	return discount

}

func ShirtsAre45DollarsIf2OrMoreShirts(cart Cart) float64 {

	discount := 0.0

	quantity, some := cart.quantity["shirts"]

	if some != true {
		return discount
	}

	shirtPrice, _ := cart.GetPrice("shirts")

	shirtDiscount := shirtPrice - 45.0

	discount := float64(quantity-2) * shirtDiscount

	return discount
}

func PriceWithoutPromotions(cart Cart) float64 {

	total := 0.0

	for name, quantity := range cart.quantity {

		price, _ := cart.GetPrice(name)

		total = total + price*float64(quantity)

	}

	return total
}

// total := 0.0

// 	for k, v := range cart.items {

// 		switch k {

// 		case "belts":

// 			if cart.items["trousers"].quantity >= 2 {
// 				total = total + v.price*0.85*float64(v.quantity)
// 			} else {
// 				total = total + v.price*float64(v.quantity)
// 			}

// 		case "shoes":

// 			if cart.items["trousers"].quantity >= 2 {
// 				total = total + v.price*0.85*float64(v.quantity)
// 			} else {
// 				total = total + v.price*float64(v.quantity)
// 			}

// 		case "shirts":

// 			if cart.items["shirts"].quantity >= 2 {
// 				total = total + v.price*2.0 + 45.0*float64(v.quantity-2)
// 			} else {
// 				total = total + v.price*float64(v.quantity)

// 			}

// 		case "ties":

// 			if cart.items["shirts"].quantity >= 3 {
// 				total = total + v.price*0.5*float64(v.quantity)
// 			} else {
// 				total = total + v.price*float64(v.quantity)

// 			}

// 		default:
// 			total = total + v.price*float64(v.quantity)

// 		}
// 	}
