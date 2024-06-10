package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	h := StringHelper{}
	want := "foo"
	got := h.String(want)
	assert.EqualValues(t, want, *got)
}

func TestStringValue(t *testing.T) {
	t.Run("when nil returns empty string", func(t *testing.T) {
		h := StringHelper{}
		want := ""
		got := h.StringValue(nil)
		assert.Equal(t, want, got)
	})
	t.Run("when set returns the given string", func(t *testing.T) {
		h := StringHelper{}
		want := "foobar"
		got := h.StringValue(String(want))
		assert.Equal(t, want, got)
	})
}
