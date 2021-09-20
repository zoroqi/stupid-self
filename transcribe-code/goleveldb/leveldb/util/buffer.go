package util

import (
	"bytes"
	"io"
)

const smallBufferSize = 64

// 兼容32/64
const maxInt = int(^uint(0) >> 1)

type Buffer struct {
	buf []byte
	off int //offset
}

func (b *Buffer) Bytes() []byte {
	// 循环使用
	return b.buf[b.off:]
}

func (b *Buffer) String() string {
	if b == nil {
		return "<nil>"
	}
	//return string(b.Bytes())
	// 可能更好吧
	return string(b.buf[b.off:])
}

func (b *Buffer) Len() int {
	return len(b.buf) - b.off
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
}

func (b *Buffer) Truncate(n int) {
	if n == 0 {
		b.Reset()
		return
	}
	if n < 0 || n > b.Len() {
		// 前缀没啥用了
		panic("leveldb/util.Buffer: truncation out of range")
	}
	b.buf = b.buf[:b.off+n]
}

// 尝试进行扩容
func (b *Buffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l {
		// 这个扩容方案很有意思
		b.buf = b.buf[:l+n]
		return l, true
	}
	return 0, false
}

func makeSlice(n int) []byte {
	// 这里为啥要改变panic? 不是很理解
	// 我是不会进行这个recover
	defer func() {
		if recover() != nil {
			panic(bytes.ErrTooLarge)
		}
	}()
	return make([]byte, n)
}

// 扩容n个字节
func (b *Buffer) grow(n int) int {
	//
	m := b.Len()
	if m == 0 && b.off != 0 {
		b.Reset()
	}
	if i, ok := b.tryGrowByReslice(n); ok {
		return i
	}

	// 初始化
	if b.buf == nil && n <= smallBufferSize {
		b.buf = make([]byte, n, smallBufferSize)
		return 0
	}

	c := cap(b.buf)
	// 重新使用容量
	if n <= c/2-m {
		copy(b.buf, b.buf[b.off:])
	} else if c > maxInt-c-n {
		// 超出容量
		panic(bytes.ErrTooLarge)
	} else {
		// 扩容
		buf := makeSlice(2*c + n)
		copy(buf, b.buf[b.off:])
		b.buf = buf
	}
	b.off = 0
	b.buf = b.buf[:m+n]
	return m
}

func (b *Buffer) Alloc(n int) []byte {
	if n < 0 {
		panic("leveldb/util.Buffer.Alloc: negative count")
	}
	m, ok := b.tryGrowByReslice(n)
	if !ok {
		m = b.grow(n)
	}
	return b.buf[m:]
}

// 返回实际长度, 测试使用
func (b *Buffer) Grow(n int) {
	if n < 0 {
		panic("negative count")
	}
	m := b.grow(n)
	// TODO @WHY 这里为啥会将n的部分抹掉
	b.buf = b.buf[:m]
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	m, ok := b.tryGrowByReslice(len(p))
	if !ok {
		//
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

const MixRead = 512

func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	for {
		i := b.grow(MixRead)
		// 抹掉扩容部分进行拷贝
		b.buf = b.buf[:i]
		// 扩容, 神奇的扩容
		m, e := r.Read(b.buf[i:cap(b.buf)])
		if m < 0 {
			panic(" reader returned negative count from Read")
		}
		b.buf = b.buf[:i+m]
		n += int64(m)
		if e == io.EOF {
			return n, nil
		}
		if e != nil {
			return n, e
		}
	}
}

// TODO 不懂
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	// 为啥和Read的判断反着
	if b.off < len(b.buf) {
		nBytes := b.Len()
		m, e := w.Write(b.buf[b.off:])
		if m > nBytes {
			panic("invalid Write count")
		}
		b.off += m
		if e != nil {
			return n, e
		}
		if m != nBytes {
			return n, io.ErrShortWrite
		}
	}
	b.Reset()
	return
}

func (b *Buffer) WriteByte(c byte) error {
	m, ok := b.tryGrowByReslice(1)
	if !ok {
		m = b.grow(1)
	}
	b.buf[m] = c
	return nil
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.off >= len(b.buf) {
		b.Reset()
		if len(p) == 0 {
			return
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return
}

func (b *Buffer) Next(n int) []byte {
	m := b.Len()
	if n > m {
		n = m
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	return data
}

func (b *Buffer) ReadByte() (c byte, err error) {
	if b.off >= len(b.buf) {
		b.Reset()
		return 0, io.EOF
	}
	c = b.buf[b.off]
	b.off++
	return c, nil
}

func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	slice, err := b.readSlice(delim)
	line = append(line, slice...)
	return
}

func (b *Buffer) readSlice(delim byte) (line []byte, err error) {
	i := bytes.IndexByte(b.buf[b.off:], delim)
	end := b.off + i + 1
	if i < 0 {
		end = len(b.buf)
		err = io.EOF
	}
	line = b.buf[b.off:end]
	b.off = end
	return line, err
}

// 这里我更倾向于传一个长度
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{buf: buf}
}
