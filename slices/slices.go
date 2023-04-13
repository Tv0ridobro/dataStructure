package slices

// Equal returns true if given slices contain same elements without checking capacity,
// false otherwise.
func Equal[T comparable](f, s []T) bool {
	if len(f) != len(s) {
		return false
	}
	for i := range f {
		if f[i] != s[i] {
			return false
		}
	}
	return true
}

// Reverse returns reversed slice.
func Reverse[T any](f []T) []T {
	s := make([]T, len(f))
	for i := 0; i < len(f); i++ {
		s[i] = f[len(f)-i-1]
	}
	return s
}

func Generate[T any](size int, f func(i int) T) []T {
	s := make([]T, size)
	for i := 0; i < size; i++ {
		s[i] = f(i)
	}

	return s
}
