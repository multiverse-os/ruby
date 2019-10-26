package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretWeakBooleanOr(vm *vm, weakOr ast.WeakLogicalOr, context Value) (Value, error) {
	lhs, err := vm.executeWithContext(context, weakOr.LHS)
	if err != nil {
		return nil, err
	}

	if lhs.IsTruthy() {
		return lhs, nil
	}

	return vm.executeWithContext(context, weakOr.RHS)
}
