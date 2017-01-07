package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/isido/padding-oracle/demo"
	"github.com/isido/padding-oracle/oracle"
)

func main() {
	// TODO make an illustrative GUI for this
	fmt.Println("A Demonstration of the Vaudenay or CBC-Padding Oracle Attack")
	fmt.Println("")

	var plaintext string

	if len(os.Args) > 1 {
		plaintext = os.Args[1]
	} else {
		plaintext = "This is the text to be decrypted"
	}

	fmt.Printf("Text to be decrypted: %q\n", plaintext)
	fmt.Println("-----------------------------------------")
	fmt.Println("Oracle output")
	// prepare ciphertext for the oracle and crypt it
	d := demo.MakeDemo()

	ciphertext, err := demo.Encrypt([]byte(plaintext), d.IV(), d.Key, d.Blocksize())

	if err != nil {
		panic("Cannot encrypt text") // TODO fixme
	}

	// use padding oracle to decrypt the ciphertext

	res := hex.EncodeToString(oracle.PaddingOracle(d, ciphertext, d.Blocksize()))
	fmt.Printf("Decoded: %q\n", res)

}
