package builtins

type NilClass struct {
	valueStub
	classStub
}

func NewNilClass(provider Provider) Class {
	n := &NilClass{}
	n.initialize()
	n.setStringer(n.String)
	n.class = provider.ClassProvider().ClassWithName("Class")
	n.superClass = provider.ClassProvider().ClassWithName("Object")
	return n
}

func (n *NilClass) String() string {
	return "NilClass"
}

func (n *NilClass) Name() string {
	return "NilClass"
}

type nilInstance struct {
	valueStub
}

func (class *NilClass) New(provider Provider, args ...Value) (Value, error) {
	n := &nilInstance{}
	n.initialize()
	n.setStringer(n.String)
	n.class = class

	return n, nil
}

func (n *nilInstance) String() string {
	return ""
}

func (n *nilInstance) IsTruthy() bool {
	return false
}
