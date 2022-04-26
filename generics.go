package main

import (
	"fmt"
	"io"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func Generics(w io.Writer) {
	var m1 = map[int]string{1: "2", 2: "4", 4: "8"}
	var m2 = map[string]int{"1": 1, "2": 4, "3": 8}
	fmt.Fprintln(w, "keys m1:", MapKeys(m1))
	fmt.Fprintln(w, "keys m2:", MapKeys(m2))

	lstInt := List[int]{}
	lstInt.Push(10)
	lstInt.Push(13)
	lstInt.Push(23)

	lstString := List[string]{}
	lstString.Push("Hello")
	lstString.Push("World")
	fmt.Fprintln(w, "int list:", lstInt.GetAll())
	fmt.Fprintln(w, "string list:", lstString.GetAll())
}
