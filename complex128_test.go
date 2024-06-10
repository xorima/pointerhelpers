package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplex128(t *testing.T) {
	h := Complex128Helper{}
	want := complex128(42)
	got := h.Complex128(want)
	assert.EqualValues(t, want, *got)
}

func TestComplex128Value(t *testing.T) {
	t.Run("when nil returns 0", func(t *testing.T) {
		h := Complex128Helper{}
		got := h.Complex128Value(nil)
		assert.Equal(t, complex128(0), got)
	})
	t.Run("when set returns the given complex128", func(t *testing.T) {
		h := Complex128Helper{}
		want := h.Complex128(complex128(42))
		got := h.Complex128Value(want)
		assert.Equal(t, *want, got)
	})
}
