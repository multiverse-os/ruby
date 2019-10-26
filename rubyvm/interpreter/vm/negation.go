package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretNegationInContext(vm *vm, negation ast.Negation, context Value) (Value, error) {
	target, err := vm.executeWithContext(context, negation.Target)
	if err != nil {
		return nil, err
	}

	if target.IsTruthy() {
		return vm.SingletonWithName("false"), nil
	} else {
		return vm.SingletonWithName("true"), nil
	}
}
