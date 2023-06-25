package utils

import (
	facade2 "advt/internal/facade"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var secretKey string

func init() {
	var facade = new(facade2.Facade)
	secretKey = facade.GetAppKey()
}

type JsonHelper struct{}

func Encrypt(content string, secret ...string) (string, error) {
	key := []byte(resoleSecret(secret))
	if len(key) > 16 {
		key = key[:16]
	}
	plaintext := []byte(content)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = padding(plaintext, blockSize)
	if len(plaintext)%aes.BlockSize != 0 {
		return "", nil
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

func Decrypt(content string, secret ...string) (string, error) {
	key := []byte(resoleSecret(secret))
	ciphertext, _ := base64.StdEncoding.DecodeString(content)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		//panic("ciphertext too short")
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = unPadding(ciphertext)
	return string(ciphertext), nil
}

func getSecretKey() string {
	if secretKey == "" {
		secretKey = "abcdefghijklmnop"
	}
	return secretKey
}

func padding(plainText []byte, blockSize int) []byte {
	isPadding := blockSize - len(plainText)%blockSize
	char := []byte{byte(isPadding)}
	newPlain := bytes.Repeat(char, isPadding)
	return append(plainText, newPlain...)
}

func unPadding(plainText []byte) []byte {
	length := len(plainText)
	lastChar := plainText[length-1]
	isPadding := int(lastChar)
	return plainText[:length-isPadding]
}

// 用...传参时一个数组形态
func resoleSecret(secret []string) string {
	switch len(secret) {
	case 0:
		return getSecretKey()
	case 1:
		return secret[0]
	default:
		//默认使用
		return getSecretKey()
	}

}
