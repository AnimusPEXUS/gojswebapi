package promise

import (
	"syscall/js"
)

type Promise struct {
	js.Value
}

func NewPromiseFromJSValue(value js.Value) (*Promise, error) {

	self := &Promise{value}

	return self, nil
}

func (self *Promise) Then(funcs ...js.Func) (*Promise, error) {

	funcs2 := make([]interface{}, len(funcs))

	for i := 0; i != len(funcs); i++ {
		funcs2[i] = funcs[i]
	}

	ret, err := NewPromiseFromJSValue(self.Value.Call("then", funcs2...))
	if err != nil {
		return nil, err
	}

	return ret, nil
}
