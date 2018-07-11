package loki

// ParamMatcher is a function that dictates whether an actual parameter matches the expected one
type ParamMatcher func(MethodMetadata, interface{}) bool

// MethodMetadata keeps track of how many times a method has been called, in case it matters to the matcher
type MethodMetadata struct {
	CallCount int
}

// Anything is a `ParamMatcher` that will always match, no matter the value
var Anything ParamMatcher = func(MethodMetadata, interface{}) bool {
	return true
}

// NthCall is a `ParamMatcher` that will match when the method has been called `n` times
func NthCall(n int) ParamMatcher {
	return func(meta MethodMetadata, value interface{}) bool {
		return meta.CallCount == n
	}
}
