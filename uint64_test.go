package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUint64(t *testing.T) {
	h := Uint64Helper{}
	want := uint64(42)
	got := h.Uint64(want)
	assert.EqualValues(t, want, *got)
}

func TestUint64Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Uint64Helper{}
		got := h.Uint64Value(nil)
		assert.Equal(t, uint64(0), got)
	})
	t.Run("when set returns the given uint64", func(t *testing.T) {
		h := Uint64Helper{}
		want := h.Uint64(uint64(42))
		got := h.Uint64Value(want)
		assert.Equal(t, *want, got)
	})
}
