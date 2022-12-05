/*
 * Copyright (c) $originalComment.match("Copyright \(c\) (\d+)", 1, "-")2022.
 */

package main

import (
	"fmt"
)

func executeParallel(ch chan<- int, functions ...func() int) {
	for _, f := range functions {
		go func(f func() int) {
			ch <- f()
		}(f)
	}
}

func exampleFunction2(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	expensiveFunction := func() int { return exampleFunction2(200000000) }
	cheapFunction := func() int { return exampleFunction2(10000000) }

	ch := make(chan int)

	go executeParallel(ch, expensiveFunction, cheapFunction)
	for result := range ch {
		fmt.Printf("Result: %d\n", result)
	}
	close(ch)
}

