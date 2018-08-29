package cart

type Promotion func(Cart, float64) float64

// type ActivePromotions struct {
// 	promotions []Promotion(*Cart, uint)
// }

// func (active *ActivePromotions) AddPromotion(f func(*Cart, uint) uint) {

// 	active.promotions = append(active.promotions, f)

// }

// func (active *ActivePromotions) Init() {

// 	active.AddPromotion(Belt15PercentOffIf2OrMoreTrousers func)

// }

func ApplyPromotions(cart Cart, promotions []func(Cart, float64) float64) float64 {

	total := PriceWithoutPromotions(cart)

	for _, f := range promotions {
		total -= f(cart, total)
	}

	return total
}

func Belt15PercentOffIf2OrMoreTrousers(cart Cart, total float64) float64 {
	
	
	for name, quantity := cart.quantity {
		switch k {

		case "belts":

			if cart.items["trousers"].quantity >= 2 {
				total = total + v.price*0.85*float64(v.quantity)
			} else {
				total = total + v.price*float64(v.quantity)
			}
		default
	}
}

func Shoes15PercentOffIf2OrMoreTrousers(a, b int) int {
	return a + b
}

func p3(a, b int) int {
	return a - b
}

func p4(a, b int) int {
	return a * b
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
