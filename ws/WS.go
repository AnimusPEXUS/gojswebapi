package ws

import (
	"errors"
	"syscall/js"

	utils_panic "github.com/AnimusPEXUS/utils/panic"
)

type ReadyState int

const (
	ReadyState_CONNECTING ReadyState = 0
	ReadyState_OPEN                  = 1
	ReadyState_CLOSING               = 2
	ReadyState_CLOSED                = 3
)

// if both url and js_value are specified, js_value is used
type WSOptions struct {
	URL       *string
	JSValue   *js.Value
	Protocols []string

	OnClose   func(js.Value) // function(event)
	OnError   func(js.Value) // function(event)
	OnMessage func(js.Value) // function(event)
	OnOpen    func(js.Value) // function(event)
}

type WS struct {
	options *WSOptions
}

func NewWS(options *WSOptions) (*WS, error) {

	// mutex_checkable := utils_sync.NewMutexCheckable(true)
	// open_cond := sync.NewCond(mutex_checkable)

	var wsoc js.Value

	if options.JSValue != nil {
		wsoc = *options.JSValue
		options.JSValue = &wsoc
		options.URL = &([]string{wsoc.Get("url").String()}[0])
		// options.URL = wsoc.Get("url").String()
	} else {
		wsoc_go := js.Global().Get("WebSocket")
		if wsoc_go.IsUndefined() {
			return nil, errors.New("WebSocket is undefined")
		}
		url := *options.URL
		wsoc = wsoc_go.New(url, js.Undefined()) //options.Protocols
		options.JSValue = &wsoc
	}

	options.JSValue.Set(
		"onopen",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if options.OnOpen != nil {
					options.OnOpen(args[0])
				}
				return nil
			},
		),
	)

	self := &WS{
		options: options,
	}

	return self, nil
}

func (self *WS) Close(code *int, reason *string) error {
	if reason != nil && code == nil {
		return errors.New("reason can't be specified without code")
	}

	var args []interface{}

	if code != nil {
		args = append(args, *code)
		if reason != nil {
			args = append(args, *reason)
		}
	}

	self.options.JSValue.Call("close", args...)
	return nil
}

func (self *WS) Send(value js.Value) js.Value {
	return self.options.JSValue.Call("send", value)
}

///////////////// properties

func (self *WS) BinaryTypeGet() (ret string, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("binaryType").String()
	return
}

func (self *WS) BinaryTypeSet(value string) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	self.options.JSValue.Set("binaryType", value)
	return
}

func (self *WS) BufferedAmountGet() (ret int, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("bufferedAmount").Int()
	return
}

// func (self *WS) ExtensionsGet()

// TODO:
// func (self *WS) OnCloseSet()

// TODO:
// func (self *WS) OnCloseGet()

// TODO:
// func (self *WS) OnErrorSet()

// TODO:
// func (self *WS) OnErrorGet()

// TODO:
// func (self *WS) OnMessageSet()

// TODO:
// func (self *WS) OnMessageGet()

// TODO:
// func (self *WS) OnOpenGet()

// TODO:
// func (self *WS) OnOpenSet()

func (self *WS) ProtocolGet() (ret string, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("protocol").String()
	return
}

func (self *WS) ReadyStateGet() (ret ReadyState, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = ReadyState(self.options.JSValue.Get("readyState").Int())
	return
}

func (self *WS) URLGet() (ret string, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("url").String()
	return
}
