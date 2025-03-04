// Code generated by go generate
// This file was generated by robots

package optional

import (
	"encoding/json"
	"errors"
)

// Uint64 is an optional uint64.
type Uint64 struct {
	value *uint64
}

// NewUint64 creates an optional.Uint64 from a uint64.
func NewUint64(v uint64) Uint64 {
	return Uint64{&v}
}

// Set sets the uint64 value.
func (u *Uint64) Set(v uint64) {
	u.value = &v
}

// Get returns the uint64 value or an error if not present.
func (u Uint64) Get() (uint64, error) {
	if !u.Present() {
		var zero uint64
		return zero, errors.New("value not present")
	}
	return *u.value, nil
}

// MustGet returns the uint64 value or panics if not present.
func (u Uint64) MustGet() uint64 {
	if !u.Present() {
		panic("value not present")
	}
	return *u.value
}

// Present returns whether or not the value is present.
func (u Uint64) Present() bool {
	return u.value != nil
}

// OrElse returns the uint64 value or a default value if the value is not present.
func (u Uint64) OrElse(v uint64) uint64 {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint64) If(fn func(uint64)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	return json.Marshal(nil)
}

func (u *Uint64) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		u.value = nil
		return nil
	}

	var value uint64

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}
