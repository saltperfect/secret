package main

import (
	"fmt"

	"github.com/saltperfect/exercise/secret/encrypt"
)

func main() {

	plaintext := "hello"
	key := "this is my secret key"
	cipher, err := encrypt.Encrypt(key, plaintext)
	must(err)
	fmt.Printf("plain: %s cipher: %s\n", plaintext, cipher)

	decryptPlain, err := encrypt.Decrypt(key, cipher)
	must(err)
	fmt.Printf("plain: %s, decrypted: %s\n", plaintext, decryptPlain)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
