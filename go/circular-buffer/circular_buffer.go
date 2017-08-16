package circular

const testVersion = 4

type Buffer struct {
	size   int
	oldest int
	ptr    int
	ring   []byte
}

type BufferError struct {
	message string
}

func (err BufferError) Error() string {
	return err.message
}

func NewBufferError(message string) error {
	return &BufferError{
		message: message,
	}
}

var (
	BufferReadError  = NewBufferError("read error")
	BufferWriteError = NewBufferError("write error")
)

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size:   size,
		oldest: 0,
		ptr:    0,
		ring:   make([]byte, size),
	}
}

func (buf *Buffer) ReadByte() (byte, error) {
	if b := buf.ring[buf.oldest]; b != byte(0) {
		buf.ring[buf.oldest] = byte(0)
		buf.oldest = (buf.oldest + 1) % buf.size
		return b, nil
	}
	return 0, BufferReadError
}

func (buf *Buffer) WriteByte(b byte) error {
	if buf.ring[buf.ptr] == byte(0) {
		buf.ring[buf.ptr] = b
		buf.ptr = (buf.ptr + 1) % buf.size
		return nil
	}
	return BufferWriteError
}

func (buf *Buffer) Overwrite(b byte) {
	err := buf.WriteByte(b)
	if err != nil {
		buf.ring[buf.oldest] = b
		buf.oldest = (buf.oldest + 1) % buf.size
	}
}

func (buf *Buffer) Reset() {
	buf.ring = make([]byte, buf.size)
	buf.ptr = 0
	buf.oldest = 0
}
