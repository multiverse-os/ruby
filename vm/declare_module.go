package vm

import (
	"strings"

	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
)

func interpretModuleDeclarationInContext(vm *vm, moduleNode ast.ModuleDecl, context Value) (Value, error) {

	originalName := vm.currentModuleName

	defer func() {
		vm.currentModuleName = originalName
	}()

	var currentModule Module
	var fullModuleName string
	if vm.currentModuleName != "" {
		fullModuleName = strings.Join([]string{vm.currentModuleName, moduleNode.FullName()}, "::")
		var ok bool
		currentModule, ok = vm.CurrentClasses[vm.currentModuleName]
		if !ok {
			currentModule = vm.CurrentModules[vm.currentModuleName]
		}
	} else {
		fullModuleName = moduleNode.FullName()
	}

	theModule := NewModule(moduleNode.Name, vm)
	vm.CurrentModules[fullModuleName] = theModule

	if currentModule != nil {
		currentModule.SetConstant(moduleNode.Name, theModule)
	}

	vm.currentModuleName = fullModuleName
	_, err := vm.executeWithContext(theModule, moduleNode.Body...)
	if err != nil {
		return nil, err
	}

	return theModule, nil
}
