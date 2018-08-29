package promotions

type Promotion struct {
	priority uint
}

total := 0.0

	for k, v := range cart.items {

		switch k {

		case "belts":

			if cart.items["trousers"].quantity >= 2 {
				total = total + v.price*0.85*float64(v.quantity)
			} else {
				total = total + v.price*float64(v.quantity)
			}

		case "shoes":

			if cart.items["trousers"].quantity >= 2 {
				total = total + v.price*0.85*float64(v.quantity)
			} else {
				total = total + v.price*float64(v.quantity)
			}

		case "shirts":

			if cart.items["shirts"].quantity >= 2 {
				total = total + v.price*2.0 + 45.0*float64(v.quantity-2)
			} else {
				total = total + v.price*float64(v.quantity)

			}

		case "ties":

			if cart.items["shirts"].quantity >= 3 {
				total = total + v.price*0.5*float64(v.quantity)
			} else {
				total = total + v.price*float64(v.quantity)

			}

		default:
			total = total + v.price*float64(v.quantity)

		}
	}

