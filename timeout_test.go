/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func Test_timeout(t *testing.T) {
	tests := []struct {
		name     string
		wantOut1 string
		wantOut2 string
	}{
		{
			name:     "test 1",
			wantOut1: "result 2",
			wantOut2: "timeout 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out1 := &bytes.Buffer{}
			out2 := &bytes.Buffer{}
			timeout(out1, out2)
			if gotOut1 := out1.String(); gotOut1 != tt.wantOut1 {
				t.Errorf("timeout() gotOut1 = %v, want %v", gotOut1, tt.wantOut1)
			}
			if gotOut2 := out2.String(); gotOut2 != tt.wantOut2 {
				t.Errorf("timeout() gotOut2 = %v, want %v", gotOut2, tt.wantOut2)
			}
		})
	}
}
