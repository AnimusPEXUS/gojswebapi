package arraybuffer

import (
	"errors"
	"io"
	"syscall/js"
)

var ERR_ARRAYBUFFER_UNSUPPORTED = errors.New("ArrayBuffer unsupported")

func GetArrayBufferJSValue() (js.Value, error) {
	return js.Global().Get("ArrayBuffer"), nil
}

func IsArrayBufferSupported() (bool, error) {

	abjv, err := GetArrayBufferJSValue()
	if err != nil {
		return false, err
	}

	undef := abjv.IsUndefined()

	return !undef, nil
}

func IsArrayBuffer(value js.Value) (bool, error) {
	abjv, err := GetArrayBufferJSValue()
	if err != nil {
		return false, err
	}

	return value.InstanceOf(abjv), nil
}

var _ io.Reader = &ArrayBuffer{}

type ArrayBuffer struct {
	jsvalue js.Value
}

func NewArrayBufferFromJSValue(jsvalue js.Value) (*ArrayBuffer, error) {
	self := &ArrayBuffer{jsvalue: jsvalue}
	return self, nil
}

func NewArrayBuffer(length int) (*ArrayBuffer, error) {
	jsv_c, err := GetArrayBufferJSValue()
	if err != nil {
		return nil, err
	}

	jsv := jsv_c.New(length)

	self, err := NewArrayBufferFromJSValue(jsv)
	if err != nil {
		return nil, err
	}

	return self, nil
}

func (self *ArrayBuffer) IsArrayBuffer() (bool, error) {
	return IsArrayBuffer(self.jsvalue)
}

func (self *ArrayBuffer) Len() (int, error) {
	return self.jsvalue.Get("byteLength").Int(), nil
}

// TODO: maybe int64 is better solution, but I'm not sure
func (self *ArrayBuffer) Slice(begin int, end *int, contentType *string) (*ArrayBuffer, error) {

	begin_p := js.ValueOf(begin)
	end_p := js.Undefined()
	contentType_p := js.Undefined()

	if end != nil {
		end_p = js.ValueOf(*end)
	}

	if contentType != nil {
		contentType_p = js.ValueOf(*contentType)
	}

	ret_array := self.jsvalue.Call("slice", begin_p, end_p, contentType_p)

	return NewArrayBufferFromJSValue(ret_array)
}

func (self *ArrayBuffer) MakeReader() *ArrayBufferReader {
	return NewArrayBufferReader()
}
