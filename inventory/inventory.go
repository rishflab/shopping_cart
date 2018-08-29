package inventory

type Inventory struct {
	items map[string]Item
}

type Item struct {
	quantity uint
	price    float64
}
