package padding

import "fmt"

// add PKCS7 type padding
func PadPKCS7(b []byte, blocksize int) ([]byte, error) {

	if blocksize < 1 || blocksize > 255 {
		err := fmt.Errorf("PadPKCS7: Blocksize must be between 0 and 256 (was %d).",
			blocksize)
		return nil, err
	}

	rem := blocksize - (len(b) % blocksize)
	if rem == 0 {
		rem = blocksize
	}

	padding := make([]byte, rem, rem)
	for i, _ := range padding {
		padding[i] = byte(rem)
	}

	return append(b, padding...), nil
}

func UnpadPKCS7(b []byte) ([]byte, error) {

	last_byte := b[len(b)-1]

	// various checks for padding (also introduce nice timing oracles...)
	if last_byte < 0 || last_byte > 255 {
		err := fmt.Errorf("UnpadPKCS7: Invalid value for the last byte of padding: %d", last_byte)
		return nil, err
	}

	idx := len(b) - int(last_byte)

	if idx < 0 || idx > (len(b)-1) {
		err := fmt.Errorf("UnpadPKCS7: Invalid padding value %d.", idx)
		return nil, err
	}

	for i := idx; i < len(b); i++ {
		if b[i] != last_byte {
			err := fmt.Errorf("UnpadPKCS7: Invalid value (%d) in the padding, expected %d", b[i], last_byte)
			return nil, err
		}
	}

	return b[:idx], nil
}
