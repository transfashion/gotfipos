package edcmega

// -- https://gist.github.com/brettscott/2ac58ab7cb1c66e2b4a32d6c1c3908a7?permalink_comment_id=3388968

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/mergermarket/go-pkcs7"
)

// Cipher key must be 32 chars long because block size is 16 bytes

// Encrypt encrypts plain text string into cipher text string
func Encrypt(unencrypted string, keydata string) (string, error) {

	// fmt.Println(ck)

	key := []byte(keydata)
	plainText := []byte(unencrypted)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}
