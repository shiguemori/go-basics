/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
	"time"
)

func worker(done chan bool, writer io.Writer) {
	fmt.Fprint(writer, "working...")
	time.Sleep(time.Second)
	fmt.Fprint(writer, "done")
	done <- true
}

func channelSynchronization(writer io.Writer) {
	done := make(chan bool, 1)
	go worker(done, writer)
	<-done
}
