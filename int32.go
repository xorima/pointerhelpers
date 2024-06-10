package pointerhelpers

// Int32Helper contains all Int32 related
// pointer helpers
type Int32Helper struct {
}

// Int32 returns a pointer to the int32 value passed in.
func Int32(v int32) *int32 {
	return &v
}

// Int32 returns a pointer to the int32 value passed in.
func (i *Int32Helper) Int32(v int32) *int32 {
	return Int32(v)
}

// Int32Value returns the value of the int32 pointer passed in or
// 0 if the pointer is nil.
func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32Value returns the value of the int32 pointer passed in or
// 0 if the pointer is nil.
func (i *Int32Helper) Int32Value(v *int32) int32 {
	return Int32Value(v)
}
