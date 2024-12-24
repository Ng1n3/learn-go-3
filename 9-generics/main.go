package main

import (
	"cmp"
	"fmt"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, true
}

func (s Stacks[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

type Stacks[T comparable] struct {
	vals []T
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	res := make([]T2, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

type OrderableFUnc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFUnc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFUnc[T]) *Tree[T] {
	return &Tree[T]{f: f}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFUnc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	if f(v, n.val) < 0 {
		n.left = n.left.Add(f, v)
	} else {
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFUnc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

func main() {
	//introduction to generics
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))
	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Emmanue", 20})
	t2.Add(Person{"Tom", 30})
	t2.Add(Person{"Sandra", 22})
	fmt.Println(t2.Contains(Person{"Tom", 20}))
	fmt.Println(t2.Contains(Person{"Fred", 30}))
	// fmt.Println(s.Contains(20))
}
