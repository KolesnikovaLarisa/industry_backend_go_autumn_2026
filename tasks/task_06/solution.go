package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	if c.capacity <= 0 {
		return
	}

	el, ok := c.items[key]
	if !ok {
		return
	}

	item := el.Value.(entry[K, V])
	c.ll.MoveToFront(el)
	return item.value, ok
}

func (c *LRUCache[K, V]) Set(key K, value V) {

	if c.capacity > 0 {
		el, ok := c.items[key]

		if ok {
			// обновляем только значение по существующему ключу
			item := el.Value.(entry[K, V])
			item.value = value
			el.Value = item
			return
		}

		if len(c.items) >= c.capacity {
			// кэш заполнен, удаляем самый старый элемент
			old_el := c.ll.Back()
			old_ent := old_el.Value.(entry[K, V])
			delete(c.items, old_ent.key)
			c.ll.Remove(old_el)
		}

		// добавляем новый элемент в начало списка
		new_el := c.ll.PushFront(entry[K, V]{
			key:   key,
			value: value,
		})

		c.items[key] = new_el
	}
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
