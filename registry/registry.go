package registry

import (
	"errors"
)

// Registry error.
var (
	ErrKeyLocked = errors.New("key is locked")
)

// Registry is a simple interface allowing untyped common values to be shared across packages.
// This is particularly useful for resolving import cycles.
type Registry struct {
	data  map[string]any
	locks map[string]bool
}

// Get the value of a key.
// This function returns nil if no value exists.
func (reg *Registry) Get(key string) any {
	return reg.data[key]
}

// Has checks whether a key is non-nil.
func (reg *Registry) Has(key string) bool {
	return reg.data[key] == nil
}

// Lock a key.
// This prevents the value being overwritten by Set.
func (reg *Registry) Lock(key string) {
	reg.locks[key] = true
}

// Set the value for a key.
// This function returns an error if the key is locked.
func (reg *Registry) Set(key string, value any) error {
	if reg.locks[key] {
		return ErrKeyLocked
	}
	reg.data[key] = value
	return nil
}

// New creates a new registry.
func New() *Registry {
	return &Registry{
		data:  map[string]any{},
		locks: map[string]bool{},
	}
}
