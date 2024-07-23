package edcmega

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Encrypt2(data []byte, key []byte) []byte {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Create a CBC encrypter
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(data)+aes.BlockSize)
	encrypter.CryptBlocks(ciphertext[aes.BlockSize:], data)

	// Prepend the IV to the ciphertext
	copy(ciphertext[:aes.BlockSize], iv)

	//fmt.Println("Encrypted data:", ciphertext)

	return ciphertext
}
