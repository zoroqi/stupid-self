package filter

import "github.com/zoroqi/transcribe-code/goleveldb/leveldb/util"

func bloomHash(key []byte) uint32 {
	return util.Hash(key, 0xbc9f1d34)
}

type bloomFilter int

func (bloomFilter) Name() string {
	return "leveldb.BuiltinBloomFilter"
}

// 这里有一种老版hbase中bloomfilter实现方案.
// bloomfilter只提供算法部分, 不进行存储.
func (f bloomFilter) Contains(filter, key []byte) bool {
	nBytes := len(filter) - 1
	if nBytes < 1 {
		return false
	}

	nBits := uint32(nBytes * 8)
	k := filter[nBytes]

	if k > 30 {
		return true
	}

	kh := bloomHash(key)
	delta := (kh >> 17) | (kh << 15)
	for j := uint8(0); j < k; j++ {
		bitops := kh % nBits
		// 偏移量计算和寻址那个更快呢?
		if (uint32(filter[bitops/8]) & (1 << (bitops % 8))) == 0 {
			return false
		}
		kh += delta
	}
	return true
}

// 这个不是同步添加, 是批量产生添加, 每一次处理都会将过去的key清理.
//这里的清理不是reset这个[]byte, 而是在之上继续写入新的, 但可能会产生扩容, 导致老key无法在判断.
//处理又外部buffer进行存储
type bloomFilterGenerator struct {
	n int
	k uint8

	keyHashes []uint32
}

func (g *bloomFilterGenerator) Add(key []byte) {
	g.keyHashes = append(g.keyHashes, bloomHash(key))
}

func (g *bloomFilterGenerator) Generate(b Buffer) {
	nBits := uint32(len(g.keyHashes) * g.n)
	if nBits < 64 {
		nBits = 64
	}
	nBytes := (nBits + 7) / 8
	nBits = nBytes * 8
	// 这里进行扩容, 这里扩容需要buffer初试空间必须是0.
	//长度计算是基于nBits算的, 而不是实际空间计算的, 不是0的会产生一定偏移量, 在转换成byte数组会产生问题.
	//一直使用同一个buffer,产生问题, 测试也可以复现.
	// 不知道为什么要这么设计. 这里简单来说就是脱离这个项目就没法用了.
	dest := b.Alloc(int(nBytes) + 1)
	dest[nBytes] = g.k
	for _, kh := range g.keyHashes {
		delta := (kh >> 17) | (kh << 15)
		for j := (uint8(0)); j < g.k; j++ {
			bitpos := kh % nBits
			dest[bitpos/8] |= (1 << (bitpos % 8))
			kh += delta
		}
	}
	g.keyHashes = g.keyHashes[:0]
}

func (f bloomFilter) NewGenerator() FilterGenerator {
	k := uint8(f * 69 / 100) // 0.69 =~ ln(2)
	if k < 1 {
		k = 1
	} else if k > 30 {
		k = 30
	}
	return &bloomFilterGenerator{
		n: int(f),
		k: k,
	}

}

func NewBloomFilter(bitsPerKey int) Filter {
	return bloomFilter(bitsPerKey)
}