package main

import (
	"fmt"

	"github.com/sambit81/gocrypto/decrypt"
	"github.com/sambit81/gocrypto/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
