package bitarray

import (
	"math/rand"
	"testing"
)

func TestBitArray_Set_Get_simple(t *testing.T) {
	t.Parallel()
	ba := New(1, 100)
	for i := 0; i < 100; i++ {
		ba.Set(i, 1)
	}
	for i := 0; i < 100; i++ {
		if v := ba.Get(i); v != 1 {
			t.Errorf("wa at %d got %d want %d", i, v, 1)
		}
	}
}

func TestBitArray_Set_Get_overwrite(t *testing.T) {
	t.Parallel()
	ba := New(2, 10)
	ba.Set(0, 1)
	ba.Set(0, 2)
	if v := ba.Get(0); v != 2 {
		t.Errorf("rewrite failed %d", v)
	}
}

func TestBitArray_Set_Get_random(t *testing.T) {
	t.Parallel()
	const size = 1000
	const bits = 7
	ba := New(bits, size)
	arr := make([]byte, size)

	for i := 0; i < size; i++ {
		v := byte(rand.Intn(1 << bits))
		ba.Set(i, v)
		arr[i] = v
	}

	for i := 0; i < size; i++ {
		if v := ba.Get(i); v != arr[i] {
			t.Errorf("wa at %d got %d want %d", i, v, arr[i])
		}
	}
}

func TestBitArray_Set_Get_random_overwrite(t *testing.T) {
	t.Parallel()
	const size = 10000
	bits := []byte{1, 2, 3, 4, 5, 6, 7}

	for bit := range bits {
		ba := New(bits[bit], size)
		arr := make([]byte, size)

		for j := 0; j < 5; j++ {
			for i := 0; i < size; i++ {
				v := byte(rand.Intn(1 << bits[bit]))
				ba.Set(i, v)
				v = byte(rand.Intn(1 << bits[bit]))
				ba.Set(i, v)
				v = byte(rand.Intn(1 << bits[bit]))
				ba.Set(i, v)
				v = byte(rand.Intn(1 << bits[bit]))
				ba.Set(i, v)

				arr[i] = v
			}

			for i := 0; i < size; i++ {
				if v := ba.Get(i); v != arr[i] {
					t.Errorf("wa at %d got %d want %d", i, v, arr[i])
				}
			}
		}
	}
}

func TestBitArray_Get_zero(t *testing.T) {
	t.Parallel()
	const size = 1000
	bits := []byte{1, 2, 3, 4, 5, 6, 7}

	for bit := range bits {
		ba := New(bits[bit], size)

		for i := 0; i < size; i++ {
			if v := ba.Get(i); v != 0 {
				t.Errorf("wa at %d got %d want %d", i, v, 0)
			}
		}
	}
}
