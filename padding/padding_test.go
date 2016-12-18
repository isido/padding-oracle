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
		res := PadPKCS7(tt.in, 8)
		if !bytes.Equal(res, tt.out) {
			t.Errorf("padPKCS7(%q, 8) => %q, want %q", tt.in, res, tt.out)
		}
	}
}
