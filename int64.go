package pointerhelpers

// Int64Helper contains all Int64 related
// pointer helpers
type Int64Helper struct {
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64 returns a pointer to the int64 value passed in.
func (i *Int64Helper) Int64(v int64) *int64 {
	return Int64(v)
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func (i *Int64Helper) Int64Value(v *int64) int64 {
	return Int64Value(v)
}
