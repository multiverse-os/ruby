package rbmarshal

import "reflect"

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch  {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

func decode(input interface{}, output reflect.Value) {
	var inputVal reflect.Value
	if input != nil {
		inputVal = reflect.ValueOf(input)
		if inputVal.Kind() == reflect.Ptr && inputVal.IsNil() {
			input = nil
		}
	}

	if input == nil {
		output.Set(reflect.Zero(output.Type()))
		return
	}

	if !inputVal.IsValid() {
		output.Set(reflect.Zero(output.Type()))
		return
	}

	outputKind := getKind(output)
	switch outputKind {
	case reflect.Bool:
		decodeBool(input, output)
	case reflect.String:
		decodeString(input, output)
	case reflect.Int:
		decodeInt(input, output)
	case reflect.Float64:
		decodeFloat64(input, output)
	case reflect.Struct:
		decodeStruct(input, output)
	case reflect.Ptr:
		decodePtr(input, output)
	case reflect.Slice:
		decodeSlice(input, output)
	case reflect.Array:
		decodeArray(input, output)
	//case reflect.Map:
	//	decodeMap(input, output)
	}
}

func decodeBool(input interface{}, output reflect.Value) {
	inputVal := reflect.Indirect(reflect.ValueOf(input))
	output.SetBool(inputVal.Bool())
}

func decodeInt(input interface{}, output reflect.Value) {
	inputVal := reflect.Indirect(reflect.ValueOf(input))
	output.SetInt(inputVal.Int())
}

func decodeFloat64(input interface{}, output reflect.Value) {
	inputVal := reflect.Indirect(reflect.ValueOf(input))
	output.SetFloat(inputVal.Float())
}

func decodeString(input interface{}, output reflect.Value) {
	inputVal := reflect.Indirect(reflect.ValueOf(input))
	output.SetString(inputVal.String())
}

func decodeStruct(input interface{}, output reflect.Value) {
	inputVal := reflect.Indirect(reflect.ValueOf(input))
	if inputVal.Type() == output.Type() {
		output.Set(inputVal)
		return
	}

	inputMap := input.(map[string]interface{})

	outputType := output.Type()
	for i := 0; i < outputType.NumField(); i++ {
		field := outputType.Field(i)
		tagValue := field.Tag.Get("ruby")
		if len(tagValue) == 0 {
			continue
		}
		tagInfo, err := parseTag(tagValue)
		if err != nil {
			return
		}
		tagName := tagInfo.name
		inputData := inputMap[tagName]

		// 反序列化到每个field上
		fieldValue := output.Field(i)
		decode(inputData, fieldValue)
	}
}

func decodeSlice(input interface{}, output reflect.Value) {
	dataVal := reflect.Indirect(reflect.ValueOf(input))

	valType := output.Type()
	valElemType := valType.Elem()
	sliceType := reflect.SliceOf(valElemType)

	valSlice := output
	if valSlice.IsNil() {
		if dataVal.Len() == 0 {
			return
		}

		valSlice = reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())
	}

	for i :=0; i<dataVal.Len(); i++{
		currentData := dataVal.Index(i).Interface()
		if valSlice.Len() <= i {
			valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
		}

		currentField := valSlice.Index(i)
		decode(currentData, currentField)
	}

	output.Set(valSlice)
}

func decodeArray(input interface{}, output reflect.Value) {
	panic("not support")
}

func decodePtr(input interface{}, output reflect.Value) {
	isNil := input == nil

	if !isNil {
		switch v := reflect.Indirect(reflect.ValueOf(input)); v.Kind() {
		case reflect.Chan,reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Slice:
			isNil = v.IsNil()
		}
	}

	if isNil {
		if !output.IsNil() && output.CanSet() {
			nilValue := reflect.New(output.Type()).Elem()
			output.Set(nilValue)
		}
	}

	outputType := output.Type()
	outputTypeElem := outputType.Elem()

	if output.CanSet() {
		realVal := output
		if realVal.IsNil()  {
			realVal = reflect.New(outputTypeElem)
		}
		decode(input, reflect.Indirect(realVal))
		output.Set(realVal)
	} else {
		decode(input, reflect.Indirect(output))
	}
}

