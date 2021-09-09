/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
)

func ClosingChannels(w io.Writer) {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Fprintln(w, fmt.Sprintf("received job %v", j))
			} else {
				fmt.Fprintln(w, "received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Fprintln(w, "sent job")
	}
	close(jobs)
	fmt.Fprintln(w, "sent all jobs")

	d := <-done
	fmt.Fprintln(w, d)
}
