package ws

import (
	"errors"
	"log"
	"syscall/js"

	"github.com/AnimusPEXUS/gojswebapi/events"
	utils_panic "github.com/AnimusPEXUS/utils/panic"
)

type WSReadyState int

const (
	WSReadyState_CONNECTING WSReadyState = 0
	WSReadyState_OPEN                    = 1
	WSReadyState_CLOSING                 = 2
	WSReadyState_CLOSED                  = 3
)

// if both url and js_value are specified, js_value is used
type WSOptions struct {
	URL       *string
	JSValue   *js.Value
	Protocols []string

	OnClose   func(*events.CloseEvent)   // function(event)
	OnError   func(*events.ErrorEvent)   // function(event)
	OnMessage func(*events.MessageEvent) // function(event)
	OnOpen    func(*events.Event)        // function(event)
}

type WS struct {
	options *WSOptions
}

func NewWS(options *WSOptions) (*WS, error) {

	self := &WS{
		options: options,
	}

	var wsoc js.Value

	if options.JSValue != nil {
		wsoc = *options.JSValue
		options.JSValue = &wsoc
		options.URL = &([]string{wsoc.Get("url").String()}[0])
	} else {
		wsoc_go := js.Global().Get("WebSocket")
		if wsoc_go.IsUndefined() {
			return nil, errors.New("WebSocket is undefined")
		}
		url := *options.URL
		wsoc = wsoc_go.New(url, js.Undefined()) //options.Protocols
		options.JSValue = &wsoc
	}

	err := self.SetOnOpen(options.OnOpen)
	if err != nil {
		return nil, err
	}

	err = self.SetOnClose(options.OnClose)
	if err != nil {
		return nil, err
	}

	err = self.SetOnMessage(options.OnMessage)
	if err != nil {
		return nil, err
	}

	err = self.SetOnError(options.OnError)
	if err != nil {
		return nil, err
	}

	return self, nil
}

func (self *WS) SetOnOpen(f func(*events.Event)) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()

	if f == nil {
		self.options.JSValue.Set("onopen", js.Undefined())
		self.options.OnOpen = nil
		return
	}

	self.options.OnOpen = f

	self.options.JSValue.Set(
		"onopen",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if self.options.OnOpen != nil {
					ev, err := events.NewEventFromJSValue(args[0])
					if err != nil {
						return err
					}
					self.options.OnOpen(ev)
				} else {
					self.SetOnOpen(nil)
				}
				return nil
			},
		),
	)
	return nil
}

func (self *WS) SetOnClose(f func(*events.CloseEvent)) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()

	if f == nil {
		self.options.JSValue.Set("onclose", js.Undefined())
		self.options.OnOpen = nil
		return
	}

	self.options.OnClose = f

	self.options.JSValue.Set(
		"onclose",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if self.options.OnClose != nil {
					ev, err := events.NewCloseEventFromJSValue(args[0])
					if err != nil {
						return err
					}
					self.options.OnClose(ev)
				} else {
					self.SetOnClose(nil)
				}
				return nil
			},
		),
	)
	return nil
}

func (self *WS) SetOnMessage(f func(*events.MessageEvent)) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()

	if f == nil {
		self.options.JSValue.Set("onmessage", js.Undefined())
		self.options.OnOpen = nil
		return
	}

	self.options.OnMessage = f

	self.options.JSValue.Set(
		"onmessage",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if self.options.OnMessage != nil {
					ev, err := events.NewMessageEventFromJSValue(args[0])
					if err != nil {
						return err
					}
					self.options.OnMessage(ev)
				} else {
					self.SetOnMessage(nil)
				}
				return nil
			},
		),
	)
	return nil
}

func (self *WS) SetOnError(f func(*events.ErrorEvent)) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()

	if f == nil {
		self.options.JSValue.Set("onerror", js.Undefined())
		self.options.OnOpen = nil
		return
	}

	self.options.OnError = f

	self.options.JSValue.Set(
		"onerror",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				if self.options.OnError != nil {
					ev, err := events.NewErrorEventFromJSValue(args[0])
					if err != nil {
						return err
					}
					self.options.OnError(ev)
				} else {
					self.SetOnError(nil)
				}
				return nil
			},
		),
	)
	return nil
}

func (self *WS) Close(code *int, reason *string) (err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()

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

func (self *WS) Send(value js.Value) (err error) {
	log.Print("WS Send called")
	defer func() {
		err = utils_panic.PanicToError()
	}()
	self.options.JSValue.Call("send", value)
	return
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

func (self *WS) ProtocolGet() (ret string, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("protocol").String()
	return
}

func (self *WS) ReadyStateGet() (ret WSReadyState, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = WSReadyState(self.options.JSValue.Get("readyState").Int())
	return
}

func (self *WS) URLGet() (ret string, err error) {
	defer func() {
		err = utils_panic.PanicToError()
	}()
	ret = self.options.JSValue.Get("url").String()
	return
}
