package main

import "fmt"

type AddItem struct {
	itemType  string
	quantity  uint
	timestamp int
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

type Item struct {
	itemType string
	quantity uint
}

type Cart struct {
	items map[string]uint
}

func (cart *Cart) CheckoutPrice() float64 {

	if cart.items["trousers"] == 2 {
		return 1.0
	} else {
		return 2.
	}

}

func (cart Cart) filterByItemName(itemName string) {
	m := len(cart.items)
	fmt.Printf("%d", m)
}

type EventStore struct {
	itemAddedEvents   []ItemAdded
	itemRemovedEvents []ItemRemoved
}

func (events *EventStore) numberOfAnItemTypeInCart(itemType string) uint {

	var count uint = 0

	for _, itemAdded := range events.itemAddedEvents {
		if itemAdded.itemType == itemType {
			count++
		}
	}

	for _, ItemRemoved := range events.itemRemovedEvents {
		if ItemRemoved.itemType == itemType {
			count--
		}
	}

	return count
}

func (events *EventStore) AddItemToCart(addItem ItemAdded) {

	events.itemAddedEvents = append(events.itemAddedEvents, addItem)
}

func (events *EventStore) RemoveItemFromCart(removeItem ItemRemoved) {

	if removeItem.quantity >= events.numberOfAnItemTypeInCart(removeItem.itemType) {

		events.itemRemovedEvents = append(events.itemRemovedEvents, removeItem)

	}
}

func (events *EventStore) ViewCart(removeItem ItemRemoved) Cart {

	Car
	for _, itemAdded := range events.itemAddedEvents {

	}

}

func main() {

	a := []ItemAdded{}

	i := ItemAdded{"trousers", 1, 109201929}

	a = append(a, i)

	fmt.Println(a)

}
