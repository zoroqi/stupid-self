package opt

import (
	"github.com/zoroqi/transcribe-code/goleveldb/leveldb/cache"
	"github.com/zoroqi/transcribe-code/goleveldb/leveldb/filter"
)


type Cacher interface {
	New(capacity int) cache.Cacher
}

type Options struct {
	AltFilters []filter.Filter
	BlockCacher Cacher
	BlockCacheCapacity int
	BlockCacheEvictRemoved bool
	BlockRestartInterval int
	BlockSize int
	CompactionExpandLimitFactor int
	CompactionGPOverlapsFactor int
}
