package paasio

import (
	"io"
	"sync"
)

type DefaultReadWriteCounter struct {
	readWriter io.ReadWriter
	nRead      int64
	nopsRead   int
	nWrite     int64
	nopsWrite  int
	mutex      sync.Mutex
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &DefaultReadWriteCounter{
		readWriter: readWriter,
	}

}

func (rw *DefaultReadWriteCounter) ReadCount() (n int64, nops int) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	return rw.nRead, rw.nopsRead
}

func (rw *DefaultReadWriteCounter) Read(p []byte) (n int, err error) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.nopsRead++
	n, err = rw.readWriter.Read(p)
	rw.nRead += int64(n)
	return n, err
}

func (rw *DefaultReadWriteCounter) WriteCount() (n int64, nops int) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	return rw.nWrite, rw.nopsWrite
}

func (rw *DefaultReadWriteCounter) Write(p []byte) (n int, err error) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.nopsWrite++
	n, err = rw.readWriter.Write(p)
	rw.nWrite += int64(n)
	return n, err

}
