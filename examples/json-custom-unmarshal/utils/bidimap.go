package utils

type BidiMap[K, V comparable] struct {
	data map[any]*BidiMapItem[K, V]
}

type BidiMapItem[K, V comparable] struct {
	Key   K
	Value V
}

func BuildBidiMap[K, V comparable](items []BidiMapItem[K, V]) BidiMap[K, V] {
	m := BidiMap[K, V]{
		data: make(map[any]*BidiMapItem[K, V], len(items)),
	}
	for _, item := range items {
		m.data[item.Key] = &item
		m.data[item.Value] = &item
	}
	return m
}

func (m BidiMap[K, V]) GetKey(value V) (key K, ok bool) {
	item, ok := m.data[value]
	if !ok {
		return key, false
	}
	return item.Key, true
}

func (m BidiMap[K, V]) GetValue(key K) (value V, ok bool) {
	item, ok := m.data[key]
	if !ok {
		return value, false
	}
	return item.Value, true
}
