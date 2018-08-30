package inventory

import (
	"fmt"
)

//Inventory of online story items
type Inventory struct {
	stock map[string]uint
	price map[string]float64
}

func StartingInventory() Inventory {
	inv := Inventory{}
	inv.SetPriceAndStock("shirts", 100, 12)
	inv.SetPriceAndStock("belts", 50, 4)
	inv.SetPriceAndStock("trousers", 150, 8)
	return inv
}

func (inv *Inventory) SetPriceAndStock(name string, price float64, stock uint) {
	inv.stock[name] = stock
	inv.price[name] = price
}

func (inv *Inventory) GetStock(name string) (uint, error) {
	stock, some := inv.stock[name]
	if some == true {
		return stock, nil
	}
	return 0, &itemNotFoundError{name}
}

func (inv *Inventory) GetPrice(name string) (float64, error) {
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
