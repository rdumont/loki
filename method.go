package loki

import "sync"

type Method struct {
	sync.Mutex
	calls        []Params
	expectations []*ExpectedCall
}

func (m *Method) CallCount() int {
	return len(m.calls)
}

func (m *Method) GetCall() Params {
	if len(m.calls) == 0 {
		return nil
	}

	return m.calls[0]
}

func (m *Method) GetNthCall(i int) Params {
	if len(m.calls) == 0 {
		return nil
	}

	return m.calls[i]
}

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

func (m *Method) On(a ...interface{}) *ExpectedCall {
	c := &ExpectedCall{m, a, nil, nil}
	m.expectations = append(m.expectations, c)
	return c
}
