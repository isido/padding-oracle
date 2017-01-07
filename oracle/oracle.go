package oracle

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

type Distinguisher interface {
	CorrectPadding(b []byte) bool
	Blocksize() int
	IV() []byte
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
	var pts []byte

	for buf.Len() > 0 {
		blocks = append(blocks, buf.Next(blocksize))
	}

	// last byte
	for i, block := range blocks {
		/*
			// last byte
			fmt.Printf("Block %d/%d...", (i + 1), len(blocks))
			b := lastByteOracle(d, block)
			var p byte

			if i == 0 {
				p = b ^ d.IV()[len(block)-1]
			} else {
				p = b ^ blocks[i-1][len(block)-1]

			}
			fmt.Printf("%q\n", p)
		*/
		// whole block
		fmt.Printf("Block %d/%d: ", (i + 1), len(blocks))
		var prev []byte
		if i == 0 {
			if d.IV() != nil {
				prev = d.IV()
			} else {
				fmt.Println("No IV available, skipping the first block") // TODO fixme
				continue
			}
		} else {
			prev = blocks[i-1]
		}
		pt := blockDecryptionOracle(d, prev, block)
		pts = append(pts, pt...)
		fmt.Println(string(pt)) // TODO fixme
		fmt.Println(pt)
	}

	return pts
}

// Vaudenay's Last Word Oracle (3.1)
func lastByteOracle(d Distinguisher, y []byte) byte {

	r := make([]byte, d.Blocksize())
	_, _ = rand.Read(r) // TODO check for errors
	b := d.Blocksize() - 1

	for i := 0; i < 256; i++ {
		r[b] = byte(i)
		ry := append(r, y...)
		if d.CorrectPadding(ry) {
			// first byte found

			// Vaudenay's original algorithm
			// checked if there were more correct
			// bytes (i.e. the padding is 02 02, 03 03 03
			// etc, but we'll skip the full implementation
			// for now
			bytes_found := byte(1)
			for n := b; n >= 1; n-- {
				r[n] = r[n] ^ 1
				ry = append(r, y...)
				if !d.CorrectPadding(ry) {
					if bytes_found > 1 {
						fmt.Printf("Bytes found %d", bytes_found)
					}
					// TODO implement me there -> (*)
					break
				}
				bytes_found++
			}
			// TODO implement me here (*)
			return byte(i ^ 1)
		}
	}
	panic("lastByteOracle(): Didn't work. A bug somewhere")
}

func blockDecryptionOracle(d Distinguisher, prev, y []byte) []byte {

	r := make([]byte, d.Blocksize())
	_, _ = rand.Read(r) // TODO check for errors

	var known []byte

	//	a_b := lastByteOracle(d, y) ^ prev[len(prev)-1]
	a_b := lastByteOracle(d, y) //^ prev[len(prev)-1]
	known = append(known, a_b)

	// k loops from 2 to number of bytes in the block
	for k := 2; k <= d.Blocksize(); k++ {
		a := byteOracle(d, known, y)
		known = append([]byte{a}, known...)
	}

	pt := make([]byte, d.Blocksize())
	for i, _ := range pt {
		pt[i] = known[i] ^ prev[i]
	}
	//return known
	return pt
}

// Vaudenay 3.2 Block Decryption Oracle
// known is a slice of known bytes from the block (from the end),
// corresponding to a_j ... a_b from the paper
// y is the block to decrypt
func byteOracle(d Distinguisher, known, y []byte) byte {

	r := make([]byte, d.Blocksize())
	// 2. pic r1, ... , r_j-1 at random
	_, _ = rand.Read(r) // TODO check for errors
	j := d.Blocksize() - len(known)
	no := len(known) + 1

	// 1. take r_k = a_k xor (b - j + 2) for k = j, ... b
	for i, v := range known {
		r[j+i] = v ^ byte(no)
	}

	// 3. r = r1 ... r_j-2 (r_j-1 xor i) r_j ... r_b
	for i := 0; i < 256; i++ {
		r[j-1] = byte(i)
		ry := append(r, y...)
		if d.CorrectPadding(ry) {
			return byte(i) ^ byte(no)
		}
	} // 4. if O(r|y) = false, increment i
	panic("byteOracle(): Didn't work. A bug somewhere")
}
