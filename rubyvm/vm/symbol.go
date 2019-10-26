package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretSymbol(vm *vm, symbol ast.Symbol) Value {
	name := symbol.Name
	maybe, ok := vm.CurrentSymbols[name]

	if !ok {
		symbol := NewSymbol(name, vm)
		vm.CurrentSymbols[name] = symbol
		return symbol
	} else {
		return maybe
	}
}
