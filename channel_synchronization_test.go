/*
 * Copyright (c) 2021.
 */

package main

import (
	"bytes"
	"testing"
)

func Test_channelSynchronization(t *testing.T) {
	tests := []struct {
		name       string
		wantWriter string
	}{
		{
			name:       "test 1",
			wantWriter: "working...done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			channelSynchronization(writer)
			compareStringsFunction(writer.String(), tt.wantWriter, "channelSynchronization", t)
		})
	}
}

func Test_worker(t *testing.T) {

	channel := make(chan bool, 1)
	channel2 := make(chan bool, 1)
	tests := []struct {
		name       string
		channel    chan bool
		wantWriter string
	}{
		{
			name:       "test 1",
			channel:    channel,
			wantWriter: "working...done",
		},
		{
			name:       "test 2",
			channel:    channel2,
			wantWriter: "working...done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			worker(tt.channel, writer)
			compareStringsFunction(writer.String(), tt.wantWriter, "worker", t)
		})
	}
}
