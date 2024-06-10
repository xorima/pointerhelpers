package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt32(t *testing.T) {
	h := Int32Helper{}
	want := int32(42)
	got := h.Int32(want)
	assert.EqualValues(t, want, *got)
}

func TestInt32Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Int32Helper{}
		got := h.Int32Value(nil)
		assert.Equal(t, int32(0), got)
	})
	t.Run("when set returns the given int32", func(t *testing.T) {
		h := Int32Helper{}
		want := h.Int32(int32(42))
		got := h.Int32Value(want)
		assert.Equal(t, *want, got)
	})
}
