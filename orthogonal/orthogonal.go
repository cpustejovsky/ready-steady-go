package orthogonal

type Buyable interface {
	Cost() int
}

type Item struct {
	Price int
	Name  string
}

type Items []Item

func (i *Item) Cost() int {
	return i.Price
}

func (i Items) Cost() int {
	var total int
	for _, item := range i {
		total += item.Price
	}
	return total
}
