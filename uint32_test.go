package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint32(t *testing.T) {
	h := Uint32Helper{}
	want := uint32(42)
	got := h.Uint32(want)
	assert.EqualValues(t, want, *got)
}

func TestUint32Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Uint32Helper{}
		got := h.Uint32Value(nil)
		assert.Equal(t, uint32(0), got)
	})
	t.Run("when set returns the given uint32", func(t *testing.T) {
		h := Uint32Helper{}
		want := h.Uint32(uint32(42))
		got := h.Uint32Value(want)
		assert.Equal(t, *want, got)
	})
}
