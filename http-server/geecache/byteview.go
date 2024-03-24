package geecache

type ByteView struct {
	b []byte
}

// 返回长度，在cache中要求要实现这样的接口
func (v ByteView) Len() int {
	return len(v.b)
}

// 返回字节切片类型的数组
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// 将数据转换成字符串
func (v ByteView) String() string {
	return string(v.b)
}

// 深拷贝
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
