package loki

type ParamMatcher func(MethodMetadata, interface{}) bool

type MethodMetadata struct {
	CallCount int
}

var Anything ParamMatcher = func(MethodMetadata, interface{}) bool {
	return true
}

func NthCall(n int) ParamMatcher {
	return func(meta MethodMetadata, value interface{}) bool {
		return meta.CallCount == n
	}
}
