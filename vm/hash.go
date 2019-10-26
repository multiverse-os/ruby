package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretHashInContext(vm *vm, hashNode ast.Hash, context Value) (Value, error) {
	hashValue, _ := vm.CurrentClasses["Hash"].New(vm)
	hash := hashValue.(*Hash)
	for _, keyPair := range hashNode.Pairs {
		key, err := vm.executeWithContext(context, keyPair.Key)
		if err != nil {
			return nil, err
		}

		val, err := vm.executeWithContext(context, keyPair.Value)
		if err != nil {
			return nil, err
		}

		hash.Add(key, val)
	}

	return hash, nil
}
