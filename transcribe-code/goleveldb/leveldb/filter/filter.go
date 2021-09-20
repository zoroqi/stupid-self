package filter

type Filter interface {
	Name() string
	NewGenerator() FilterGenerator
	Contains(filter, key []byte) bool
}

type FilterGenerator interface {
	Add(key []byte)
	Generate(b Buffer)
}
// 这就是鸭子的好处, 实现可能在任何地方.
//这种隐形的联系好吗?
type Buffer interface {
	Alloc(n int) []byte
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
}
