package blob

import (
	"errors"
	"syscall/js"

	"github.com/AnimusPEXUS/wasmtools/arraybuffer"
	"github.com/AnimusPEXUS/wasmtools/promise"
)

var ERR_BLOB_UNSUPPORTED = errors.New("Blob unsupported")

func GetBlobJSValue() js.Value {
	return js.Global().Get("Blob")
}

func IsBlobSupported() bool {
	return !GetBlobJSValue().IsUndefined()
}

func IsBlob(value js.Value) (bool, error) {
	return value.InstanceOf(GetBlobJSValue()), nil
}

var _ io.Reader = &Blob{}

type Blob struct {
	jsvalue js.Value
}

func NewBlobFromJSValue(jsvalue js.Value) (*Blob, error) {
	self := &Blob{jsvalue: jsvalue}
	return self, nil
}

func (self *Blob) Read(p []byte) (n int, err error) {

}

func (self *Blob) IsBlob() (bool, error) {
	return IsBlob(self.jsvalue)
}

func (self *Blob) Size() (int, error) {
	return self.jsvalue.Get("size").Int(), nil
}

func (self *Blob) Type() (string, error) {
	return self.jsvalue.Get("type").String(), nil
}

func (self *Blob) ArrayBuffer() (*arraybuffer.ArrayBuffer, error) {

	blob_arraybuffer_result := self.jsvalue.Call("arrayBuffer")

	pro, err := promise.NewPromiseFromJSValue(blob_arraybuffer_result)
	if err != nil {
		return nil, err
	}

	psucc := make(chan bool)
	perr := make(chan bool)
	var array_data js.Value

	pro.Then(
		js.FuncOf(func(
			this js.Value,
			args []js.Value,
		) interface{} {
			if len(args) == 0 {
				perr <- true
				return false
			}
			array_data = args[0].Get("data")
			psucc <- true
			return false
		},
		),
		js.FuncOf(func(
			this js.Value,
			args []js.Value,
		) interface{} {
			perr <- true
			return false
		},
		),
	)

	select {
	case <-psucc:
		return arraybuffer.NewArrayBufferFromJSValue(array_data)
	case <-perr:
		return nil, errors.New("error getting Blob's ArrayBuffer")
	}

	return nil, errors.New("invalid behavior")
}

// TODO: maybe int64 is better solution, but I'm not sure
func (self *Blob) Slice(start *int, end *int, contentType *string) (*Blob, error) {
	start_p := js.Undefined()
	end_p := js.Undefined()
	contentType_p := js.Undefined()

	if start != nil {
		start_p = js.ValueOf(*start)
	}

	if end != nil {
		end_p = js.ValueOf(*end)
	}

	if contentType != nil {
		contentType_p = js.ValueOf(*contentType)
	}

	ret_blob := self.jsvalue.Call("slice", start_p, end_p, contentType_p)

	return NewBlobFromJSValue(ret_blob)
}

// TODO: maybe later :)
// func (self *Blob) Stream() (*ReadableStream, error)

func (self *Blob) Text() (*promise.Promise, error) {
	blob_text_result := self.jsvalue.Call("text")
	pro, err := promise.NewPromiseFromJSValue(blob_text_result)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
