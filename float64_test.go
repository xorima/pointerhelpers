package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64(t *testing.T) {
	h := Float64Helper{}
	want := float64(42)
	got := h.Float64(want)
	assert.EqualValues(t, want, *got)
}

func TestFloat64Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Float64Helper{}
		got := h.Float64Value(nil)
		assert.Equal(t, float64(0), got)
	})
	t.Run("when set returns the given float64", func(t *testing.T) {
		h := Float64Helper{}
		want := h.Float64(float64(42))
		got := h.Float64Value(want)
		assert.Equal(t, *want, got)
	})
}
