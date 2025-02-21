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
