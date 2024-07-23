package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"

	"crypto/rand"
	"fmt"
)

func main() {

	//GetConnectedPort()
	//TestDummyData()

	TestSaleTransaction()
}

func Emain() {
	unencrypted := "34343501013102023031030C303030303035303030303030040431323334050C313233343536373839303132060C4A4B542D3030303030303031070C2020202020202020202020203030303030303030"

	enc, err := GetAESEncrypted(unencrypted)
	if err != nil {
		panic(err)
	}

	fmt.Println(enc)
}

func GetAESEncrypted(plaintext string) (string, error) {
	key := "61626364656667686162636465666768"
	iv := "my16digitIvKey12"

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	for i := 0; i < len(ciphertext); i++ {
		fmt.Println(ciphertext[i])
	}

	str := base64.StdEncoding.EncodeToString(ciphertext)
	return str, nil
}

func Fmain() {
	key, err := hex.DecodeString("61626364656667686162636465666768") // Hex encoded key
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := hex.DecodeString("37323501013102023031030C303030303035303030303030040431323334050C313233343536373839303132060C4A4B542D3030303030303031070C2020202020202020202020203030303030303030") // Hex encoded data
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv) // Use rand.Read directly
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a CBC encrypter
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the data
	ciphertext := make([]byte, len(data)+aes.BlockSize)
	encrypter.CryptBlocks(ciphertext[aes.BlockSize:], data)

	// Prepend the IV to the ciphertext
	copy(ciphertext[:aes.BlockSize], iv)

	// Encode the encrypted data in hexadecimal format
	encryptedHex := hex.EncodeToString(ciphertext)
	fmt.Println("Encrypted data:", encryptedHex)
}
