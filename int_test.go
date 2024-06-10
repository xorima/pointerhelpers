package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	h := IntHelper{}
	want := 42
	got := h.Int(want)
	assert.EqualValues(t, want, *got)
}

func TestIntValue(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := IntHelper{}
		got := h.IntValue(nil)
		assert.Equal(t, int(0), got)
	})
	t.Run("when set returns the given int", func(t *testing.T) {
		h := IntHelper{}
		want := h.Int(42)
		got := h.IntValue(want)
		assert.Equal(t, *want, got)
	})
}
