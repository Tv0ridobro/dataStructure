// Package bloom_filter implements a bloom filter.
// See https://en.wikipedia.org/wiki/Bloom_filter for more details.
package bloom_filter

import (
	"math"

	"github.com/dolthub/maphash"
)

// BloomFilter represents a bloom filter.
// Zero value of BloomFilter is bloom filter, should be used only with New().
type BloomFilter[T comparable] struct {
	hashers []maphash.Hasher[T]
	bits    []byte
	size    int
}

// New returns an initialized bloom filter.
// n is expected number of elements,
// probabilityOfMistake is probability of false positive BloomFilter.Contains().
func New[T comparable](n int, probabilityOfMistake float64) BloomFilter[T] {
	mFloat := -math.Log(probabilityOfMistake) * float64(n) / (math.Ln2 * math.Ln2)
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

// Contains returns true if skiplist contains given value, false otherwise.
// Contains can give false positive results.
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

// Insert inserts value in a bloom filter.
func (b BloomFilter[T]) Insert(elem T) {
	for _, hasher := range b.hashers {
		h := hasher.Hash(elem)
		index := h % uint64(b.size)
		bucketNumber := index / 8
		b.bits[bucketNumber] |= (1 << (index % 8))
	}
}
