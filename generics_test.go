/*
 * Copyright (c) 2022.
 */

package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestGenerics(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Generics",
			want: "keys m1: [1 2 4]keys m2: [1 2 3]int list: [10 13 23]string list: [Hello World]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Generics(w)
			got := strings.Replace(w.String(), "\n", "", -1)
			compareStringsFunction(got, tt.want, "Generics", t)
		})
	}
}

func TestList_GetAll(t *testing.T) {

	tests := []struct {
		name string
		want []any
		list []any
	}{
		{
			name: "Test List int",
			list: []any{10, 13, 23},
			want: []any{10, 13, 23},
		},
		{
			name: "Test List String",
			list: []any{"1", "3", "5"},
			want: []any{"1", "3", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := List[any]{}
			for _, v := range tt.list {
				lst.Push(v)
			}
			if got := lst.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
