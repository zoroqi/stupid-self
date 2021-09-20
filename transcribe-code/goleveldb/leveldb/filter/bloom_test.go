package filter

import (
	"fmt"
	"github.com/zoroqi/transcribe-code/goleveldb/leveldb/util"
	"testing"
)

func Test_bloomFilter(t *testing.T) {
	hello := []byte("hello")
	world := []byte("world")
	donnt := []byte("don't")
	hells := []byte("hells")

	bf := NewBloomFilter(1000)
	bg := bf.NewGenerator()
	bg.Add(hello)
	bg.Add(world)

	buffer := util.NewBuffer(make([]byte,1,10))
	bg.Generate(buffer)
	bb := buffer.Bytes()

	fmt.Printf("%s, r:%t want:%t\n", hello, bf.Contains(bb, hello), true)
	fmt.Printf("%s, r:%t want:%t\n", world, bf.Contains(bb, world), true)
	fmt.Printf("%s, r:%t want:%t\n", donnt, bf.Contains(bb, donnt), false)
	fmt.Printf("%s, r:%t want:%t\n", hells, bf.Contains(bb, hells), false)
	fmt.Println("-------")

	bg.Add(donnt)
	bg.Add(hells)
	bg.Generate(buffer)
	bb = buffer.Bytes()
	// 输出会错误
	fmt.Printf("%s, r:%t want:%t\n", hello, bf.Contains(bb, hello), false)
	fmt.Printf("%s, r:%t want:%t\n", world, bf.Contains(bb, world), false)
	fmt.Printf("%s, r:%t want:%t\n", donnt, bf.Contains(bb, donnt), true)
	fmt.Printf("%s, r:%t want:%t\n", hells, bf.Contains(bb, hells), true)
}
