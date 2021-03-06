package main

import "github.com/smartystreets/go-disruptor"

func publish(writer *disruptor.Writer) {
	for {
		sequence := writer.Reserve(1)
		if sequence != disruptor.Gating {
			ringBuffer[sequence&RingMask] = sequence
			writer.Commit(sequence)
		}
	}
}
