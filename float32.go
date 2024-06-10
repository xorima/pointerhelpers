package pointerhelpers

// Float32Helper contains all Float32 related
// pointer helpers
type Float32Helper struct {
}

// Float32 returns a pointer to the float32 value passed in.
func Float32(v float32) *float32 {
	return &v
}

// Float32 returns a pointer to the float32 value passed in.
func (i *Float32Helper) Float32(v float32) *float32 {
	return Float32(v)
}

// Float32Value returns the value of the float32 pointer passed in or
// 0 if the pointer is nil.
func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Value returns the value of the float32 pointer passed in or
// 0 if the pointer is nil.
func (i *Float32Helper) Float32Value(v *float32) float32 {
	return Float32Value(v)
}
