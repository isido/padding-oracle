package main

import (
	"fmt"
	"os"
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

	// crypt it with CBC + AES256

	// pad the text, if necessary

	// encrypt it

	// start decryption

	// generate false padding

	// test it

	// repeat

	// until go the first byte

	// repeat until the message is decrypted

}
