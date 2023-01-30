package pointerhelpers

// BoolHelper contains all Bool related
// pointer helpers
type BoolHelper struct {
}

// Bool returns a pointer to the Bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// Bool returns a pointer to the bool value passed in.
func (b *BoolHelper) Bool(v bool) *bool {
	return Bool(v)
}

// BoolValue returns the value of the bool pointer passed in or
// 0 if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolValue returns the value of the bool pointer passed in or
// 0 if the pointer is nil.
func (b *BoolHelper) BoolValue(v *bool) bool {
	return BoolValue(v)
}

// BoolIsUnset returns true if the pointer is nil
func BoolIsUnset(v *bool) bool {
	return v == nil
}

// BoolIsUnset returns true if the pointer is nil
func (b *BoolHelper) BoolIsUnset(v *bool) bool {
	return BoolIsUnset(v)
}
