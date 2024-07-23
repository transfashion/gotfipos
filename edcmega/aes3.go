package edcmega

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Encrypt3(unencrypted string, keydata string) string {
	key, err := hex.DecodeString(keydata) // Hex encoded key
	if err != nil {
		panic(err)
	}

	data, err := hex.DecodeString(unencrypted) // Hex encoded data
	if err != nil {
		panic(err)
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv) // Use rand.Read directly
	if err != nil {
		panic(err)
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
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

	return encryptedHex
}

func Encrypt4(unencrypted string, keydata string) string {

	key, err := hex.DecodeString(keydata) // Hex encoded key
	if err != nil {
		panic(err)
	}

	data, err := hex.DecodeString(unencrypted) // Hex encoded data
	if err != nil {
		panic(err)
	}
	// data := []byte(unencrypted)
	// key := []byte(keydata)

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Create a CBC encrypter (no padding)
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the data without padding
	ciphertext := make([]byte, len(data)) // Allocate ciphertext size without padding
	encrypter.CryptBlocks(ciphertext, data)

	// Prepend the IV to the ciphertext
	combined := append(iv, ciphertext...)

	// Encode the encrypted data in hexadecimal format
	encryptedHex := hex.EncodeToString(combined)
	fmt.Println("Encrypted data (no padding):", encryptedHex)

	return encryptedHex
}

func Encrypt5(plaintext string, key string) []byte {

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
		panic(err)
	}

	ciphertext := make([]byte, len(plainTextBlock))

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	return ciphertext
	/*
		hex := make([]string, len(ciphertext))
		for i := 0; i < len(ciphertext); i++ {
			// fmt.Println(ciphertext[i])
			hex[i] = fmt.Sprintf("%02X", ciphertext[i])
		}

		fmt.Println(hex)
		return strings.Join(hex[:], "")
	*/

}

//    162BA0C6D18C8176745BB08B7A281214EA3EADFE4A17CFEA591D3C2788983D69FC1D5C83EA6D312934AEFD4A6E0B3D65C204D5A247E0B3C07476C1801280F8FB370861943EC71D0B951346CB84FD236B0825FF5C75917DFC30165B64CADF5A9F9B251E1DB0FE45DA3E5ED1106C11D78C826EF8E16D1486D0B7CADB82A48F266C0666386573F451EA8F4409851241F24CADC75008EBC37DFEFABDF7B625407600
//0050904C126B3DB30916B05FCEE2CE90A73FEE83D0293E9B24DA9E0A46FB9FF1ADE96A6FE86EA59E75E14CD45B49DF10C2976AA778DA65D3E01BBBC95D95EAB79C96AF82780A28A87B2CA9CE10865A857191
