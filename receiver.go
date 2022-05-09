package loki

func Receive[TIn Matcher[TIn], TOut any](m *Method[TIn, TOut], in TIn) TOut {
	for _, s := range m.setups {
		if s.matches(in) {
			s.call()
			out := s.out
			m.calls = append(m.calls, call[TIn, TOut]{In: in, Out: out, Expected: true})
			return out
		}
	}

	var x TOut
	m.calls = append(m.calls, call[TIn, TOut]{In: in, Expected: false})
	return x
}
