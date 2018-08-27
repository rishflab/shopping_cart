package main

import (
	"fmt"
)

type addItem struct {
	itemType  string
	quantity  uint
	timestamp int
}

type pricesUpdated struct {
	priceList map[string]float64
}

type removeItem struct {
	itemType  string
	quantity  uint
	timestamp int
}

type itemAdded struct {
	itemType  string
	quantity  uint
	timestamp int
}

type itemRemoved struct {
	itemType  string
	quantity  uint
	timestamp int
}

type cartItem struct {
	quantity uint
	price    float64
}

type cart struct {
	items map[string]cartItem
}

// func (cart *Cart) CheckoutPrice() float64 {

// 	if cart.items["trousers"] == 2 {
// 		return 1.0
// 	} else {
// 		return 2.
// 	}

// }

func (cart cart) filterByItemName(itemName string) {
	m := len(cart.items)
	fmt.Printf("%d", m)
}

type eventStore struct {
	itemAdded     []itemAdded
	itemRemoved   []itemRemoved
	pricesUpdated []pricesUpdated
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

func main() {

	a := []itemAdded{}

	p := pricesUpdated{map[string]float64{"trousers": 20.0, "shirts": 10.0}}

	i := itemAdded{"shirts", 1, 109201929}

	b := itemAdded{"trousers", 1, 109400100}

	// f := itemRemoved{"trousers", 1, 102919293}

	// g := itemRemoved{"trousers", 1, 102919293}

	a = append(a, i)

	//fmt.Println(a)

	events := eventStore{itemAdded: []itemAdded{}, itemRemoved: []itemRemoved{}}

	events.updatePrices(p)
	//fmt.Println(events.buildCart())
	events.addItem(i)
	events.addItem(b)
	// events.removeItem(f)
	// events.removeItem(g)
	fmt.Println(events.quantityOfItem("shirts"))
	//fmt.Println(events.buildCart())
	fmt.Println(events)

	fmt.Println(events.buildCart())
	//fmt.Println(events.buildCart())

}
