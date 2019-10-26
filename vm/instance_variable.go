package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretInstanceVariableInContext(vm *vm, ref ast.InstanceVariable, context Value) (Value, error) {
	return context.GetInstanceVariable(ref.Name), nil
}
