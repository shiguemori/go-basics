/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func Test_channelDirections(t *testing.T) {
	tests := []struct {
		name      string
		wantWrite string
	}{
		{
			name:      "test 1",
			wantWrite: "passed message",
		},
		{
			name:      "test 2",
			wantWrite: "passed message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			channelDirections(out)
			compareStringsFunction(out.String(), tt.wantWrite, "channelDirections", t)
		})
	}
}
