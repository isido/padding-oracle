package padding

import "fmt"

func CheckBlocksize(blocksize int) error {
	if blocksize < 1 || blocksize > 255 {
		err := fmt.Errorf("Blocksize must be between 1 and 255 (was %d.", blocksize)
		return err
	}
	return nil
}

// add PKCS7 type padding
func PadPKCS7(b []byte, blocksize int) ([]byte, error) {

	err := CheckBlocksize(blocksize)
	if err != nil {
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

// remove PKCS7 type padding, return error if the padding is not well-formed
func UnpadPKCS7(b []byte) ([]byte, error) {

	last_byte := b[len(b)-1]

	// various checks for padding (also introduce nice timing oracles...)
	if last_byte < 1 || last_byte > 255 {
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
