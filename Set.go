// Package Set implements set and operations on set.
package Set

import (
	"errors"
	"reflect"
)

type Set struct {
	set     map[interface{}]bool
	setType reflect.Type
}

// creates a Set for a slice of any type
func NewSet(keys interface{}) *Set {
	s := new(Set)
	s.set = make(map[interface{}]bool)
	s.setType = reflect.TypeOf(keys)
	s.Append(keys)
	return s
}

// returns the length of the set.
func (s *Set) Len() int {
	return len(s.set)
}

// Appends a slice of same type to the set.
func (s *Set) Append(keys interface{}) {
	slice := reflect.ValueOf(keys)
	for i := 0; i < slice.Len(); i++ {
		s.set[slice.Index(i).Interface()] = true
	}
}

// returns the type of set.
func (s *Set) GetType() reflect.Type {
	return s.setType
}

// updates the set with one or more sets.
func (s *Set) Update(sets ...*Set) {
	for _, s_obj := range sets {
		for key, _ := range s_obj.set {
			s.set[key] = true
		}
	}
}

// Adds one or more element to the set.
func (s *Set) Add(keys ...interface{}) {
	for _, key := range keys {
		s.set[key] = true
	}
}

// returns the Set as Slice or slice of interfaces (in case of hetrogeneous set.)
// need to typcast to appropriate type before using.
func (s *Set) Set() interface{} {
	set := reflect.MakeSlice(s.setType, 0, s.Len())
	for key, _ := range s.set {
		set = reflect.Append(set, reflect.ValueOf(key))
	}
	return set.Interface()
}

// Removes one or more element from the set.
func (s *Set) Remove(keys ...interface{}) (err error) {
	for _, key := range keys {
		_, isok := s.set[key]
		if !isok {
			err = errors.New("Error: element not found.")
		}
		delete(s.set, key)
	}
	return
}

// clears the set.
func (s *Set) Clear() {
	for key := range s.set {
		delete(s.set, key)
	}
}

// returns a copy of the set.
func Copy(s *Set) *Set {
	scopy := new(Set)
	scopy.set = make(map[interface{}]bool)
	for key, _ := range s.set {
		scopy.set[key] = true
	}
	scopy.setType = s.setType
	return scopy
}

//returns the union of two or more sets.
func Union(s1 *Set, s2sets ...*Set) *Set {
	s := Copy(s1)
	for _, s2 := range s2sets {
		for key, _ := range s2.set {
			_, isok := s.set[key]
			if !isok {
				s.set[key] = true
			}
		}
	}
	return s
}

// returns the Intersection of two or more sets.
func Intersection(s1 *Set, s2sets ...*Set) *Set {
	s := new(Set)
	s.set = make(map[interface{}]bool)
	count := make(map[interface{}]int)
	for _, s2 := range s2sets {
		for key, _ := range s2.set {
			_, isok := s1.set[key]
			if isok {
				count[key] = count[key] + 1
			}
		}
	}
	l := len(s2sets)
	for key, _ := range s1.set {
		if l == count[key] {
			s.set[key] = true
		}
	}
	s.setType = s1.setType
	return s
}

// Returns the difference of two or more sets as a new set.
// (i.e. all elements that are in this set but not the others.)
func Difference(s1 *Set, s2sets ...*Set) *Set {
	s := Copy(s1)
	for _, s2 := range s2sets {
		for key, _ := range s2.set {
			_, isok := s.set[key]
			if isok {
				delete(s.set, key)
			}
		}
	}
	return s
}
