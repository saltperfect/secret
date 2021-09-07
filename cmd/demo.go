package main

import (
	"fmt"

	"github.com/saltperfect/exercise/secret"
)
func main() {
	valt:= secret.NewVault("my-fake-key", ".secret")

	err := valt.Set("some-key", "some-value 1 ")
	if err != nil {
		panic(err)
	}

	err = valt.Set("some-key2", "some-value 2")
	if err != nil {
		panic(err)
	}
	out, err := valt.Get("some-key")
	if err != nil {
		panic(err)
	}
	fmt.Printf("value: %s\n", out)
	out, err = valt.Get("some-key2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("value: %s\n", out)
}
//package main
//
//import (
//	"fmt"
//
//	"github.com/saltperfect/exercise/secret/encrypt"
//)
//
//func main() {
//
//	plaintext := "hello"
//	key := "this is my secret key"
//	cipher, err := encrypt.Encrypt(key, plaintext)
//	must(err)
//	fmt.Printf("plain: %s cipher: %s\n", plaintext, cipher)
//
//	decryptPlain, err := encrypt.Decrypt(key, cipher)
//	must(err)
//	fmt.Printf("plain: %s, decrypted: %s\n", plaintext, decryptPlain)
//}
//
//func must(err error) {
//	if err != nil {
//		panic(err)
//	}
//}