package pointerhelpers

// Complex128Helper contains all Complex128 related
// pointer helpers
type Complex128Helper struct {
}

// Complex128 returns a pointer to the complex128 value passed in.
func Complex128(v complex128) *complex128 {
	return &v
}

// Complex128 returns a pointer to the complex128 value passed in.
func (i *Complex128Helper) Complex128(v complex128) *complex128 {
	return Complex128(v)
}

// Complex128Value returns the value of the complex128 pointer passed in or
// 0 if the pointer is nil.
func Complex128Value(v *complex128) complex128 {
	if v != nil {
		return *v
	}
	return 0
}

// Complex128Value returns the value of the complex128 pointer passed in or
// 0 if the pointer is nil.
func (i *Complex128Helper) Complex128Value(v *complex128) complex128 {
	return Complex128Value(v)
}
