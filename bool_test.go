package pointerhelpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	h := BoolHelper{}
	want := true
	got := h.Bool(want)
	assert.True(t, *got)
}

func TestBoolValue(t *testing.T) {
	t.Run("when nil returns false", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolValue(nil)
		assert.False(t, got)
	})
	t.Run("When true, returns true", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolValue(Bool(true))
		assert.True(t, got)
	})
	t.Run("When false, returns false", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolValue(Bool(false))
		assert.False(t, got)
	})
}

func TestBoolIsUnset(t *testing.T) {
	t.Run("when nil returns true", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolIsUnset(nil)
		assert.True(t, got)
	})
	t.Run("When true, returns false", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolIsUnset(Bool(true))
		assert.False(t, got)
	})
	t.Run("When false, returns false", func(t *testing.T) {
		h := BoolHelper{}
		got := h.BoolIsUnset(Bool(false))
		assert.False(t, got)
	})
}
