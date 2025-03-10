package goutils

import (
	"reflect"
	"sync"
)

type Set interface {
	Add(element interface{}) bool
	Remove(element interface{})
	Contains(element interface{}) bool
	Size() int
	Clear()
	Intersection(other Set) Set // 交集
	Union(other Set) Set        // 并集
	Difference(other Set) Set   // 差集
	ToSlice() []interface{}
}

type GenericSet struct {
	sync.RWMutex
	setType  reflect.Type
	elements map[interface{}]struct{}
}

func NewSet() *GenericSet {
	return &GenericSet{
		elements: make(map[interface{}]struct{}),
		setType:  nil,
	}
}

func (s *GenericSet) Add(element interface{}) bool {
	s.Lock()
	defer s.Unlock()
	elementType := reflect.TypeOf(element)
	if len(s.elements) == 0 {
		if !elementType.Comparable() {
			return false
		}
		s.elements[element] = struct{}{}
		s.setType = reflect.TypeOf(element)
		return true
	} else if elementType != s.setType {
		return false
	}
	if _, exists := s.elements[element]; !exists {
		s.elements[element] = struct{}{}
	}
	return true
}

func (s *GenericSet) Remove(element interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.elements, element)
}

func (s *GenericSet) Contains(element interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.elements[element]
	return ok
}
func (s *GenericSet) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.elements)
}

func (s *GenericSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.elements = make(map[interface{}]struct{})
}

func (s *GenericSet) ToSlice() []interface{} {
	s.RLock()
	defer s.RUnlock()
	slice := make([]interface{}, 0, len(s.elements))
	for element := range s.elements {
		slice = append(slice, element)
	}
	return slice
}

func (s *GenericSet) Intersection(other Set) Set {
	s.RLock()
	defer s.RUnlock()
	result := NewSet()
	for element := range s.elements {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func (s *GenericSet) Union(other Set) Set {
	s.RLock()
	defer s.RUnlock()
	result := NewSet()
	for element := range s.elements {
		result.Add(element)
	}
	for element := range other.(*GenericSet).elements {
		result.Add(element)
	}
	return result
}

func (s *GenericSet) Difference(other Set) Set {
	s.RLock()
	defer s.RUnlock()
	result := NewSet()
	for element := range s.elements {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
