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
