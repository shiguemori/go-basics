/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMyGreeterHandler(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		addr string
	}{
		{
			name: "test 1",
			addr: ":5050",
		},
		{
			name: "test 1",
			addr: ":5051",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := http.ListenAndServe(tt.addr, http.HandlerFunc(MyGreeterHandler))
			if err != nil {
				t.Errorf("error %q", err)
				return
			}
		})
	}
}
