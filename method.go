package loki

import "sync"

// Method is the basic structure used to fake method calls
//
// Example:
// type FakeTodoList struct {
//     AddCalls		loki.Method
// }
type Method struct {
	sync.Mutex
	calls        []Params
	expectations []*ExpectedCall
}

// CallCount retrieves how many times the method has been called
func (m *Method) CallCount() int {
	return len(m.calls)
}

// GetCall returns the parameters of the first call made to the method, or nil
func (m *Method) GetCall() Params {
	if len(m.calls) == 0 {
		return nil
	}

	return m.calls[0]
}

// GetNthCall returns the parameters of the nth call made to a method, or nil
func (m *Method) GetNthCall(i int) Params {
	if len(m.calls) == 0 {
		return nil
	}

	return m.calls[i]
}

// Receive represents the method receiving a call, and should be used in the method implementation
//
// Example:
// func (tl *FakeTodoList) Add(name string) {
//     tl.AddCalls.Receive(name)
// }
func (m *Method) Receive(a ...interface{}) Params {
	m.Lock()
	defer m.Unlock()

	m.calls = append(m.calls, a)
	for i := len(m.expectations) - 1; i >= 0; i-- {
		c := m.expectations[i]
		if c.matches(a) {
			return c.returns
		}
	}

	return nil
}

// On is used to set up an expected call to the method
//
// Example:
// tl := new(FakeTodoList)
// tl.AddCalls.On("buy milk").Run(func(p) { fmt.Prinln("Called Add with parameters ", p"); })
func (m *Method) On(a ...interface{}) *ExpectedCall {
	c := &ExpectedCall{m, a, nil, nil}
	m.expectations = append(m.expectations, c)
	return c
}
