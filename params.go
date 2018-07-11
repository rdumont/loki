package loki

import "fmt"

type Params []interface{}

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

func (p Params) GetOr(i int, or interface{}) interface{} {
	if len(p) > i {
		return p[i]
	}

	return or
}
