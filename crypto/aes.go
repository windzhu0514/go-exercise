package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"go-exercise/crypto/ecb"
)

var commonIV = []byte{0x5a, 0xe3, 0xf0, 0x46, 0xcc, 0x11, 0xb4, 0x45, 0x09, 0x04, 0x47, 0x58, 0x00, 0xbf, 0x88, 0xd5}

func AESEncrypt(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = PKCS7Padding(src, block.BlockSize())
	dst := make([]byte, len(src))
	blockMode := cipher.NewCBCEncrypter(block, commonIV)
	blockMode.CryptBlocks(dst, src)

	return dst, nil
}

func AESDecrypt(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(src)%block.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	dst := make([]byte, len(src))
	blockMode := cipher.NewCBCDecrypter(block, commonIV)
	blockMode.CryptBlocks(dst, src)

	dst = PKCS7UnPadding(dst)

	return dst, nil
}

func AESECBEncrypt(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = PKCS5Padding(src, block.BlockSize())
	dst := make([]byte, len(src))
	blockMode := ecb.NewECBEncrypter(block)
	blockMode.CryptBlocks(dst, src)

	return dst, nil
}

func AESECBDecrypt(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(src)%block.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	dst := make([]byte, len(src))
	blockMode := ecb.NewECBDecrypter(block)
	blockMode.CryptBlocks(dst, src)
	dst = PKCS5UnPadding(dst)

	return dst, nil
}
