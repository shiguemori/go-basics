/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
)

func channelBuffering(out io.Writer, args []string) {

	messages := make(chan string, 3)

	for _, v := range args {
		messages <- v
	}

	for i := 0; i < len(args); i++ {
		fmt.Fprintln(out, <-messages)
	}
}
