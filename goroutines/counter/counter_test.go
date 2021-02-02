package counter

import (
	"sync"
	"testing"
	"time"
)

type Value struct {
	m sync.Mutex
	i int
}

func (c *Value) add(value int) {
	c.m.Lock()         // require to always pass test
	defer c.m.Unlock() // require to always pass test
	c.i += value
}

func TestCounter_Add(t *testing.T) {
	c := Value{}

	for i := 0; i <= 100; i++ {
		go c.add(i)
	}

	time.Sleep(1 * time.Second)

	if c.i != 5050 {
		t.Error("invalid value: ", c.i)
	}
}
