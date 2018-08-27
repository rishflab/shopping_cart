package main

import (
	"fmt"
)

type addItem struct {
	itemType string
	quantity uint
}

type inventoryUpdated struct{}

type removeItem struct {
	itemType string
	quantity uint
}

type itemAdded struct {
	itemType string
	quantity uint
}

type itemRemoved struct {
	itemType string
	quantity uint
}

type cartItem struct {
	quantity uint
	price    float64
}

type cart struct {
	items map[string]cartItem
}

type inventory struct {
	items map[string]uint
}

type pricesUpdated struct {
	priceList map[string]float64
}

func (cart *cart) CheckoutPrice() float64 {

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

	return total
}

type eventStore struct {
	itemAdded        []itemAdded
	itemRemoved      []itemRemoved
	pricesUpdated    []pricesUpdated
	inventorySet []inventorySet
}

func (events *eventStore) quantityOfItem(itemType string) uint {

	return events.buildCart().items[itemType].quantity
}

func (events *eventStore) updatePrices(updatePrices pricesUpdated) {

	events.pricesUpdated = append(events.pricesUpdated, updatePrices)

}

func (events *eventStore) addItem(addItem itemAdded) (bool, error) {

	_, err := events.priceOfItem(addItem.itemType)

	if err != nil {
		return false, err
	}

	if events.quantityOfItem("add")

	events.itemAdded = append(events.itemAdded, addItem)
	return true, err

}

func (events *eventStore) removeItem(removeItem itemRemoved) {

	if removeItem.quantity <= events.quantityOfItem(removeItem.itemType) {
		events.itemRemoved = append(events.itemRemoved, removeItem)
	}
}

type priceNotFoundError struct {
	itemType string
}

func (e *priceNotFoundError) Error() string {
	return fmt.Sprintf("No price for %s", e.itemType)
}

func (events *eventStore) priceOfItem(itemType string) (float64, error) {

	if (len(events.pricesUpdated)) < 1 {
		return -1.0, &priceNotFoundError{itemType}
	}
	currentPrices := events.pricesUpdated[len(events.pricesUpdated)-1]
	price, exists := currentPrices.priceList[itemType]

	if exists == false {
		return -1.0, &priceNotFoundError{itemType}
	}

	return price, nil

}

func (events *eventStore) buildCart() cart {

	c := cart{make(map[string]cartItem)}

	for _, itemAdded := range events.itemAdded {

		price, _ := events.priceOfItem(itemAdded.itemType)

		item, exists := c.items[itemAdded.itemType]

		if exists == true {
			item = cartItem{quantity: item.quantity + itemAdded.quantity, price: price}
		} else {
			item = cartItem{quantity: itemAdded.quantity, price: price}
		}

		c.items[itemAdded.itemType] = item

	}

	for _, itemRemoved := range events.itemRemoved {

		item, exists := c.items[itemRemoved.itemType]

		if exists == true {
			if item.quantity-itemRemoved.quantity == 0 {
				delete(c.items, itemRemoved.itemType)
			} else {
				price, _ := events.priceOfItem(itemRemoved.itemType)
				item = cartItem{quantity: item.quantity - itemRemoved.quantity, price: price}
				c.items[itemRemoved.itemType] = item
			}
		}

	}

	return c
}

func initEventStore() eventStore {

	events := eventStore{itemAdded: []itemAdded{}, itemRemoved: []itemRemoved{}, pricesUpdated: []pricesUpdated{}}
	events.updatePrices(pricesUpdated{map[string]float64{"belts": 20.0, "shirts": 60.0, "suits": 300.0, "trousers": 70.0, "shoes": 120.0, "ties": 20.0}})
	events.addItem(itemAdded{"belts", 10})
	events.addItem(itemAdded{"shirts", 5})
	events.addItem(itemAdded{"suits", 2})
	events.addItem(itemAdded{"trousers", 4})
	events.addItem(itemAdded{"shoes", 1})
	events.addItem(itemAdded{"ties", 8})
	return events
}

func main() {

	// a := []itemAdded{}

	// p := pricesUpdated{map[string]float64{"trousers": 20.0, "shirts": 10.0}}

	// i := itemAdded{"shirts", 1}

	// b := itemAdded{"trousers", 1}

	// // f := itemRemoved{"trousers", 1}

	// // g := itemRemoved{"trousers", 1}

	// a = append(a, i)

	// //fmt.Println(a)

	// events := eventStore{itemAdded: []itemAdded{}, itemRemoved: []itemRemoved{}}

	// events.updatePrices(p)
	// //fmt.Println(events.buildCart())
	// events.addItem(i)
	// events.addItem(b)
	// // events.removeItem(f)
	// // events.removeItem(g)
	// fmt.Println(events.quantityOfItem("shirts"))
	// //fmt.Println(events.buildCart())
	// fmt.Println(events)
	events := initEventStore()
	fmt.Println(events.buildCart())
	//fmt.Println(events.buildCart())

}
