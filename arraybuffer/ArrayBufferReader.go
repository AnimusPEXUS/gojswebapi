package arraybuffer

import (
	"github.com/dennwc/dom/js"
)

type ArrayBufferReader struct {
	ab     *ArrayBuffer
	lenght int
	done   int
	EOF    bool
}

func NewArrayBufferReader(ab *ArrayBuffer) (*ArrayBufferReader, error) {

	length, err := ab.Len()
	if err != nil {
		return 0, err
	}

	self := &ArrayBufferReader{
		ab:     ab,
		lenght: length,
		done:   0,
	}

	return self
}

func (self *ArrayBufferReader) Read(p []byte) (n int, err error) {
	len_p := len(p)

	var end_index int

	{
		selfdonelenp := self.done + len_p
		if selfdonelenp > self.lenght {
			end_index = self.done - self.lenght
			self.EOF = true
		} else {
			end_index = selfdonelenp
		}
	}

	self.ab.Slice(self.done, end_index)

	uint8array := js
}
