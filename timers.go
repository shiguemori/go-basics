/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
	"time"
)

func Timers(w io.Writer) {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Fprintf(w, "Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Fprintf(w, "Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Fprintf(w, "Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
