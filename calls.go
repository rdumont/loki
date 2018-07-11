package loki

// ExpectedCall is the definition of how a method call is expected to be made
//
// To set up a call, use `new(loki.Method).On(...)`
type ExpectedCall struct {
	method  *Method
	params  Params
	run     []func(Params)
	returns Params
}

// Run sets up the call to execute a function without returning any values
func (c *ExpectedCall) Run(f func(Params)) *ExpectedCall {
	c.run = append(c.run, f)
	return c
}

// Return sets up a call to return the provided values `a`
func (c *ExpectedCall) Return(a ...interface{}) *ExpectedCall {
	c.returns = a
	return c
}

func (c *ExpectedCall) matches(p Params) bool {
	if len(c.params) != len(p) {
		return false
	}

	for i, ep := range c.params {
		if f, ok := ep.(ParamMatcher); ok {
			meta := MethodMetadata{len(c.method.calls)}
			if !f(meta, p[i]) {
				return false
			}
		} else if ep != p[i] {
			return false
		}
	}

	return true
}
