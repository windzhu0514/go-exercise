#include <zlib.h>
#include<stdio.h>
#include<string.h>
#include<string>
#include<stdlib.h>
#include "polarssl/aes.h"
#include "polarssl/base64.h"
#include"sign.h"
#include "aes_wrap.h"



std::string Base64Encode(std::string& in, bool with_new_line)
{
	size_t base64len = 0;
	base64_encode(NULL, &base64len, (unsigned char*)in.c_str(), in.length());
	char*  base64buf = (char*)malloc(base64len);
	base64_encode((unsigned char*)base64buf, &base64len, (unsigned char*)in.c_str(), in.length());

	std::string base = base64buf;
	free(base64buf);
	return base;
}


std::string siua(const std::string& src)
{
	int boundLen = compressBound(src.length());
	unsigned char* zipBody = (unsigned char*)malloc(boundLen);
	int            zipLen = boundLen;
	compress((Bytef*)zipBody, (uLongf*)&zipLen, (Bytef*)src.c_str(), src.length());
	unsigned char key[] = "meituan0sankuai1";
	unsigned char iv[] = "0102030405060708";
	unsigned char* aes_out = NULL;
	unsigned int            aes_len = 0;


	/*unsigned char data[459] = {
		0x78, 0x9C, 0x5D, 0x51, 0xDB, 0x6E, 0xE2, 0x30, 0x10, 0xFD, 0x15, 0x3F, 0xEE, 0xAE, 0xB0, 0x33,
		0x76, 0xEE, 0x7E, 0x0B, 0x17, 0xF5, 0x61, 0x09, 0x20, 0x40, 0x5D, 0xF5, 0xA9, 0x32, 0x89, 0xA3,
		0x46, 0x85, 0xB0, 0x4A, 0x80, 0xAA, 0x68, 0xFC, 0xEF, 0x3B, 0x21, 0x55, 0x4B, 0xD7, 0x76, 0x62,
		0xC7, 0x73, 0x72, 0xCE, 0x99, 0x19, 0x29, 0xC0, 0xB9, 0x3C, 0xC2, 0xDC, 0xD6, 0xD7, 0x33, 0x1E,
		0xFA, 0xF7, 0x73, 0x1E, 0x3D, 0x4F, 0x16, 0x9F, 0x67, 0x5C, 0xAC, 0xA7, 0x29, 0xE4, 0x78, 0x7D,
		0x41, 0xBA, 0x1D, 0x70, 0x74, 0x1B, 0x0B, 0x40, 0x15, 0x60, 0x6B, 0xF7, 0xD6, 0x74, 0x96, 0xBF,
		0xDA, 0xF7, 0x6E, 0x08, 0x7A, 0x77, 0x24, 0x9F, 0x67, 0x4D, 0x70, 0x6F, 0x20, 0xF2, 0x64, 0x28,
		0x15, 0x48, 0x5F, 0x25, 0xA9, 0x3E, 0x77, 0xB6, 0xF5, 0xBE, 0x51, 0x1C, 0x4E, 0x51, 0x1C, 0x86,
		0x98, 0x5F, 0xF9, 0xF8, 0x5C, 0xEF, 0x4B, 0xDB, 0xF2, 0xB9, 0x4C, 0xB0, 0xC7, 0xF5, 0xA2, 0xA6,
		0x3D, 0x44, 0x01, 0xBF, 0x24, 0x06, 0xEF, 0xAD, 0xF2, 0x3E, 0xCC, 0x48, 0x81, 0x0D, 0x0A, 0xEC,
		0x4B, 0x81, 0x7D, 0x23, 0x97, 0x08, 0xE8, 0xDC, 0xE1, 0xF4, 0x77, 0x64, 0xCA, 0x1D, 0xFE, 0xBF,
		0xE7, 0xCB, 0xF9, 0x93, 0x98, 0xAF, 0xA5, 0x14, 0x7F, 0x64, 0xE4, 0x83, 0xC8, 0xA7, 0x22, 0x5F,
		0x89, 0x47, 0xE5, 0x8B, 0x50, 0xAC, 0x7C, 0x35, 0x62, 0xC4, 0x19, 0x7B, 0x90, 0x7A, 0x2A, 0x66,
		0x20, 0xB5, 0x1F, 0xA3, 0x69, 0xCA, 0xF6, 0x58, 0x97, 0x24, 0x52, 0xD9, 0xD6, 0x36, 0x85, 0xE5,
		0x6D, 0xBD, 0x67, 0x92, 0x2A, 0xF3, 0xB6, 0x37, 0x0D, 0x60, 0x36, 0xDE, 0xCC, 0x16, 0xDB, 0xD1,
		0xB0, 0x21, 0x89, 0x03, 0x59, 0xF8, 0x9A, 0x1F, 0x5F, 0xCE, 0x65, 0x59, 0x5B, 0xBC, 0x44, 0x01,
		0x5B, 0xB5, 0xC7, 0xC2, 0x76, 0xDD, 0xB1, 0x25, 0xC6, 0x0B, 0x53, 0xEC, 0x87, 0x31, 0xB7, 0xC0,
		0x4F, 0xCC, 0xB7, 0x54, 0x16, 0x78, 0xF4, 0xA6, 0x0B, 0x4C, 0x30, 0x9B, 0x4C, 0x66, 0xF3, 0xD9,
		0x7A, 0x99, 0xCF, 0xB6, 0xB3, 0x35, 0x85, 0x7E, 0xE3, 0x43, 0x6B, 0x2E, 0xF5, 0xE9, 0x9D, 0x6D,
		0x6C, 0x43, 0x7F, 0x63, 0xB6, 0xDC, 0xAC, 0x9C, 0x4B, 0xA2, 0x58, 0x26, 0x0A, 0xFC, 0x24, 0x4D,
		0x20, 0xA5, 0x56, 0x71, 0x9A, 0x52, 0x25, 0xF0, 0x2B, 0x56, 0x80, 0xD4, 0xBB, 0x87, 0x31, 0x82,
		0xD2, 0x00, 0xF7, 0x0B, 0xF1, 0xAD, 0xAE, 0x6A, 0xE7, 0xF8, 0xCD, 0x5C, 0x6F, 0x98, 0xA3, 0x8D,
		0x54, 0x0A, 0x61, 0x10, 0x71, 0xA5, 0xD2, 0x92, 0x07, 0x31, 0x48, 0x9E, 0x96, 0x3B, 0xC9, 0x8D,
		0xDA, 0xA9, 0x54, 0xC5, 0x7E, 0x0C, 0xA5, 0x75, 0x6E, 0x40, 0x67, 0x34, 0x28, 0x9F, 0xA1, 0x2E,
		0x58, 0x1C, 0x0F, 0xA2, 0x33, 0xCD, 0xEB, 0xD9, 0xD4, 0x82, 0xFA, 0x75, 0x3A, 0x9B, 0x06, 0x25,
		0x08, 0x10, 0x01, 0x48, 0xBC, 0x39, 0xA2, 0x92, 0xA6, 0x1C, 0x88, 0x3A, 0x62, 0x32, 0xD0, 0x0A,
		0x74, 0x98, 0x68, 0x08, 0x13, 0xE7, 0x08, 0x85, 0xFD, 0x33, 0x7F, 0xE2, 0xC7, 0xAA, 0xAA, 0x0B,
		0x8B, 0x26, 0xD0, 0x65, 0xA1, 0x77, 0x56, 0x57, 0xA0, 0x6D, 0xA8, 0xCB, 0x98, 0x0C, 0xF2, 0x28,
		0xBC, 0xA5, 0xC5, 0xA9, 0xAD, 0xFF, 0x00, 0x09, 0x64, 0xE2, 0xB2
	};*/
	aes_cbc_crypt_auto_padding(zipBody, zipLen, &aes_out, &aes_len, 128, key, iv);
	//aes_cbc_crypt_auto_padding(data, sizeof(data), &aes_out, &aes_len, 128, key, iv);

	std::string str(aes_out, aes_out + aes_len);
	std::string base=Base64Encode(str, false);
	free(zipBody);
	free(aes_out);
	return base;
}

void siua(const char* in, size_t  in_size, char* out, size_t * out_size)
{
	std::string ss(in,in+in_size);
	std::string ret = siua(ss);
	strcpy(out,ret.c_str());
	*out_size = ret.length();
}

void siuadec(const char* in, size_t  in_size, char* out, size_t * out_size)
{
	unsigned char* basedecode = (unsigned char*)malloc(in_size);
	size_t baselen = in_size;
	base64_decode(basedecode, &baselen, (unsigned char*)in, in_size);
	unsigned char* gziped = NULL;
	unsigned int gzipedlen = 0;
	unsigned char key[] = "meituan0sankuai1";
	unsigned char iv[] = "0102030405060708";
	aes_cbc_decrypt(basedecode, baselen, &gziped, &gzipedlen, 128, key, iv);
	
	int osize = compressBound(gzipedlen);
	FILE* fp=fopen("D://zippedsiua.txt","wb");
	fwrite(gziped,1,gzipedlen,fp);
	fclose(fp);

	uncompress((Bytef*)out,(uLongf*)&osize,gziped,gzipedlen);
}



