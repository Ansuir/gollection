package gollection

type Collection[T any] struct {
	items []T
}

func NewCollection[T any](items []T) Collection[T] {
	return Collection[T]{items: items}
}

func (c *Collection[T]) Add(item T) {
	c.items = append(c.items, item)
}

func (c *Collection[T]) Each(action func(T)) {
	for _, item := range c.items {
		action(item)
	}
}
