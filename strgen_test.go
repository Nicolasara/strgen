package strgen

import (
	"testing"
)

// TestGenerate checks that generate returns a string of the right length.
func TestGenerate(t *testing.T) {
	const MaxLengthString = 16
	for i := 0; i < 16; i++ {
		str := generate(i)
		strLen := len(str)
		if strLen != i {
			t.Errorf("generate returned the string: %q, which was of size %d, when %d was expected", str, strLen, i)
		}
	}
}
