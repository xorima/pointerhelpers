package pointerhelpers

// Uint64Helper contains all Uint64 related
// pointer helpers
type Uint64Helper struct {
}

// Uint64 returns a pointer to the uint64 value passed in.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64 returns a pointer to the uint64 value passed in.
func (i *Uint64Helper) Uint64(v uint64) *uint64 {
	return Uint64(v)
}

// Uint64Value returns the value of the uint64 pointer passed in or
// 0 if the pointer is nil.
func Uint64Value(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64Value returns the value of the uint64 pointer passed in or
// 0 if the pointer is nil.
func (i *Uint64Helper) Uint64Value(v *uint64) uint64 {
	return Uint64Value(v)
}
