/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func TestNonBlockingChannel(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test 1",
			want: "no message received no message sent no activity ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			NonBlockingChannel(out)
			gotOut := bufferToStringReplace(out)
			compareStringsFunction(gotOut, tt.want, "NonBlockingChannel", t)
		})
	}
}
