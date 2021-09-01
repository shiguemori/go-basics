/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
)

func channels(out io.Writer, args string) {

	messages := make(chan string)

	go func() {
		messages <- args
	}()

	msg := <-messages
	fmt.Fprintln(out, msg)
}
