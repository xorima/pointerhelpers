package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt16(t *testing.T) {
	h := Int16Helper{}
	want := int16(42)
	got := h.Int16(want)
	assert.EqualValues(t, want, *got)
}

func TestInt16Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Int16Helper{}
		got := h.Int16Value(nil)
		assert.Equal(t, int16(0), got)
	})
	t.Run("when set returns the given int16", func(t *testing.T) {
		h := Int16Helper{}
		want := h.Int16(int16(42))
		got := h.Int16Value(want)
		assert.Equal(t, *want, got)
	})
}
