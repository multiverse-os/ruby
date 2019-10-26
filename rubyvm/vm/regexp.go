package vm

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretRegexpInContext(vm *vm, regexpNode ast.Regex, context Value) (Value, error) {
	return NewRegexp(vm, regexpNode.Value), nil
}
