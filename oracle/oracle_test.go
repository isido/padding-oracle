package oracle

import (
	"bytes"
	"testing"
	"testing/quick"
)

type byteslices [][]byte

func TestMakeByteSlice(t *testing.T) {

	tests := byteslices{
		{},
		{1},
		{2, 2},
		{3, 3, 3},
		{4, 4, 4, 4},
		{5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8},
	}

	for i, exp := range tests {
		out := makeByteSlice(i)
		if !bytes.Equal(out, exp) {
			t.Errorf("makeByteSlice(%q) => %q, want %q", i, out, exp)
		}
	}
}

func TestMakeByteSliceLength(t *testing.T) {
	f := func(x int) bool {
		if x < 0 {
			x = -x
		}
		if x > 255 {
			return true // TODO figure out how to limit values
		}
		xs := makeByteSlice(x)
		return len(xs) == x
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
