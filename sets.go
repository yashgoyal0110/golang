package main

import (
	"fmt"
)

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[string]struct{})}
}

func (s *Set) Add(item string) {
	s.data[item] = struct{}{}
}

func (s *Set) Remove(item string) {
	delete(s.data, item)
}

func (s *Set) Contains(item string) bool {
	_, exists := s.data[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Elements() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func caller() {
	set := NewSet()
	set.Add("apple")
	set.Add("banana")
	set.Add("cherry")
	fmt.Println("Set contains 'banana':", set.Contains("banana"))
	set.Remove("banana")
	fmt.Println("Set contains 'banana':", set.Contains("banana"))
	fmt.Println("Set size:", set.Size())
	fmt.Println("Set elements:", set.Elements())
}
