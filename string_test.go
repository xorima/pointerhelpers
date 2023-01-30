package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64(t *testing.T) {
	h := Int64Helper{}
	want := int64(42)
	got := h.Int64(want)
	assert.EqualValues(t, want, *got)
}

func TestInt64Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Int64Helper{}
		got := h.Int64Value(nil)
		assert.Equal(t, int64(0), got)
	})
	t.Run("when set returns the given int64", func(t *testing.T) {
		h := Int64Helper{}
		want := h.Int64(int64(42))
		got := h.Int64Value(want)
		assert.Equal(t, *want, got)
	})
}
