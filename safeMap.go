package goutils

import "sync"

// SafeMap string key
type SafeMap struct {
	sync.RWMutex
	M map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		M: make(map[string]interface{}),
	}
}

func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.RLock()
	defer sm.RUnlock()
	val, ok := sm.M[key]
	return val, ok
}

func (sm *SafeMap) Set(key string, val interface{}) {
	sm.Lock()
	defer sm.Unlock()
	sm.M[key] = val
}

func (sm *SafeMap) Delete(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.M, key)
}

func (sm *SafeMap) Len() int {
	sm.RLock()
	defer sm.RUnlock()
	return len(sm.M)
}

// SafeIntMap int key
type SafeIntMap struct {
	sync.RWMutex
	M map[int]interface{}
}

func NewSafeIntMap() *SafeIntMap {
	return &SafeIntMap{
		M: make(map[int]interface{}),
	}
}

func (sm *SafeIntMap) Get(key int) (interface{}, bool) {
	sm.RLock()
	defer sm.RUnlock()
	val, ok := sm.M[key]
	return val, ok
}

func (sm *SafeIntMap) Set(key int, val interface{}) {
	sm.Lock()
	defer sm.Unlock()
	sm.M[key] = val
}

func (sm *SafeIntMap) Delete(key int) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.M, key)
}

func (sm *SafeIntMap) Len() int {
	sm.RLock()
	defer sm.RUnlock()
	return len(sm.M)
}

// SafeMaps 1.18+
type SafeMaps[K comparable, V any] struct {
	sync.RWMutex
	M map[K]V
}

func NewSafeMaps[K comparable, V any]() *SafeMaps[K, V] {
	return &SafeMaps[K, V]{
		M: make(map[K]V),
	}
}

func (sm *SafeMaps[K, V]) Get(key K) (interface{}, bool) {
	sm.RLock()
	defer sm.RUnlock()
	val, ok := sm.M[key]
	return val, ok
}

func (sm *SafeMaps[K, V]) Set(key K, val V) {
	sm.Lock()
	defer sm.Unlock()
	sm.M[key] = val
}

func (sm *SafeMaps[K, V]) Delete(key K) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.M, key)
}

func (sm *SafeMaps[K, V]) Len() int {
	sm.RLock()
	defer sm.RUnlock()
	return len(sm.M)
}

type OrderMap struct {
	*SafeMap
	keys []string
}

type MapItem struct {
	Key   string
	Value interface{}
}

func NewOrderMap() *OrderMap {
	return &OrderMap{
		SafeMap: NewSafeMap(),
		keys:    make([]string, 0),
	}
}

func (om *OrderMap) Get(key string) (interface{}, bool) {
	om.RLock()
	defer om.RUnlock()
	val, ok := om.M[key]
	return val, ok
}

func (om *OrderMap) Set(key string, val interface{}) {
	om.Lock()
	defer om.Unlock()
	if _, ok := om.M[key]; !ok {
		om.keys = append(om.keys, key)
	}
	om.M[key] = val
}

func (om *OrderMap) Delete(key string) {
	om.Lock()
	defer om.Unlock()
	delete(om.M, key)
	for i, k := range om.keys {
		if k == key {
			om.keys = append(om.keys[:i], om.keys[i+1:]...)
			break
		}
	}
}

func (om *OrderMap) Len() int {
	om.RLock()
	defer om.RUnlock()
	return len(om.M)
}

func (om *OrderMap) Range() []MapItem {
	om.RLock()
	defer om.RUnlock()
	result := make([]MapItem, 0, len(om.M))
	for _, k := range om.keys {
		if val, ok := om.M[k]; ok {
			result = append(result, MapItem{Key: k, Value: val})
		}
	}
	return result
}
