package pointerhelpers

// Int8Helper contains all Int8 related
// pointer helpers
type Int8Helper struct {
}

// Int8 returns a pointer to the int8 value passed in.
func Int8(v int8) *int8 {
	return &v
}

// Int8 returns a pointer to the int8 value passed in.
func (i *Int8Helper) Int8(v int8) *int8 {
	return Int8(v)
}

// Int8Value returns the value of the int8 pointer passed in or
// 0 if the pointer is nil.
func Int8Value(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int8Value returns the value of the int8 pointer passed in or
// 0 if the pointer is nil.
func (i *Int8Helper) Int8Value(v *int8) int8 {
	return Int8Value(v)
}
