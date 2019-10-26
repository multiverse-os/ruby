package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretDefinedKeyword(vm *vm, defined ast.Defined, context Value) (Value, error) {
	_, err := vm.executeWithContext(context, defined.Node)
	if err == nil {
		return NewString("garbage", vm), nil
	} else {
		return vm.SingletonWithName("nil"), nil
	}
}
