package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretSuperCall(vm *vm, superCall ast.SuperclassMethodImplCall, context Value) (Value, error) {
	methodName := vm.stack.Frames[0].Method
	superClass := context.Class().SuperClass()
	superMethod, err := superClass.InstanceMethod(methodName)
	if err != nil {
		return nil, NewNoMethodError(methodName, superClass.String(), superClass.Class().String(), vm.stack.String())
	}

	return superMethod.Execute(context, nil)
}
