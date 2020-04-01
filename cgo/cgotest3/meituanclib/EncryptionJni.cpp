#include <stdlib.h>
#include <polarssl/base64.h>
#include "aes_wrap.h"
#include "EncryptionJni.h"
void encyptDataAES2(const char* in, size_t  in_size, char* out, size_t * out_size)
{
	//unsigned char key[] = "Jljvk2z0ZcK20p7hpQBQqeQCXyYa3kov";//meituan1sankuai0
	unsigned char key[] = "meituan1sankuai0";
	unsigned char iv[] = "0102030405060708";

	unsigned char*  aesout = NULL;
	unsigned int    aeslen = 0;
	aes_cbc_crypt_auto_padding((unsigned char*)in, in_size, (unsigned char**)&aesout, &aeslen, 128, key, iv);

	if (aeslen <= 0)return;
	char* baseout = (char*)malloc(aeslen * 2);
	size_t baselen = aeslen*2;
	base64_encode((unsigned char*)baseout,&baselen,aesout,aeslen);
	free(aesout);

	strcpy(out, baseout);
	*out_size = baselen;
	free(baseout);
}

void decyptDataAES2(const char* in, size_t  in_size, char* out, size_t * out_size)
{
	unsigned char key[] = "meituan1sankuai0";
	unsigned char iv[] = "0102030405060708";

	char* base64dec = (char*)malloc(in_size);
	size_t baselen= in_size;
	base64_decode((unsigned char*)base64dec, &baselen, (unsigned char*)in, in_size);

	unsigned char*  aesout = NULL;
	unsigned int    aeslen = 0;
	aes_cbc_decrypt((unsigned char*)base64dec, baselen, (unsigned char**)&aesout, &aeslen, 128, key, iv);
	free(base64dec);

	memcpy(out, aesout,aeslen);
	*out_size = aeslen;
	free(aesout);
}