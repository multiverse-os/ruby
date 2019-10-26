package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
	monads "github.com/multiverse-os/ruby/rubyvm/vm/monads"
)

// TODO: If monads is going to be required we should migrate it otherwise
// preferably lets just drop the requirement using whatever means necessary

func interpretBareReferenceInContext(vm *vm, ref ast.BareReference, context Value) (Value, error) {

	var name string = ref.Name
	var returnErr error

	maybe := monads.Maybe(func() interface{} {
		m, err := vm.localVariableStack.Retrieve(name)
		if err == nil {
			return m
		} else {
			return nil
		}
	}).OrSome(monads.Maybe(func() interface{} {
		m, ok := vm.ObjectSpace[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(monads.Maybe(func() interface{} {
		m, ok := vm.CurrentClasses[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(monads.Maybe(func() interface{} {
		m, ok := vm.CurrentModules[name]
		if ok {
			return m
		} else {
			return nil
		}
	})).OrSome(monads.Maybe(func() interface{} {
		maybeMethod := context.Method(name)
		if maybeMethod == nil {
			return nil
		}

		value, err := maybeMethod.Execute(context, nil)
		if err != nil {
			returnErr = err
			return nil
		} else {
			return value
		}
	}))

	if returnErr != nil {
		return nil, returnErr
	}

	value, ok := maybe.Value().(Value)
	if ok {
		return value, nil
	} else {
		return nil, NewNameError(
			name,
			context.String(),
			context.Class().String(),
			vm.stack.String(),
		)
	}
}
