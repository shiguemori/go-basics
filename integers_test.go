/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	assertCorrectSum := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("basic sum of 1,2,3", func(t *testing.T) {
		arrayIntegers := []int{1, 2, 3}
		got := Add(arrayIntegers)
		want := 6
		assertCorrectSum(t, got, want)
	})

	t.Run("testing empty array", func(t *testing.T) {
		arrayIntegers := []int{}
		got := Add(arrayIntegers)
		want := 0
		assertCorrectSum(t, got, want)
	})

	t.Run("testing negative numbers ", func(t *testing.T) {
		arrayIntegers := []int{-2, -3, -5}
		got := Add(arrayIntegers)
		want := -10
		assertCorrectSum(t, got, want)
	})
}

func TestSub(t *testing.T) {
	assertCorrectSub := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("basic subtraction of 1,2,3", func(t *testing.T) {
		arrayIntegers := []int{3, 2, 1}
		got := Sub(arrayIntegers)
		want := -6
		assertCorrectSub(t, got, want)
	})

	t.Run("testing empty array", func(t *testing.T) {
		arrayIntegers := []int{}
		got := Sub(arrayIntegers)
		want := 0
		assertCorrectSub(t, got, want)
	})

	t.Run("testing negative numbers ", func(t *testing.T) {
		arrayIntegers := []int{-2, -3, -5}
		got := Sub(arrayIntegers)
		want := 10
		fmt.Println(got)
		assertCorrectSub(t, got, want)
	})
}
