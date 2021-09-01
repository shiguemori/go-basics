/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func routineLoop(out io.Writer, from string) {
	for i := 0; i < 3; i++ {
		fmt.Fprintln(out, from+":"+strconv.Itoa(i))
	}
}

func runGoRoutine(out io.Writer) {

	routineLoop(out, "direct")

	go routineLoop(out, "goroutine")

	go func(msg string) {
		fmt.Fprintln(out, msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Fprintln(out, "done")
}
