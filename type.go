package ucode

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	sizePhrase = 6
	bufCap     = 10000
)

type Type struct {
	buf     [bufCap]int
	idx     uint
	low, hi int
	once    sync.Once

	sync.RWMutex
}

func (c *Type) randBuf() {
	var r func(compare int) int
	var buf [bufCap]int

	r = func(compare int) int {
		x := c.low + rand.Intn(c.hi-c.low)
		if x == compare {
			return r(compare)
		}

		return x
	}

	for i := 0; i < bufCap; i++ {
		if i > 0 {
			buf[i] = r(buf[i-1])

			continue
		}

		buf[i] = r(0)
	}

	c.buf = buf
}

func (c *Type) get() int {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	// init buf
	c.once.Do(func() {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		c.low, c.hi = int(math.Pow10(sizePhrase-1)), int(math.Pow10(sizePhrase))-1
		c.randBuf()
	})

	if c.idx == bufCap {
		c.randBuf()
		c.idx = 0
	}

	el := c.buf[c.idx]
	c.idx++

	return el
}

func (c *Type) GetStr() string {
	return fmt.Sprintf("%d", c.get())
}

func (c *Type) Get() int {
	return c.get()
}
