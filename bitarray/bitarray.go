package bitarray

type BitArray struct {
	bitsPerElement byte
	data           []byte
}

// New returns new bitArray. bitsPerElement is a value in [1, 7].
func New(bitsPerElement byte, size int) BitArray {
	return BitArray{
		bitsPerElement: bitsPerElement,
		data:           make([]byte, size*int(bitsPerElement)/8+1),
	}
}

func (b BitArray) Set(i int, value byte) {
	leftBorderBits := i * int(b.bitsPerElement)
	leftBorderBytes := leftBorderBits / 8
	rightBorderBits := (i+1)*int(b.bitsPerElement) - 1
	rightBorderBytes := rightBorderBits / 8

	old := b.Get(i)
	// to destroy old value with itself
	value = value ^ old

	switch rightBorderBytes - leftBorderBytes {
	case 0:
		value = value << (8 - int(b.bitsPerElement) - (leftBorderBits % 8))
		b.data[leftBorderBytes] ^= value
	case 1:
		b.data[leftBorderBytes] ^= value >> ((rightBorderBits + 1) % 8)
		b.data[rightBorderBytes] ^= value << (8 - (rightBorderBits % 8) - 1)
	}
}

func (b BitArray) Get(i int) byte {
	leftBorderBits := i * int(b.bitsPerElement)
	leftBorderBytes := leftBorderBits / 8
	rightBorderBits := (i+1)*int(b.bitsPerElement) - 1
	rightBorderBytes := rightBorderBits / 8

	var value byte
	switch rightBorderBytes - leftBorderBytes {
	case 0:
		value = b.data[leftBorderBytes] << (leftBorderBits % 8) >> (8 - b.bitsPerElement)
	case 1:
		value = b.data[leftBorderBytes]
		// zeroing bytes
		value = value << (8 - int(b.bitsPerElement) + ((rightBorderBits + 1) % 8))
		value = value >> (8 - int(b.bitsPerElement))
		value |= b.data[rightBorderBytes] >> (8 - ((rightBorderBits + 1) % 8))
	}

	return value
}
