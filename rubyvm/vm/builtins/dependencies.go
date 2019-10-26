package builtins

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
)

type ArgEvaluator interface {
	EvaluateArgInContext(ast.Node, Value) (Value, error)
}

type ClassProvider interface {
	ClassWithName(string) Class
}

type SingletonProvider interface {
	SingletonWithName(string) Value
	NewSingletonWithName(string, Value)

	SymbolWithName(string) Value
	AddSymbol(Value)
}

type StackProvider interface {
	CurrentStack() string
	UnshiftStackFrame(string, string, int)
	ShiftStackFrame()
}

type MethodProvider interface {
	AddMethod(name string, context Value, body func(self Value, block Block, args ...Value) (Value, error))
}

type Provider interface {
	ArgEvaluator() ArgEvaluator
	ClassProvider() ClassProvider
	SingletonProvider() SingletonProvider
	StackProvider() StackProvider
	MethodProvider() MethodProvider
}
