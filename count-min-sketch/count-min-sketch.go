package countminsketch

import (
	"math"

	"github.com/dolthub/maphash"
)

type CountMinSketch[T comparable] struct {
	matrix  [][]uint64
	hashers []maphash.Hasher[T]
}

func New[T comparable](errorRate, confidence float64) CountMinSketch[T] {
	n := int(math.Ceil(math.E / errorRate))
	k := int(math.Ceil(math.Log(1 / confidence)))

	hashers := make([]maphash.Hasher[T], k)
	for i := 0; i < k; i++ {
		hashers[i] = maphash.NewHasher[T]()
	}

	matrix := make([][]uint64, n)
	for i := 0; i < k; i++ {
		matrix[i] = make([]uint64, k)
	}

	return CountMinSketch[T]{
		matrix:  matrix,
		hashers: hashers,
	}
}

func (c CountMinSketch[T]) Insert(elem T) {
	c.InsertN(elem, 1)
}

func (c CountMinSketch[T]) InsertN(elem T, count uint64) {
	for i, hasher := range c.hashers {
		hash := hasher.Hash(elem)
		c.matrix[i][hash%uint64(len(c.matrix[i]))] += count
	}
}

func (c CountMinSketch[T]) Count(elem T) uint64 {
	var min uint64 = math.MaxUint64
	for i, hasher := range c.hashers {
		hash := hasher.Hash(elem)
		if value := c.matrix[i][hash%uint64(len(c.matrix[i]))]; value < min {
			min = value
		}
	}
	return min
}
