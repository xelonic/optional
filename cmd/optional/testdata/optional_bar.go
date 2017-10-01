// Code generated by go generate
// This file was generated by robots at 2017-10-01 20:44:25.670012108 +0000 UTC

package bar

import "errors"

// optionalBar is an optional bar
type optionalBar struct {
	value *bar
}

// NewoptionalBar creates a optional.optionalBar from a bar
func NewoptionalBar(v bar) optionalBar {
	return optionalBar{&v}
}

// Set sets the bar value
func (o optionalBar) Set(v bar) {
	o.value = &v
}

// Get returns the bar value or an error if not present
func (o optionalBar) Get() (bar, error) {
	if !o.Present() {
		return *o.value, errors.New("value not present")
	}
	return *o.value, nil
}

// Present returns whether or not the value is present
func (o optionalBar) Present() bool {
	return o.value != nil
}

// OrElse returns the bar value or a default value if the value is not present
func (o optionalBar) OrElse(v bar) bar {
	if o.Present() {
		return *o.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (o optionalBar) If(fn func(bar)) {
	if o.Present() {
		fn(*o.value)
	}
}
