/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
	"time"
)

func timeout(out1, out2 io.Writer) {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Fprint(out1, res)
	case <-time.After(1 * time.Second):
		fmt.Fprint(out2, "timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Fprint(out1, res)
	case <-time.After(3 * time.Second):
		fmt.Fprint(out2, "timeout 2")
	}
}
