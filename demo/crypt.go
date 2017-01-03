package demo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/isido/padding-oracle/padding"
)

type DemoOracle struct {
	Iv        []byte
	Key       []byte
	blocksize int
}

func (d DemoOracle) CorrectPadding(b []byte) bool {

	_, err := Decrypt(b, d.Iv, d.Key, d.Blocksize())

	if err != nil {
		return true
	} else {
		return false
	}
}

func (d DemoOracle) Blocksize() int {
	return d.blocksize
}

// TODO create demo with random iv and key
func MakeDemo() DemoOracle {

	blocksize := 16
	iv := make([]byte, blocksize)
	_, err := rand.Read(iv)
	if err != nil {
		panic("Cannot read random bytes for IV") // TODO fixme
	}

	key := make([]byte, blocksize)
	_, err = rand.Read(key)
	if err != nil {
		panic("Cannot read random bytes for key") // TODO fixme
	}
	return DemoOracle{iv, key, blocksize}
}

// Encrypt bytes with CBC-AES256
// Pad plaintext with PKCS7
// Note: bad api, iv and key can be mistaken for each other
// maybe structify this at some point
func Encrypt(plaintext, iv, key []byte, blocksize int) ([]byte, error) {

	err := padding.CheckBlocksize(blocksize)
	if err != nil {
		return nil, err
	}

	if len(iv) != blocksize {
		err := fmt.Errorf("IV length (%d) must be the same as the blocksize (%d)", len(iv), blocksize)
		return nil, err
	}

	if len(key) != blocksize {
		err := fmt.Errorf("Key length (%d) must be the same as the blocksize (%d)", len(key), blocksize)
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padded_plaintext, err := padding.PadPKCS7(plaintext, blocksize)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(padded_plaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padded_plaintext)

	return ciphertext, nil

}

// Decrypt bytes with CBC-AES 256
// Remove padding and report errors with invalid padding,
// thus creating a padding oracle
func Decrypt(ciphertext, iv, key []byte, blocksize int) ([]byte, error) {

	err := padding.CheckBlocksize(blocksize)
	if err != nil {
		return nil, err
	}

	if len(iv) != blocksize {
		err := fmt.Errorf("IV length (%d) must be the same as the blocksize (%d)", len(iv), blocksize)
		return nil, err
	}

	if len(key) != blocksize {
		err := fmt.Errorf("Key length (%d) must be the same as the blocksize (%d)", len(iv), blocksize)
		return nil, err
	}

	if len(ciphertext)%blocksize != 0 {
		err := fmt.Errorf("Ciphertext does not divide evenly into blocks (ct lenght %d, blocksize %d)", len(ciphertext),
			blocksize)
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	padded_plaintext := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(padded_plaintext, ciphertext)

	unpadded_plaintext, err := padding.UnpadPKCS7(padded_plaintext)
	if err != nil {
		return nil, err
	}

	return unpadded_plaintext, nil

}
