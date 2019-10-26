package vm

import (
	"strings"

	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
	. "github.com/multiverse-os/ruby/rubyvm/vm/builtins"
	monads "github.com/multiverse-os/ruby/rubyvm/vm/monads"
)

func interpretDoubleQuotedStringInContext(vm *vm, stringValue ast.InterpolatedString, context Value) (Value, error) {

	str := stringValue.Value
	currentBytes := []byte{}
	bytesToInterpret := [][]byte{}
	insideInterpolation := false
	for i := 0; i < len(str); i++ {
		if insideInterpolation {
			currentBytes = append(currentBytes, str[i])
		}

		if str[i] == '}' && i > 0 && str[i-1] != '\\' {
			insideInterpolation = false
			bytesToInterpret = append(bytesToInterpret, currentBytes)
			currentBytes = []byte{}
		} else if str[i] == '#' && len(str) > i && str[i+1] == '{' {
			insideInterpolation = true
		}
	}

	for _, bytes := range bytesToInterpret {
		substringToReplace := string(bytes)
		rubyValue, err := vm.EvaluateStringInContext(substringToReplace[1:len(substringToReplace)-1], context)
		if err != nil {
			return nil, err
		}

		valueAsString := monads.Maybe(func() interface{} {
			method := rubyValue.Method("to_s")
			if method == nil {
				return nil
			}

			result, err := method.Execute(rubyValue, nil)
			if err != nil {
				return nil
			}

			return result.(*StringValue).RawString()
		}).OrSome(rubyValue.String()).Value().(string)

		str = strings.Replace(str, "#"+substringToReplace, valueAsString, 1)
	}

	return NewString(str, vm), nil
}
