#include <zlib.h>
#include<stdio.h>
#include<string.h>
#include<string>
#include<stdlib.h>
#include "polarssl/aes.h"
#include "polarssl/base64.h"
#include "aes_wrap.h"
void aes_cbc_crypt_auto_padding(unsigned char* raw, unsigned int rawlen, unsigned char** out, unsigned int* outlen, unsigned int keylen, unsigned char* key, unsigned char* iv)
{
	aes_context ctx = { 0 };
	int padding_len = 16 - rawlen % 16;
	int after_padding = rawlen + padding_len;
	*outlen = after_padding;
	*out = (unsigned char*)malloc(after_padding);

	unsigned char* after = (unsigned char*)malloc(after_padding);
	memcpy(after, raw, rawlen);
	memset(after + rawlen, padding_len, padding_len);

	aes_setkey_enc(&ctx, key, keylen);
	aes_crypt_cbc(&ctx, AES_ENCRYPT, after_padding, iv, after, *out);
	free(after);
	*outlen = after_padding;
}

void aes_cbc_decrypt(unsigned char* raw, unsigned int rawlen, unsigned char** out, unsigned int* outlen, unsigned int keylen, unsigned char* key, unsigned char* iv)
{
	aes_context ctx = { 0 };
	*out = (unsigned char*)malloc(rawlen);
	aes_setkey_dec(&ctx, key, keylen);
	aes_crypt_cbc(&ctx, AES_DECRYPT, rawlen, iv, raw, *out);
	if ((*out)[rawlen - 1] <= 0x10)
	{
		*outlen = rawlen - (*out)[rawlen - 1];
	}
}
