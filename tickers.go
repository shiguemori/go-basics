/*
 * Copyright (c) 2021.
 */
package main

import (
	"fmt"
	"io"
	"time"
)

func Tickers(w io.Writer) {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println(fmt.Sprintf("Tick at %v", t))
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Fprintln(w, "Ticker stopped")
}
