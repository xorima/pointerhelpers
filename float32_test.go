package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32(t *testing.T) {
	h := Float32Helper{}
	want := float32(42)
	got := h.Float32(want)
	assert.EqualValues(t, want, *got)
}

func TestFloat32Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Float32Helper{}
		got := h.Float32Value(nil)
		assert.Equal(t, float32(0), got)
	})
	t.Run("when set returns the given float32", func(t *testing.T) {
		h := Float32Helper{}
		want := h.Float32(float32(42))
		got := h.Float32Value(want)
		assert.Equal(t, *want, got)
	})
}
