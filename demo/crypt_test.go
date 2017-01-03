package demo

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestEnDeCryption(t *testing.T) {

	// test with 100 random plaintexts, ivs and keys - the end result
	// should be the same

	for i := 0; i < 100; i++ {
		blocksize := 16
		iv := make([]byte, blocksize)
		_, err := rand.Read(iv)
		if err != nil {
			panic("Cannot read random bytes for IV")
		}
		key := make([]byte, blocksize)
		_, err = rand.Read(key)
		if err != nil {
			panic("Cannot read random bytes for key")
		}

		pt_size := rand.Int31n(1000) + 24
		pt := make([]byte, pt_size)
		_, err = rand.Read(pt)
		if err != nil {
			panic("Cannot read random bytes for message")
		}

		ct, err := Encrypt(pt, iv, key, blocksize)
		if err != nil {
			t.Errorf("Encrypt: encryption failed: %q", err)
		}
		pt2, err := Decrypt(ct, iv, key, blocksize)

		if err != nil {
			t.Errorf("Decrypt: decryption failed: %q", err)
		}

		if !bytes.Equal(pt, pt2) {
			t.Errorf("Decryption (%q) produced different result the original (%q)"+
				" with IV %q, key %q and blocksize %q",
				pt2, pt, iv, key, blocksize)
		}
	}

}
