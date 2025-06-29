package kamacache

// ByteView 只读的字节视图，用于缓存数据
type ByteView struct {
	b []byte
}

func (b ByteView) Len() int {
	return len(b.b)
}

func (b ByteView) ByteSLice() []byte {
	return cloneBytes(b.b)
}//wejopeenfferwr

func (b ByteView) String() string {
	return string(b.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
