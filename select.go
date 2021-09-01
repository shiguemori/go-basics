/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
	"time"
)

func selectFunction(out1, out2 io.Writer, message1, message2 string) {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- message1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- message2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Fprint(out1, "received "+msg1)
		case msg2 := <-c2:
			fmt.Fprint(out2, "received "+msg2)
		}
	}
}
