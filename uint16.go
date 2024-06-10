package pointerhelpers

// Uint16Helper contains all Uint16 related
// pointer helpers
type Uint16Helper struct {
}

// Uint16 returns a pointer to the uint16 value passed in.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint16 returns a pointer to the uint16 value passed in.
func (i *Uint16Helper) Uint16(v uint16) *uint16 {
	return Uint16(v)
}

// Uint16Value returns the value of the uint16 pointer passed in or
// 0 if the pointer is nil.
func Uint16Value(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16Value returns the value of the uint16 pointer passed in or
// 0 if the pointer is nil.
func (i *Uint16Helper) Uint16Value(v *uint16) uint16 {
	return Uint16Value(v)
}
