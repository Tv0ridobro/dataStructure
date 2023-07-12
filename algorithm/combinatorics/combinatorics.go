package combinatorics

import (
	"github.com/Tv0ridobro/data-structure/slices"
)

// NextPermutation returns next permutation.
func NextPermutation(s []int) []int {
	next := make([]int, len(s))
	copy(next, s)
	for i := len(s) - 2; i >= 0; i-- {
		if next[i] >= next[i+1] {
			continue
		}
		z := i + 1
		for j := i + 1; j < len(s); j++ {
			if next[j] >= next[i] && next[j] <= next[z] {
				z = j
			}
		}
		next[i], next[z] = next[z], next[i]
		next = append(next[:i+1], slices.Reverse(next[i+1:])...)
		break
	}
	return next
}
