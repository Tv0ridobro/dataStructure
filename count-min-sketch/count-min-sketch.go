// Package countminsketch implements a Count-Min Sketch: a probabilistic data
// structure that serves as a frequency table of events in a stream of data.
// See https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch for more details.
package countminsketch

import (
	"math"

	"github.com/dolthub/maphash"
)

type CountMinSketch[T comparable] struct {
	matrix  [][]uint64
	hashers []maphash.Hasher[T]
}

// New is a constructor function that creates a new count-min sketch
// with desired error rate and confidence.
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

// Insert adds an element to the count-min sketch with a count of 1.
func (c CountMinSketch[T]) Insert(elem T) {
	c.InsertN(elem, 1)
}

// InsertN adds an element to the count-min sketch with a given count.
func (c CountMinSketch[T]) InsertN(elem T, count uint64) {
	for i, hasher := range c.hashers {
		hash := hasher.Hash(elem)
		c.matrix[i][hash%uint64(len(c.matrix[i]))] += count
	}
}

// Count returns the approximate count of an element in the count-min sketch.
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
