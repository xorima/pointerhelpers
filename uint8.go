package pointerhelpers

// Uint8Helper contains all Uint8 related
// pointer helpers
type Uint8Helper struct {
}

// Uint8 returns a pointer to the uint8 value passed in.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint8 returns a pointer to the uint8 value passed in.
func (i *Uint8Helper) Uint8(v uint8) *uint8 {
	return Uint8(v)
}

// Uint8Value returns the value of the uint8 pointer passed in or
// 0 if the pointer is nil.
func Uint8Value(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8Value returns the value of the uint8 pointer passed in or
// 0 if the pointer is nil.
func (i *Uint8Helper) Uint8Value(v *uint8) uint8 {
	return Uint8Value(v)
}
