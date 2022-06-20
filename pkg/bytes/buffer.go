package bytes

import (
	"bytes"
	"strings"
)

// Buffer is a bytes.Buffer that implements the Stringer interface.
type Buffer struct {
	bytes.Buffer
}

// NewBuffer returns a new Buffer with the given initial size.
func NewBuffer(buf []byte) *Buffer {
	buffer := bytes.NewBuffer(buf)
	return &Buffer{*buffer}
}

// NewBufferString returns a new Buffer initialized with the given string.
func NewBufferString(s string) *Buffer {
	buffer := bytes.NewBufferString(s)
	return &Buffer{*buffer}
}

// String returns the string representation of the Buffer, which is trimmed space.
func (b *Buffer) String() string {
	return strings.TrimSpace(b.Buffer.String())
}

// Bytes returns the byte slice representation of the Buffer, which are trimmed space.
func (b *Buffer) Bytes() []byte {
	return bytes.TrimSpace(b.Buffer.Bytes())
}
