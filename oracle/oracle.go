package oracle

import (
	"bytes"
	"fmt"
)

type Distinguisher interface {
	CorrectPadding(b []byte) bool
	Blocksize() int
}

func makeByteSlice(n int) []byte {
	init := make([]byte, 1)
	init[0] = byte(n)
	return bytes.Repeat(init, n)
}

func PaddingOracle(d Distinguisher, ciphertext []byte, blocksize int) []byte {

	if len(ciphertext)%blocksize != 0 {
		panic("Incorrect length for ciphertext") // TODO fixme
	}

	buf := bytes.NewBuffer(ciphertext)

	var blocks [][]byte

	for buf.Len() > 0 {
		blocks = append(blocks, buf.Next(blocksize))
	}

	plaintext_block := make([]byte, blocksize)
	for i, block := range blocks {
		fmt.Printf("Block %d/%d\n", (i + 1), len(blocks))

		fake_block := make([]byte, blocksize)
		for i := len(fake_block) - 1; i >= 0; i-- {
			fmt.Printf("Byte %d/%d (Inverse order)...", (i + 1), len(fake_block))
			prepared_block := append(fake_block, block...)
			b := decryptAByte(d, prepared_block, i)
			v := b ^ block[i] ^ 1 // TODO fixme
			fmt.Printf("%q\n", v)
			plaintext_block[i] = v
		}
		return plaintext_block
	}

	return plaintext_block
}

func decryptAByte(d Distinguisher, prepared_ciphertext []byte, bytepos int) byte {

	for i := 0; i < 256; i++ {
		b := byte(i)
		prepared_ciphertext[bytepos] = b
		if d.CorrectPadding(prepared_ciphertext) {
			return b
		}
	}
	panic("Didn't work. A bug somewhere")
}
