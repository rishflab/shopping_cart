package main

import "fmt"

type AddItem struct {
	itemType  string
	quantity  uint
	timestamp int
}

type PricesUpdated struct {
	priceList map[string]float64
}

type RemoveItem struct {
	itemType  string
	quantity  uint
	timestamp int
}

type ItemAdded struct {
	itemType  string
	quantity  uint
	timestamp int
}

type ItemRemoved struct {
	itemType  string
	quantity  uint
	timestamp int
}

type CartItem struct {
	quantity uint
	price    uint
}

type Cart struct {
	items map[string]CartItem
}

// func (cart *Cart) CheckoutPrice() float64 {

// 	if cart.items["trousers"] == 2 {
// 		return 1.0
// 	} else {
// 		return 2.
// 	}

// }

func (cart Cart) filterByItemName(itemName string) {
	m := len(cart.items)
	fmt.Printf("%d", m)
}

type EventStore struct {
	itemAdded     []ItemAdded
	itemRemoved   []ItemRemoved
	pricesUpdated []PricesUpdated
}

func (events *EventStore) numberOfAnItemTypeInCart(itemType string) uint {

	var count uint = 0

	for _, itemAdded := range events.itemAdded {
		if itemAdded.itemType == itemType {
			count++
		}
	}

	for _, ItemRemoved := range events.itemRemoved {
		if ItemRemoved.itemType == itemType {
			count--
		}
	}

	return count
}

func (events *EventStore) SetPrices(updatePrices PricesUpdated) {

	events.setPricesEvents = append(events.updatePrices, updatePriceList)

}

func (events *EventStore) AddItem(addItem ItemAdded) {

	events.itemAddedEvents = append(events.itemAdded, addItem)
}

func (events *EventStore) RemoveItem(removeItem ItemRemoved) {

	if removeItem.quantity >= events.numberOfAnItemTypeInCart(removeItem.itemType) {

		events.itemRemovedEvents = append(events.itemRemoved, removeItem)

	}
}

func (events *EventStore) PriceOfItem(itemType string) {

	currentPrices := events.updatePriceListEvents[len(events.updatePriceListEvents-1)]
	return currentPrices[itemType]

}

func (events *EventStore) BuildCart() map[string]CartItem {

	cart := make(map[string]CartItem)

	for _, itemAdded := range events.itemAddedEvents {

		cartItem, exists := cart[itemAdded.itemType]

		if exists == true {
			cartItem = CartItem{quantity: cartItem.quantity + itemAdded.quantity, price: events.PriceOfItem(itemAdded.itemType)}
		} else {
			cartItem = CartItem{quantity: itemAdded.quantity, price: events.PriceOfItem(itemAdded.itemType)}
		}

		cart[itemAdded.itemType] = cartItem

	}

	for _, itemRemoved := range events.itemRemovedEvents {

		cartItem, exists := cart[itemRemoved.itemType]

		if exists == true {
			if cartItem.quantity-itemRemoved.quantity == 0 {
				delete(cart, itemRemoved.itemType)
			} else {
				cartItem = CartItem{quantity: cartItem.quantity - itemRemoved.quantity, price: 20.0}
				cart[itemRemoved.itemType] = cartItem
			}
		}

	}

	return cart

}

func main() {

	a := []ItemAdded{}

	i := ItemAdded{"director", 1, 109201929}

	b := ItemAdded{"bond", 1, 109400100}

	f := ItemRemoved{"bond", 1, 102919293}

	g := ItemRemoved{"bond", 1, 102919293}

	a = append(a, i)

	fmt.Println(a)

	events := EventStore{itemAddedEvents: []ItemAdded{}, itemRemovedEvents: []ItemRemoved{}}

	events.AddItem(i)
	events.AddItem(b)
	events.RemoveItem(f)
	events.RemoveItem(g)
	fmt.Println(events)
	fmt.Println(events.BuildCart())

}
