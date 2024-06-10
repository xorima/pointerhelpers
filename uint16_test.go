package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint16(t *testing.T) {
	h := Uint16Helper{}
	want := uint16(42)
	got := h.Uint16(want)
	assert.EqualValues(t, want, *got)
}

func TestUint16Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Uint16Helper{}
		got := h.Uint16Value(nil)
		assert.Equal(t, uint16(0), got)
	})
	t.Run("when set returns the given uint16", func(t *testing.T) {
		h := Uint16Helper{}
		want := h.Uint16(uint16(42))
		got := h.Uint16Value(want)
		assert.Equal(t, *want, got)
	})
}
