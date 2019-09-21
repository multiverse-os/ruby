// Package linewriter provides an io.Writer which calls an emitter on each line.
package linewriter

import (
	"bytes"
	"sync"
)

// Writer is an io.Writer which buffers input, flushing
// individual lines through an emitter function.
type Writer struct {
	// the mutex locks buf.
	sync.Mutex
	// buf holds the data we haven't emitted yet.
	buf  bytes.Buffer
	emit func(p []byte)
}

func NewWriter(emitter func(p []byte)) *Writer { return &Writer{emit: emitter} }

// Write implements io.Writer.Write.
// It calls emit on each line of input, not including the newline.
// Write may be called concurrently.
func (self *Writer) Write(p []byte) (int, error) {
	self.Lock()
	defer self.Unlock()

	total := 0
	for len(p) > 0 {
		emit := true
		i := bytes.IndexByte(p, '\n')
		if i < 0 {
			// No newline, we will buffer everything.
			i = len(p)
			emit = false
		}
		n, err := self.buf.Write(p[:i])
		if err != nil {
			return total, err
		}
		total += n
		p = p[i:]
		if emit {
			// Skip the newline, but still count it.
			p = p[1:]
			total++
			self.emit(self.buf.Bytes())
			self.buf.Reset()
		}
	}
	return total, nil
}
