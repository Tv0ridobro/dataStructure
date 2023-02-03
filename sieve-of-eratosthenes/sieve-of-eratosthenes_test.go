package sieveoferatosthenes

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestSieveOfEratosthenes_GetDelimiters(t *testing.T) {
	t.Parallel()
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	s := New(101)
	if v := s.PrimeFactorization(100); !slices.Equal(v, []int{2, 2, 5, 5}) {
		t.Errorf("wrong answer %v %v", v, []int{2, 2, 5, 5})
	}
	if v := s.PrimeFactorization(97); !slices.Equal(v, []int{97}) {
		t.Errorf("wrong answer %v %v", v, []int{97})
	}
	if v := s.IsPrime(97); v != true {
		t.Errorf("wrong answer %v %v", v, true)
	}
	if v := s.IsPrime(98); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(0); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(1); v != false {
		t.Errorf("wrong answer %v %v", v, false)
	}
	if v := s.IsPrime(2); v != true {
		t.Errorf("wrong answer %v %v", v, true)
	}
	if v := s.Primes(); !slices.Equal(v, primes) {
		t.Errorf("wrong answer %v %v", v, primes)
	}
}
