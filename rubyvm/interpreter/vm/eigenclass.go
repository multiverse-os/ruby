package vm

import (
	"github.com/multiverse-os/ruby/rubyvm/ast"

	. "github.com/multiverse-os/ruby/rubyvm/interpreter/vm/builtins"
)

func interpretEigenclassInContext(
	vm *vm,
	eigenclassNode ast.Eigenclass,
	context Value,
) (Value, error) {

	blockContext, err := vm.executeWithContext(context, eigenclassNode.Target)

	if err != nil {
		return nil, err
	}

	vm.inEigenclassBlock = true
	defer func() { vm.inEigenclassBlock = false }()

	return vm.executeWithContext(blockContext, eigenclassNode.Body...)
}
