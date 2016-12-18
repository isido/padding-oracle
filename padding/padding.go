package padding

// add PKCS7 type padding
func PadPKCS7(b []byte, blocksize uint) []byte {

	divisor := int(blocksize)

	rem := divisor - (len(b) % divisor)
	if rem == 0 {
		rem = divisor
	}

	padding := make([]byte, rem, rem)
	for i, _ := range padding {
		padding[i] = byte(rem)
	}

	return append(b, padding...)
}

func UnpadPKCS7(b []byte, blocksize int) ([]byte, error) {

	return b, nil
}
