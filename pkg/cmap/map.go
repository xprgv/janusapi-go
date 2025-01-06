package cmap

import "sync"

type ConcurrentMap[K comparable, V any] struct {
	items map[K]V
	mtx   sync.RWMutex
}

func NewConcurrentMap[K comparable, V any]() ConcurrentMap[K, V] {
	return ConcurrentMap[K, V]{items: make(map[K]V)}
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	v, ok := m.items[key]
	return v, ok
}

func (m *ConcurrentMap[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	keys := make([]K, 0, len(m.items))

	for key := range m.items {
		keys = append(keys, key)
	}

	return keys
}

func (m *ConcurrentMap[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	values := make([]V, 0, len(m.items))

	for _, value := range m.items {
		values = append(values, value)
	}

	return values
}

func (m *ConcurrentMap[K, V]) Items() map[K]V {
	temp := make(map[K]V)

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for key, value := range m.items {
		temp[key] = value
	}

	return temp
}

func (m *ConcurrentMap[K, V]) Exists(key K) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	_, ok := m.items[key]
	return ok
}

func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.items[key] = value
}

func (m *ConcurrentMap[K, V]) MSet(data map[K]V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for key, value := range data {
		m.items[key] = value
	}
}

func (m *ConcurrentMap[K, V]) Count() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	return len(m.items)
}

func (m *ConcurrentMap[K, V]) IsEmpty() bool { return m.Count() == 0 }

type IterFunc[K comparable, V any] func(key K, value V)

func (m *ConcurrentMap[K, V]) Iter(fn IterFunc[K, V]) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for key, value := range m.items {
		fn(key, value)
	}
}

func (m *ConcurrentMap[K, V]) Pop(key K) (V, bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	value, ok := m.items[key]
	delete(m.items, key)
	return value, ok
}

func (m *ConcurrentMap[K, V]) Del(key K) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	_, ok := m.items[key]
	delete(m.items, key)
	return ok
}

func (m *ConcurrentMap[K, V]) MDel(keys []K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for _, key := range keys {
		delete(m.items, key)
	}
}

func (m *ConcurrentMap[K, V]) Clear() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.items = make(map[K]V)
}
