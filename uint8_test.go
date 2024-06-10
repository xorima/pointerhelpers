package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint8(t *testing.T) {
	h := Uint8Helper{}
	want := uint8(42)
	got := h.Uint8(want)
	assert.EqualValues(t, want, *got)
}

func TestUint8Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Uint8Helper{}
		got := h.Uint8Value(nil)
		assert.Equal(t, uint8(0), got)
	})
	t.Run("when set returns the given uint8", func(t *testing.T) {
		h := Uint8Helper{}
		want := h.Uint8(uint8(42))
		got := h.Uint8Value(want)
		assert.Equal(t, *want, got)
	})
}
