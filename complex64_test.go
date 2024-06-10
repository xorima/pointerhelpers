package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplex64(t *testing.T) {
	h := Complex64Helper{}
	want := complex64(42)
	got := h.Complex64(want)
	assert.EqualValues(t, want, *got)
}

func TestComplex64Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Complex64Helper{}
		got := h.Complex64Value(nil)
		assert.Equal(t, complex64(0), got)
	})
	t.Run("when set returns the given complex64", func(t *testing.T) {
		h := Complex64Helper{}
		want := h.Complex64(complex64(42))
		got := h.Complex64Value(want)
		assert.Equal(t, *want, got)
	})
}
