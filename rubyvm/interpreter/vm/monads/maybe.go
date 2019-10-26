package monads

type maybeValue struct {
	v interface{}
}

func (m maybeValue) Value() interface{} {
	return m.v
}

func (m maybeValue) OrSome(v interface{}) MaybeValue {
	if m.v == nil {
		maybe, ok := v.(maybeValue)
		if ok {
			m.v = maybe.v
		} else {
			m.v = v
		}
	}

	return m
}

type MaybeValue interface {
	Value() interface{}
	OrSome(interface{}) MaybeValue
}

func Maybe(unit func() interface{}) (m maybeValue) {
	defer func() {
		recover()
	}()

	m.v = unit()
	return m
}
