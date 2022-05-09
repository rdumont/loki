package loki

import (
	"fmt"
	"strings"
)

type Matcher[T any] interface {
	Matches(T) bool
	Values() []interface{}
}

type call[TIn, TOut any] struct {
	In       TIn
	Out      TOut
	Expected bool
}

type TestReporter interface {
	Helper()
	Cleanup(func())
	Errorf(format string, args ...any)
}

type Method[TIn Matcher[TIn], TOut any] struct {
	setups []*Setup[TIn, TOut]
	calls  []call[TIn, TOut]
	strict bool
	t      TestReporter
}

func (m *Method[TIn, TOut]) Strict(t TestReporter) *Method[TIn, TOut] {
	t.Helper()

	if m.strict == true {
		return m
	}

	m.strict = true
	m.t = t
	t.Cleanup(m.assert)

	return m
}

func (m *Method[TIn, TOut]) On(in TIn) *Setup[TIn, TOut] {
	s := &Setup[TIn, TOut]{matchInput: in}
	m.setups = append(m.setups, s)
	return s
}

func (m *Method[TIn, TOut]) OnAny() *Setup[TIn, TOut] {
	s := &Setup[TIn, TOut]{matchAnyInput: true}
	m.setups = append(m.setups, s)
	return s
}

func (m *Method[TIn, TOut]) Calls() []TIn {
	var ins []TIn
	for _, c := range m.calls {
		ins = append(ins, c.In)
	}
	return ins
}

func (m *Method[TIn, TOut]) Reset() {
	m.calls = nil
	m.setups = nil
}

func (m *Method[TIn, TOut]) assert() {
	if !m.strict || m.t == nil {
		return
	}

	m.t.Helper()

	var unexpectedCalls []string
	for _, c := range m.calls {
		if c.Expected {
			continue
		}

		unexpectedCalls = append(unexpectedCalls, "\t- "+formatValues(c.In.Values()))
	}

	if len(unexpectedCalls) > 0 {
		m.t.Errorf("Received %d unexpected calls:\n%s\n", len(unexpectedCalls), strings.Join(unexpectedCalls, "\n"))
	}

	for _, s := range m.setups {
		s.assertStrict(m.t)
	}
}

func formatValues(in []interface{}) string {
	var args []string
	for _, v := range in {
		args = append(args, fmt.Sprintf("%v (%T)", v, v))
	}
	return strings.Join(args, "; ")
}
