// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2017-07-23 21:16:16.872603135 +0000 UTC

package optional

// Bool is an optional bool
type Bool struct {
	bool    bool
	present bool
}

// EmptyBool returns an empty optional.Bool
func EmptyBool() Bool {
	return Bool{}
}

// OfBool creates an optional.Bool from a bool
func OfBool(b bool) Bool {
	return Bool{bool: b, present: true}
}

// Set sets the bool value
func (o *Bool) Set(b bool) {
	o.bool = b
	o.present = true
}

// Bool returns the bool value
func (o *Bool) Bool() bool {
	return o.bool
}

// Present returns whether or not the value is present
func (o *Bool) Present() bool {
	return o.present
}

// OrElse returns the bool value or a default value if the value is not present
func (o *Bool) OrElse(b bool) bool {
	if o.present {
		return o.bool
	}
	return b
}
