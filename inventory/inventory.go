package inventory

import "fmt"

type inventory struct {
	quantity map[string]uint
	price    map[string]float64
}

func (inv *inventory) setPriceAndQuantity(name string, price float64, quantity uint) {
	inv.quantity[name] = quantity
	inv.price[name] = price
}

func (inv *inventory) getQuantity(name string) (uint, error) {
	quantity, some := inv.quantity[name]
	if some == true {
		return quantity, nil
	}
	return 0, &itemNotFoundError{name}
}

func (inv *inventory) getPrice(name string) (float64, error) {
	price, some := inv.price[name]
	if some == true {
		return price, nil
	}
	return -1.0, &itemNotFoundError{name}
}

type itemNotFoundError struct {
	itemName string
}

func (e *itemNotFoundError) Error() string {
	return fmt.Sprintf("%s not found in inventory", e.itemName)
}
