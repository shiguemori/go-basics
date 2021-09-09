/*
 * Copyright (c) 2021.
 */

package main

import (
	"fmt"
	"io"
)

func NonBlockingChannel(w io.Writer) {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
		fmt.Fprintln(w, "received message")
	default:
		fmt.Fprintln(w, "no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println(msg)
		fmt.Fprintln(w, "sent message")
	default:
		fmt.Fprintln(w, "no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println(msg)
		fmt.Fprintln(w, "received message")
	case sig := <-signals:
		fmt.Println(sig)
		fmt.Fprintln(w, "received signal")
	default:
		fmt.Fprintln(w, "no activity")
	}
}
