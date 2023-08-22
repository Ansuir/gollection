package gollection

type Collection[T interface{}] struct {
	items []T
}

func NewCollection[T any](items []T) *Collection[T] {
	return &Collection[T]{items: items}
}

func (c *Collection[T]) Add(item T) *Collection[T] {
	c.items = append(c.items, item)
	return c
}

func (c *Collection[T]) Each(action func(T)) *Collection[T] {
	for _, item := range c.items {
		action(item)
	}
	return c
}

func (c *Collection[T]) Map(f func(T) any) interface{} {
	result := NewCollection[any]([]any{})
	for _, item := range c.items {
		result.Add(f(item))
	}
	return result
}

func Map[T any, TT any](c *Collection[T], f func(T) TT) *Collection[TT] {
	result := make([]TT, len(c.items))
	for i, item := range c.items {
		result[i] = f(item)
	}
	return &Collection[TT]{result}
}
