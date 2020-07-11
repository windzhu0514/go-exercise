package crypto

import "crypto/rc4"

func RC4Encrypt(key, src []byte) []byte {
	cipher, _ := rc4.NewCipher(key)
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return dst
}
