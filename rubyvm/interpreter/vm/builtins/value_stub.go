package builtins

// this type repesents the shared behavior and data of all Ruby Values
// all values will need to store the methods that are defined on them
// (in addition to their class, and other information)
type valueStub struct {
	eigenclass_methods map[string]Method
	class              Class

	stringer      func() string
	prettyPrinter func() string

	instance_variables map[string]Value
	attrs              map[string]Value
}

func (valueStub *valueStub) initialize() {
	valueStub.eigenclass_methods = make(map[string]Method)
	valueStub.instance_variables = make(map[string]Value)
	valueStub.attrs = make(map[string]Value)
	valueStub.prettyPrinter = func() string { return valueStub.String() }
}

// Method Lookup //

/*
  1. Methods defined in the object's singleton class (i.e. the object itself)
  2. Modules mixed into the singleton class in reverse order of inclusion
  3. Methods defined by the object's class
	4. Modules included into the object's class in reverse order of inclusion
	5. Methods defined by the object's superclass, i.e. inherited methods
  6. Once BasicObject is reached, start at 1 with "method_missing" method
  7. Fail. Loudly.
*/

func (valueStub *valueStub) Method(name string) Method {
	//	  1. Methods defined in the object's singleton class (i.e. the object itself)
	m, ok := valueStub.eigenclass_methods[name]
	if ok {
		return m
	}

	//    2. Modules mixed into the singleton class in reverse order of inclusion
	// FIXME: respect step 2 here

	//	  3. Methods defined by the object's class
	for _, method := range valueStub.class.InstanceMethods() {
		if method.Name() == name {
			return method
		}
	}

	m, ok = valueStub.class.eigenclassMethods()[name]
	if ok {
		return m
	}

	//		4. Modules included into the object's class in reverse order of inclusion
	// FIXME: this should be reversed (should be fixed in Include method)
	for _, module := range valueStub.class.includedModules() {
		m, ok := module.eigenclassMethods()[name]
		if ok {
			return m
		}
	}

	//    5. Methods defined by the object's superclass, i.e. inherited methods
	super := valueStub.class.SuperClass()
	for super != nil {
		m, ok := super.eigenclassMethods()[name]
		if ok {
			return m
		}

		m, err := super.InstanceMethod(name)
		if err == nil {
			return m
		}

		for _, module := range super.includedModules() {
			m, ok := module.eigenclassMethods()[name]
			if ok {
				return m
			}
		}

		if super.String() == "BasicObject" {
			break
		}

		super = super.SuperClass()
	}

	return nil
}

func (valueStub *valueStub) Methods() []Method {
	values := make([]Method, 0, len(valueStub.eigenclass_methods))
	for _, m := range valueStub.eigenclass_methods {
		values = append(values, m)
	}

	return values
}

func (valueStub *valueStub) AddMethod(m Method) {
	valueStub.eigenclass_methods[m.Name()] = m
}

func (valueStub *valueStub) RemoveMethod(m Method) {
	delete(valueStub.eigenclass_methods, m.Name())
}

func (valueStub *valueStub) String() string {
	return valueStub.stringer()
}

func (valueStub *valueStub) PrettyPrint() string {
	return valueStub.prettyPrinter()
}

func (valueStub *valueStub) setStringer(stringer func() string) {
	valueStub.stringer = stringer
}

func (valueStub *valueStub) setPrettyPrinter(printer func() string) {
	valueStub.prettyPrinter = printer
}

func (valueStub *valueStub) Class() Class {
	return valueStub.class
}

func (valueStub *valueStub) eigenclassMethods() map[string]Method {
	return valueStub.eigenclass_methods
}

func (valueStub *valueStub) GetInstanceVariable(name string) Value {
	return valueStub.instance_variables[name]
}

func (valueStub *valueStub) SetInstanceVariable(name string, value Value) {
	valueStub.instance_variables[name] = value
}

func (valueStub *valueStub) GetClassVariable(name string) Value {
	return valueStub.class.classVariable(name)
}

func (valueStub *valueStub) SetClassVariable(name string, value Value) {
	valueStub.class.setClassVariable(name, value)
}

func (v *valueStub) IsTruthy() bool {
	return true
}

func (v *valueStub) GetAttribute(name string) (Value, bool) {
	something, ok := v.attrs[name]
	return something, ok
}

func (v *valueStub) SetAttribute(name string, value Value) {
	v.attrs[name] = value
}
