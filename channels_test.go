/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_channels(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test 1",
			args: "ping",
			want: "ping",
		},
		{
			name: "test 2",
			args: "pong",
			want: "pong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			channels(out, tt.args)
			gotOut := strings.Replace(out.String(), "\n", "", -1)
			compareStringsFunction(gotOut, tt.want, "channels", t)
		})
	}
}
