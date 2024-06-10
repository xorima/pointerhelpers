package pointerhelpers

// Uint32Helper contains all Uint32 related
// pointer helpers
type Uint32Helper struct {
}

// Uint32 returns a pointer to the uint32 value passed in.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32 returns a pointer to the uint32 value passed in.
func (i *Uint32Helper) Uint32(v uint32) *uint32 {
	return Uint32(v)
}

// Uint32Value returns the value of the uint32 pointer passed in or
// 0 if the pointer is nil.
func Uint32Value(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32Value returns the value of the uint32 pointer passed in or
// 0 if the pointer is nil.
func (i *Uint32Helper) Uint32Value(v *uint32) uint32 {
	return Uint32Value(v)
}
