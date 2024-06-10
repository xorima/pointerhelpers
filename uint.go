package pointerhelpers

// UintHelper contains all Uint related
// pointer helpers
type UintHelper struct {
}

// Uint returns a pointer to the uint value passed in.
func Uint(v uint) *uint {
	return &v
}

// Uint returns a pointer to the uint value passed in.
func (i *UintHelper) Uint(v uint) *uint {
	return Uint(v)
}

// UintValue returns the value of the uint pointer passed in or
// 0 if the pointer is nil.
func UintValue(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintValue returns the value of the uint pointer passed in or
// 0 if the pointer is nil.
func (i *UintHelper) UintValue(v *uint) uint {
	return UintValue(v)
}
