package bloom_filter

import (
	"math"

	"github.com/dolthub/maphash"
)

type BloomFilter[T comparable] struct {
	hashers []maphash.Hasher[T]
	bits    []byte
	size    int
}

func New[T comparable](n int, probability float64) BloomFilter[T] {
	mFloat := -math.Log(probability) * float64(n) / (math.Ln2 * math.Ln2)
	k := int(mFloat / float64(n) * math.Ln2)
	m := int(mFloat)

	hashers := make([]maphash.Hasher[T], k)
	for i := 0; i < k; i++ {
		hashers[i] = maphash.NewHasher[T]()
	}
	return BloomFilter[T]{
		hashers: hashers,
		bits:    make([]byte, (m-1)/8+1),
		size:    m,
	}
}

func (b BloomFilter[T]) Contains(elem T) bool {
	for _, hasher := range b.hashers {
		h := hasher.Hash(elem)
		index := h % uint64(b.size)
		bucketNumber := index / 8
		if b.bits[bucketNumber]&(1<<(index%8)) == 0 {
			return false
		}
	}
	return true
}

func (b BloomFilter[T]) Insert(elem T) {
	for _, hasher := range b.hashers {
		h := hasher.Hash(elem)
		index := h % uint64(b.size)
		bucketNumber := index / 8
		b.bits[bucketNumber] = b.bits[bucketNumber] | (1 << (index % 8))
	}
}
