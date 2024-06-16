package container

type BytesList struct {
	buf []byte
}

func NewBytesList(b ...byte) *BytesList {
	return &BytesList{
		buf: b,
	}
}

func (b *BytesList) Append(bytes ...byte) {
	b.buf = append(b.buf, bytes...)
}

func (b *BytesList) Join(bl *BytesList) *BytesList {
	return &BytesList{
		buf: append(b.buf, bl.buf...),
	}
}

func (b *BytesList) Bytes() []byte {
	return b.buf
}

func (b *BytesList) Len() int {
	return len(b.buf)
}

func (b *BytesList) String() string {
	return string(b.buf)
}
