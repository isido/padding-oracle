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

func TestCorrectPadding(t *testing.T) {

	// test with 100 random instances of en/decrypting

	for i := 0; i < 100; i++ {

		d := MakeDemo()

		pt_size := rand.Int31n(1000) + 24
		pt := make([]byte, pt_size)
		_, err := rand.Read(pt)
		if err != nil {
			panic("Cannot read random bytes for message")
		}

		ct, err := Encrypt(pt, d.IV(), d.Key, d.Blocksize())
		if err != nil {
			t.Errorf("Encrypt: encryption failed: %q", err)
		}

		if !d.CorrectPadding(ct) {
			t.Errorf("CorrectPadding() => false, expected true")
		}
	}
}

func TestInCorrectPadding(t *testing.T) {

	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	key := []byte{8, 7, 6, 5, 4, 3, 2, 1, 16, 15, 14, 13, 12, 11, 10, 9}
	pt := []byte{60, 60, 60, 60, 60, 60, 60, 60, 60}

	d := DemoOracle{iv, key, 16}

	ct, err := Encrypt(pt, iv, key, 16)
	if err != nil {
		t.Errorf("Encrypt: encryption failed: %q", err)
	}

	if !d.CorrectPadding(ct) {
		t.Errorf("CorrectPadding(%q) => false, expected true", ct)
	}

	// corrupt ciphertext and padding
	ct[len(ct)-1] = ct[len(ct)-1] ^ 4

	if d.CorrectPadding(ct) {
		t.Errorf("CorrectPadding(%q) => false, expected true", ct)
	}
}
