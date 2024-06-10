package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt8(t *testing.T) {
	h := Int8Helper{}
	want := int8(42)
	got := h.Int8(want)
	assert.EqualValues(t, want, *got)
}

func TestInt8Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Int8Helper{}
		got := h.Int8Value(nil)
		assert.Equal(t, int8(0), got)
	})
	t.Run("when set returns the given int8", func(t *testing.T) {
		h := Int8Helper{}
		want := h.Int8(int8(42))
		got := h.Int8Value(want)
		assert.Equal(t, *want, got)
	})
}
