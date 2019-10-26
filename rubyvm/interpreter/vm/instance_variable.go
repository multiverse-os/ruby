package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretInstanceVariableInContext(
	vm *vm,
	ref ast.InstanceVariable,
	context Value,
) (Value, error) {

	return context.GetInstanceVariable(ref.Name), nil
}
