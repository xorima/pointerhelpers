package pointerhelpers

// StringHelper contains all String related
// pointer helpers
type StringHelper struct{}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// String returns a pointer to the string value passed in.
func (s *StringHelper) String(v string) *string {
	return String(v)
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.

func (s *StringHelper) StringValue(v *string) string {
	return StringValue(v)
}
