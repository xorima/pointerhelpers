package pointerhelpers

// IntHelper contains all int related
// pointer helpers
type IntHelper struct {
}

// Int returns a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}

// Int returns a pointer to the int value passed in.
func (i *IntHelper) Int(v int) *int {
	return Int(v)
}

// IntValue returns the value of the int pointer passed in or
// 0 if the pointer is nil.
func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntValue returns the value of the int pointer passed in or
// 0 if the pointer is nil.
func (i *IntHelper) IntValue(v *int) int {
	return IntValue(v)
}
