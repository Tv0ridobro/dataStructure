// Package sieveoferatosthenes implements a sieve of eratosthenes.
// See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes for more details.
package sieveoferatosthenes

// SieveOfEratosthenes represents a sieve of eratosthenes.
// Zero value of SieveOfEratosthenes is invalid sieve, should be used only with New().
type SieveOfEratosthenes struct {
	numbers []int
	primes  []int
}

// New returns an initialized SieveOfEratosthenes.
func New(size int) *SieveOfEratosthenes {
	s := make([]int, size)
	primes := make([]int, 0)
	for i := 2; i < size; i++ {
		if s[i] == 0 {
			s[i] = i
			primes = append(primes, i)
		}
		for _, e := range primes {
			if s[i] < e || e*i >= size {
				break
			}
			s[i*e] = e
		}
	}
	return &SieveOfEratosthenes{numbers: s, primes: primes}
}

// IsPrime returns true if given number is prime,
// false otherwise.
func (s SieveOfEratosthenes) IsPrime(i int) bool {
	if i <= 0 {
		return false
	}
	return s.numbers[i] == i
}

// PrimeFactorization returns prime numbers, such that their product gives the original number.
func (s SieveOfEratosthenes) PrimeFactorization(i int) []int {
	delim := make([]int, 0)
	for s.numbers[i] != i {
		delim = append(delim, s.numbers[i])
		i /= s.numbers[i]
	}
	delim = append(delim, i)
	return delim
}

// Primes return all primes number up to size of sieve.
func (s SieveOfEratosthenes) Primes() []int {
	c := make([]int, len(s.primes))
	copy(c, s.primes)
	return c
}
