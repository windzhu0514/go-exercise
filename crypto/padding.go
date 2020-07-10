package crypto

import "bytes"

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// PKCS7Padding 和 PKCS5Padding 填充方式一样
// PKCS7Pad() pads an byte array to be a multiple of 16
// http://tools.ietf.org/html/rfc5652#section-6.3
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding 和 PKCS5UnPadding 去填充方式一样
// PKCS7Unpad() removes any potential PKCS7 padding added.
func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	trailing := length - unpadding
	if trailing >= 0 && trailing < length {
		return src[:(length - unpadding)]
	}

	return nil
}
