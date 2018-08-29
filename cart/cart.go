package cart

import (
	"fmt"

	"github.com/rishflab/shopping_cart/inventory"
)

type Cart struct {
	quantity  map[string]uint
	inventory *inventory.Inventory
}

func (cart *Cart) AddItem(name string, stock uint) error {

	inventoryStock, err := cart.inventory.GetStock(name)

	if err != nil {
		return &itemNotFoundError{name}
	}

	cartStock, some := cart.quantity[name]

	if some == false {
		return &itemNotFoundError{name}
	}

	if stock+cartStock > inventoryStock {
		return &itemNotFoundError{name}
	}

	cart.quantity[name] = stock + cartStock

	return nil
}

func (cart *Cart) RemoveItem(name string, quantity uint) error {

	cartQuantity, some := cart.quantity[name]

	if some != true {
		return &itemNotFoundError{name}
	}

	if quantity > cartQuantity {
		return &itemNotFoundError{name}
	}

	cart.quantity[name] = cartQuantity - quantity

	return nil

}

func (cart *Cart) GetPrice(name string) (float64, error) {

	price, err := cart.inventory.GetPrice(name)

	if err != nil {
		return -1.0, &itemNotFoundError{name}
	}

	return price, nil

}

func (cart *Cart) PriceAfterPromotions(promotions []Promotion) float64 {

}

type itemNotFoundError struct {
	itemName string
}

func (e *itemNotFoundError) Error() string {
	return fmt.Sprintf("%s not found in inventory", e.itemName)
}
