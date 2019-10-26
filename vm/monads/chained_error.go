package monads

type chainedError struct {
	v interface{}
	f func(interface{}) interface{}
}

func (e chainedError) Value() (returnValue interface{}) {
	defer func() {
		_ = recover()
	}()

	returnValue = e.f(nil)
	return returnValue
}

func (e chainedError) Compose(other ChainedError) ChainedError {
	original := e.f
	e.f = func(thunk interface{}) interface{} {
		return other.(chainedError).f(original(thunk))
	}
	return e
}

type ChainedError interface {
	Value() interface{}
	Compose(ChainedError) ChainedError
}

func Error(failableFunc func(thunk interface{}) interface{}) (chainedErr chainedError) {
	return chainedError{f: failableFunc}
}
