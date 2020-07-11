package ws

import (
	"bytes"
	"errors"
	"io"
	"net"
	"sync"
	"syscall/js"
	"time"

	wasmtools_arraybuffer "github.com/AnimusPEXUS/wasmtools/arraybuffer"
	wasmtools_blob "github.com/AnimusPEXUS/wasmtools/blob"
	"github.com/gorilla/websocket"
)

var _ net.Conn = &WSNetConn{}

// TODO: redo. this is fast temporary copypasta
type websocketAddr struct{}

func (a websocketAddr) Network() string {
	return "websocket"
}

func (a websocketAddr) String() string {
	return "websocket/unknown-addr"
}

type WSNetConnOptions struct {
	WS *WS
}

type WSNetConn struct {
	options *WSNetConnOptions

	read_buffer *bytes.Buffer

	inbound_messages       []js.Value
	inbound_messages_mutex sync.Mutex
}

func NewWSNetConn(options *WSNetConnOptions) *WSNetConn {
	self := &WSNetConn{
		options:          options,
		read_buffer:      nil,
		inbound_messages: make([]js.Value, 0),
	}

	return self
}

func (self *WSNetConn) onMessage(event js.Value) {
	self.inbound_messages_mutex.Lock()
	defer self.inbound_messages_mutex.Unlock()

	js_data = event.Get("data")
	self.inbound_messages = append(self.inbound_messages, js_data)

}

func (self *WSNetConn) readBufferWriter() {
	self.inbound_messages_mutex.Lock()

	js_data := self.inbound_messages[0]
	self.inbound_messages = append(self.inbound_messages[0:0], self.inbound_messages[1:]...)

	self.inbound_messages_mutex.Unlock()

	var re io.Reader

	if wasmtools_blob.IsBlob(js_data) {
		re, err = wasmtools_blob.NewBlobFromJSValue(js_data)
		if err != nil {
			// TODO: error?
		}
	} else if wasmtools_arraybuffer.IsArrayBuffer(js_data) {
		re, err = wasmtools_arraybuffer.NewArrayBufferFromJSValue(js_data)
		if err != nil {
			// TODO: error?
		}
	} else {
		// TODO: error?
	}

	_, err = io.Copy(self.read_buffer, re)
	if err != nil {
		// TODO: error?
	}

}

func (self *WSNetConn) Read(b []byte) (n int, err error) {

make_read:
	if self.read_buffer != nil {
		n, err = self.read_buffer.Read(b)
		if self.read_buffer.Len() == 0 {
			self.read_buffer = nil
		}
		return
	}

	{
		var mtype int
		var data []byte

		for {
			mtype, data, err = self.gorilla_conn.ReadMessage()
			if err != nil {
				break
			}
			if mtype == websocket.BinaryMessage {
				goto ok
			}
		}

		return len(data), err

	ok:
		self.read_buffer = bytes.NewBuffer(data)
		goto make_read
	}

	// magic!
	return 666, errors.New("Something strange! Who will you call?")
}

func (self *WSNetConn) Write(b []byte) (n int, err error) {
	err = self.options.WS.Send(b)
	n = len(b)
	return
}

func (self *WSNetConn) Close() error {
	return self.options.WS.Close()
}

func (self *WSNetConn) LocalAddr() net.Addr {
	return websocketAddr{}
}

func (self *WSNetConn) RemoteAddr() net.Addr {
	return websocketAddr{}
}

func (self *WSNetConn) SetDeadline(t time.Time) error {
	return nil
}

func (self *WSNetConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (self *WSNetConn) SetWriteDeadline(t time.Time) error {
	return nil
}
