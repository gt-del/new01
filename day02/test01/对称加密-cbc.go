package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// PaddingLastGroup 编写填充函数：1.若最后一个分组字节数不够，则进行填充
//2.若最后一个分组字节数刚好合适，就添加一个新的分组
func PaddingLastGroup(plainText []byte, blockSize int) []byte {
	padNum := blockSize - len(plainText)%blockSize
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)    //返回新的byte切片
	newPlain = append(plainText, newPlain...) //注意此处三个点的使用，相当于把byte切片打散
	return newPlain
}

//去掉填充的数据
func unPaddingLastGroup(plainText []byte) []byte {
	length := len(plainText)
	lastChar := plainText[length-1]
	number := int(lastChar)
	return plainText[:length-number]
}

//des加密
func desEncrypt(plainText, key []byte) []byte {
	block, err := des.NewCipher(key) //使用des加密，得到cipher.Block接口，先将des和key绑定
	if err != nil {
		panic(err)
	}
	newText := PaddingLastGroup(plainText, block.BlockSize())
	//选择分组模式,使用cbc
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv) //选定cbc模式加密
	dst := make([]byte, len(newText))
	blockMode.CryptBlocks(dst, newText)
	return dst
}

//des解密
func desDecrypt(ciphertext, key []byte) []byte {
	block, err := des.NewCipher(key) //将des和key绑定
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	blockMod := cipher.NewCBCDecrypter(block, iv) //选定cbc模式解密
	blockMod.CryptBlocks(ciphertext, ciphertext)
	return unPaddingLastGroup(ciphertext)
}
