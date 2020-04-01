#pragma once
void aes_cbc_crypt_auto_padding(unsigned char* raw, unsigned int rawlen, unsigned char** out, unsigned int* outlen, unsigned int keylen, unsigned char* key, unsigned char* iv);

void aes_cbc_decrypt(unsigned char* raw, unsigned int rawlen, unsigned char** out, unsigned int* outlen, unsigned int keylen, unsigned char* key, unsigned char* iv);