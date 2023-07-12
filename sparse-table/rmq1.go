package sparsetable

import (
	"math"

	"golang.org/x/exp/constraints"

	mathx "github.com/Tv0ridobro/data-structure/math"
)

// RMQpm1 represents a special case of RMQ.
// All neighbour elements should differ in 1.
// Allows to answer rmq queries in constant time with linear preprocessing time.
type RMQpm1[T constraints.Integer] struct {
	blocks        []block[T]
	precalculated [][][]T
	st            *SparseTable[T]

	k int
}

type block[T constraints.Integer] struct {
	id     int
	offset T
}

// NewRMQpm1 returns RMQpm1 struct.
// All neighbour elements should differ in 1.
func NewRMQpm1[T constraints.Signed](elements []T) RMQpm1[T] {
	n := len(elements)
	k := int(0.5 * math.Log2(float64(n)))
	if k == 0 {
		k = 1
	}

	// precalculate all possible blocks
	precalculated := make([][][]T, 1<<(k-1))
	for i := 0; i < len(precalculated); i++ {
		precalculated[i] = make([][]T, k)
		for j := 0; j < k; j++ {
			precalculated[i][j] = make([]T, k)
		}
	}

	for i := 0; i < len(precalculated); i++ {
		var prev T
		for j := 0; j < k-1; j++ {
			if (i & (1 << j)) != 0 {
				precalculated[i][j+1][j+1] = prev + 1
				prev++
			} else {
				precalculated[i][j+1][j+1] = prev - 1
				prev--
			}
		}

		for j := 0; j < k; j++ {
			for z := j + 1; z < k; z++ {
				precalculated[i][j][z] = mathx.Min(precalculated[i][j][z-1], precalculated[i][z][z])
			}
		}
	}

	// define type of each block
	blocks := make([]block[T], 0, 1<<k)
loop:
	for i := 0; i < n; {
		var t int
		offset := elements[i]
		prev := elements[i]
		i++

		for j := 0; j < k-1; j++ {
			if i == n {
				break loop
			}

			if prev+1 == elements[i] {
				t += 1 << j
			}
			prev = elements[i]
			i++
		}

		blocks = append(blocks, block[T]{
			id:     t,
			offset: offset,
		})
	}

	mins := make([]T, 0)
	for i := range blocks {
		val := precalculated[blocks[i].id][0][k-1]
		mins = append(mins, val)
	}

	return RMQpm1[T]{
		blocks:        blocks,
		precalculated: precalculated,
		st: New[T](func(t T, t2 T) T {
			if t <= t2 {
				return t
			}
			return t2
		}, mins),
		k: k,
	}
}

// Min returns minimum in range l, r.
func (rmq *RMQpm1[T]) Min(l, r int) T {
	if r < l || l < 0 {
		var zero T
		return zero
	}

	k := len(rmq.precalculated[0])

	l1 := l / k
	r1 := r / k

	if l1 == r1 {
		id := rmq.blocks[l1].id
		return rmq.precalculated[id][l%k][r%k] + rmq.blocks[l1].offset
	}

	id := rmq.blocks[l1].id
	f1 := rmq.precalculated[id][l%k][k-1] + rmq.blocks[l1].offset

	id = rmq.blocks[r1].id
	f2 := rmq.precalculated[id][0][r%k] + rmq.blocks[r1].offset

	if l1+1 == r {
		return mathx.Min(f1, f2)
	}
	f3 := rmq.st.Query(l1+1, r1-1)

	return mathx.Min(mathx.Min(f1, f2), f3)
}
