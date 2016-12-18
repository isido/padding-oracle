package padding

import (
	"bytes"
	"testing"
)

var paddingTests8 = []struct {
	in  []byte
	out []byte
}{
	{[]byte{65}, []byte{65, 7, 7, 7, 7, 7, 7, 7}},
	{[]byte{65, 65}, []byte{65, 65, 6, 6, 6, 6, 6, 6}},
	{[]byte{65, 65, 65}, []byte{65, 65, 65, 5, 5, 5, 5, 5}},
	{[]byte{65, 65, 65, 65}, []byte{65, 65, 65, 65, 4, 4, 4, 4}},
	{[]byte{65, 65, 65, 65, 65}, []byte{65, 65, 65, 65, 65, 3, 3, 3}},
	{[]byte{65, 65, 65, 65, 65, 65}, []byte{65, 65, 65, 65, 65, 65, 2, 2}},
	{[]byte{65, 65, 65, 65, 65, 65, 65}, []byte{65, 65, 65, 65, 65, 65, 65, 1}},
	{[]byte{65, 65, 65, 65, 65, 65, 65, 65}, []byte{65, 65, 65, 65, 65, 65, 65, 65, 8, 8, 8, 8, 8, 8, 8, 8}},
	{[]byte{65, 65, 65, 65, 65, 65, 65, 65, 65}, []byte{65, 65, 65, 65, 65, 65, 65, 65, 65, 7, 7, 7, 7, 7, 7, 7}},
}

func TestPaddingExamples(t *testing.T) {

	for _, tt := range paddingTests8 {
		res, err := PadPKCS7(tt.in, 8)
		if err != nil {
			t.Errorf("PadPKCS7(%q, 8) => Unexpected error: %v", tt.in, err)
		}

		if !bytes.Equal(res, tt.out) {
			t.Errorf("PadPKCS7(%q, 8) => %q, want %q", tt.in, res, tt.out)
		}
	}
}

func TestPaddingInvalidBlocksizes(t *testing.T) {

	tin := paddingTests8[0].in

	values := []int{0, -1, 256}

	for _, v := range values {
		_, err := PadPKCS7(tin, v)
		if err == nil {
			t.Errorf("PadPKCS7(..., %d) => Expected error, got none.", v)
		}
	}
}

func TestUnpaddingExamples(t *testing.T) {

	for _, tt := range paddingTests8 {
		res, _ := UnpadPKCS7(tt.out)
		if !bytes.Equal(res, tt.in) {
			t.Errorf("UnpadPKCS7(%q, 8) => %q, want %q", tt.out, res, tt.in)
		}
	}

}
