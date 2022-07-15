package main

import (
	"crypto/aes"
	"crypto/cipher"
)

//des加密
func aesEncrypt(plainText, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//选择分组模式,使用cbc
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, plainText)
	return plainText
}

//des解密
func aesDecrypt(ciphertext, key []byte) []byte {
	block, err := aes.NewCipher(key) //将des和key绑定
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv) //选定cbc模式解密
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext
}
