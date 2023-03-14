package ucode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandRange(t *testing.T) {
	t.Run("no duplicates", func(t *testing.T) {
		const elements = 100_000
		var codes Type

		last := codes.Get()

		for i := 0; i < elements; i++ {
			el := codes.Get()

			if last != el {
				last = el

				continue
			}

			assert.Failf(t, "duplicate", "last value %d, current %d", last, el)
		}
	})
}
