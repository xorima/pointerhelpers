package pointerhelpers

// Float64Helper contains all Float64 related
// pointer helpers
type Float64Helper struct {
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64 returns a pointer to the float64 value passed in.
func (i *Float64Helper) Float64(v float64) *float64 {
	return Float64(v)
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func (i *Float64Helper) Float64Value(v *float64) float64 {
	return Float64Value(v)
}
