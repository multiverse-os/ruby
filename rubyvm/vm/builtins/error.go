package builtins

func NewArgumentErrorClass(provider Provider) Class {
	return NewGenericClass("ArgumentError", "Exception", provider)
}

func NewStandardErrorClass(provider Provider) Class {
	return NewGenericClass("StandardError", "Exception", provider)
}
