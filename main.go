package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/isido/padding-oracle/demo"
	"github.com/isido/padding-oracle/oracle"
)

func main() {
	fmt.Println("A Demonstration of the Vaudenay or Padding Oracle Attack")

	var plaintext string

	if len(os.Args) > 1 {
		plaintext = os.Args[1]
	} else {
		plaintext = "This is the text to be decrypted"
	}

	fmt.Println("This is the text that is to be decrypted:")
	fmt.Println("-----------------------------------------")
	fmt.Println(plaintext)
	fmt.Println("-----------------------------------------")

	// prepare ciphertext for the oracle and crypt it
	d := demo.MakeDemo()

	ciphertext, err := demo.Encrypt([]byte(plaintext), d.Iv, d.Key, d.Blocksize())

	if err != nil {
		panic("Cannot encrypt text") // TODO fixme
	}

	// use padding oracle to decrypt the ciphertext

	fmt.Println("This is the padding oracle output:")
	fmt.Println("----------------------------------")
	res := hex.EncodeToString(oracle.PaddingOracle(d, ciphertext, d.Blocksize()))
	fmt.Println(res)

}
