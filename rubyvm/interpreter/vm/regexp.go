package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretRegexpInContext(vm *vm, regexpNode ast.Regex, context Value) (Value, error) {
	return NewRegexp(vm, regexpNode.Value), nil
}
