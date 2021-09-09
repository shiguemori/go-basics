/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
)

func RangeOverChannels(w io.Writer) {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Fprintln(w, elem)
	}
}
