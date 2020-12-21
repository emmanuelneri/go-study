package main

import (
	"testing"
)

func Benchmark_closeEndFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := Closable{number: i, status: "OPEN"}
		c.close()
	}
}

func Benchmark_closeWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := Closable{number: i}
		defer c.close()
	}
}

type Closable struct {
	number int
	status string
}

func (c *Closable) close() {
	c.status = "CLOSE"
}
