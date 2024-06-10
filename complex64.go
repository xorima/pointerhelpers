package pointerhelpers

// Complex64Helper contains all Complex64 related
// pointer helpers
type Complex64Helper struct {
}

// Complex64 returns a pointer to the complex64 value passed in.
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex64 returns a pointer to the complex64 value passed in.
func (i *Complex64Helper) Complex64(v complex64) *complex64 {
	return Complex64(v)
}

// Complex64Value returns the value of the complex64 pointer passed in or
// 0 if the pointer is nil.
func Complex64Value(v *complex64) complex64 {
	if v != nil {
		return *v
	}
	return 0
}

// Complex64Value returns the value of the complex64 pointer passed in or
// 0 if the pointer is nil.
func (i *Complex64Helper) Complex64Value(v *complex64) complex64 {
	return Complex64Value(v)
}
