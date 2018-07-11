package loki

import "fmt"

// Params is a slice values either used as input parameters for a method or output values
type Params []interface{}

// Get retrieves the value in position `i`, or panics if there isn't one
func (p Params) Get(i int) interface{} {
	if len(p) > i {
		return p[i]
	}

	var act interface{}
	if p == nil {
		act = nil
	} else {
		act = len(p)
	}
	panic(fmt.Sprintf("expected at least %v values, but got %v", i+1, act))
}

// GetOr retrieves the value in position `i`, or the alternative provided
func (p Params) GetOr(i int, or interface{}) interface{} {
	if len(p) > i {
		return p[i]
	}

	return or
}
