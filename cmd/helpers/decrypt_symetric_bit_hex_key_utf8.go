package helpers

import (
	"crypto/aes"
	"log"
)

const (
	alg = "aes-256-gcm"
)

func DecryptSymmetricBiHexKeyUTF8(k string) {
	key := []byte(k)

	_, err := aes.NewCipher(key)

	if IsError(err) {
		log.Fatal(err)
	}

	// d, err := cipher.NewGCMWithNonceSize(b, 16)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// txt ,err :=

}
