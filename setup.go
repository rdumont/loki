package loki

import "sync"

type Setup[TIn Matcher[TIn], TOut any] struct {
	mu sync.Mutex

	matchAnyInput bool
	matchInput    TIn
	maxCalls      int
	callCount     int

	out TOut
	fn  func(TIn) TOut
}

func (s *Setup[TIn, TOut]) Once() *Setup[TIn, TOut] {
	return s.NTimes(1)
}

func (s *Setup[TIn, TOut]) NTimes(n int) *Setup[TIn, TOut] {
	s.maxCalls = n
	return s
}

func (s *Setup[TIn, TOut]) Return(out TOut) *Setup[TIn, TOut] {
	s.out = out
	return s
}

func (s *Setup[TIn, TOut]) matches(in TIn) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If max calls was reached
	if s.maxCalls > 0 && s.callCount >= s.maxCalls {
		return false
	}

	if s.matchAnyInput {
		return true
	}

	return s.matchInput.Matches(in)
}

func (s *Setup[TIn, TOut]) call() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.callCount++
}

func (s *Setup[TIn, TOut]) assertStrict(t TestReporter) {
	s.mu.Lock()
	defer s.mu.Unlock()

	inputString := " any arguments"
	if !s.matchAnyInput {
		inputString = ": " + formatValues(s.matchInput.Values())
	}

	if s.maxCalls > 0 && s.callCount < s.maxCalls {
		t.Errorf("Expected %d, but got %d calls with%s", s.maxCalls, s.callCount, inputString)
		return
	}

	if s.maxCalls == 0 && s.callCount == 0 {
		t.Errorf("Expected some, but got %d calls with%s", s.callCount, inputString)
	}
}
