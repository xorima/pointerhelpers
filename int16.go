package pointerhelpers

// Int16Helper contains all Int16 related
// pointer helpers
type Int16Helper struct {
}

// Int16 returns a pointer to the int16 value passed in.
func Int16(v int16) *int16 {
	return &v
}

// Int16 returns a pointer to the int16 value passed in.
func (i *Int16Helper) Int16(v int16) *int16 {
	return Int16(v)
}

// Int16Value returns the value of the int16 pointer passed in or
// 0 if the pointer is nil.
func Int16Value(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16Value returns the value of the int16 pointer passed in or
// 0 if the pointer is nil.
func (i *Int16Helper) Int16Value(v *int16) int16 {
	return Int16Value(v)
}
