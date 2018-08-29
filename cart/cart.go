package cart

import "github.com/rishflab/shopping_cart/inventory"

type cart struct {
	items map[string]inventory.Item
	*inventory.Inventory
}

func (c *cart) AddItem(item inventory.Item) {

}

func (c *cart) RemoveItem(item inventory.Item) {

}
