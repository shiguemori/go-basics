/*
 * Copyright (c) 2021.
 */

package main

func Add(integers []int) (sum int) {
	for _, v := range integers {
		sum += v
	}
	return sum
}

func Sub(integers []int) (sub int) {
	for _, v := range integers {
		sub -= v
	}
	return sub
}
