package bloomfilter

import (
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	t.Parallel()
	probability := 0.01
	bloom := New[int](1000, probability)
	for i := 0; i < 1000; i++ {
		bloom.Insert(i)
	}
	for i := 0; i < 1000; i++ {
		if !bloom.Contains(i) {
			t.Errorf("doesn't containt %d", i)
		}
	}

	misses := 0
	size := 10000
	for i := 0; i < size; i++ {
		if bloom.Contains(i + 10000) {
			misses++
		}
	}
	if float64(misses)/float64(size) >= probability*2 {
		t.Errorf("too much misses")
	}
}
