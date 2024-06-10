package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint(t *testing.T) {
	h := UintHelper{}
	want := uint(42)
	got := h.Uint(want)
	assert.EqualValues(t, want, *got)
}

func TestUintValue(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := UintHelper{}
		got := h.UintValue(nil)
		assert.Equal(t, uint(0), got)
	})
	t.Run("when set returns the given uint", func(t *testing.T) {
		h := UintHelper{}
		want := h.Uint(uint(42))
		got := h.UintValue(want)
		assert.Equal(t, *want, got)
	})
}
