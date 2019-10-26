package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretWeakBooleanAnd(vm *vm, weakAnd ast.WeakLogicalAnd, context Value) (Value, error) {
	lhs, err := vm.executeWithContext(context, weakAnd.LHS)
	if err != nil {
		return nil, err
	}

	if !lhs.IsTruthy() {
		return vm.SingletonWithName("nil"), nil
	}

	rhs, err := vm.executeWithContext(context, weakAnd.RHS)
	if err != nil {
		return nil, err
	}

	if !rhs.IsTruthy() {
		return vm.SingletonWithName("nil"), nil
	}

	return rhs, nil
}
