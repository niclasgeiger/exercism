package paasio

import (
	"io"
	"sync"
)

type DefaultReadCounter struct {
	reader io.Reader
	n      int64
	nops   int
	mutex  sync.Mutex
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &DefaultReadCounter{
		reader: reader,
	}
}

func (r *DefaultReadCounter) ReadCount() (n int64, nops int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.n, r.nops
}

func (r *DefaultReadCounter) Read(p []byte) (n int, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.nops++
	n, err = r.reader.Read(p)
	r.n += int64(n)
	return n, err
}
