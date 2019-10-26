package vm

import (
	"strings"

	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
	monads "github.com/multiverse-os/ruby/rubyvm/vm/monads"
)

func interpretConstantInContext(vm *vm, constantNode ast.Constant, context Value) (Value, error) {

	maybeTarget := monads.Maybe(func() interface{} {
		if vm.currentModuleName == "" {
			return vm.CurrentClasses["Object"]
		} else {
			return nil
		}
	}).OrSome(monads.Maybe(func() interface{} {
		return vm.CurrentClasses[vm.currentModuleName]
	})).OrSome(monads.Maybe(func() interface{} {
		return vm.CurrentModules[vm.currentModuleName]
	}))

	target, ok := maybeTarget.Value().(Module)

	maybeConstant := monads.Maybe(func() interface{} {
		constant, err := target.Constant(constantNode.Name)
		if err == nil {
			return constant
		} else {
			return nil
		}
	}).OrSome(monads.Maybe(func() interface{} {
		return vm.CurrentClasses[constantNode.Name]
	})).OrSome(monads.Maybe(func() interface{} {
		return vm.CurrentModules[constantNode.Name]
	})).OrSome(monads.Maybe(func() interface{} {
		if vm.currentModuleName == "" {
			return nil
		}

		parts := strings.Split(vm.currentModuleName, "::")
		count := len(parts) - 1
		for index, _ := range parts {
			namespace := append(parts[0:(count-index)], constantNode.Name)
			nameToLookup := strings.Join(namespace, "::")

			maybe := monads.Maybe(func() interface{} {
				return vm.CurrentClasses[nameToLookup]
			}).OrSome(monads.Maybe(func() interface{} {
				return vm.CurrentModules[nameToLookup]
			}))

			something, ok := maybe.Value().(Value)
			if ok {
				return something
			}
		}

		return nil
	}))

	constant, ok := maybeConstant.Value().(Value)
	if ok {
		return constant, nil
	} else {
		return nil, NewNameError(
			constantNode.Name,
			context.String(),
			context.Class().String(),
			vm.stack.String(),
		)
	}

}
