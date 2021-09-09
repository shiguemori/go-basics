/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func TestClosingChannels(t *testing.T) {
	tests := []struct {
		name  string
		wantW string
	}{
		{
			name:  "test 1",
			wantW: "sent job sent job sent job sent all jobs received job 1 received job 2 received job 3 received all jobs true ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			ClosingChannels(out)
			gotOut := bufferToStringReplace(out)
			compareStringsFunction(gotOut, tt.wantW, "ClosingChannels", t)
		})
	}
}
