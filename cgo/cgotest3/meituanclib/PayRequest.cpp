//

#include<string>
#include"polarssl/aes.h"
#include"polarssl/sha2.h"
#include"polarssl/base64.h"
#include"polarssl/rsa.h"
#include"polarssl/bignum.h"
#include<stdlib.h>
#include"sign.h"
#include"PayRequest.h"

int rng(void * context, unsigned char * need_pading, size_t need_len)
{
	//memset(need_pading,0xff,need_len);
	int i;
	for ( i = 0; i < need_len; i++)
	{
		need_pading[i] = rand();
	}
	return 0;
}

int MTPKCS5Padding_padding(char* buf, int sz)
{
	int leave = 16 - sz % 16;
	memset(buf + sz, leave, leave);
	return sz + leave;
}

char* MTAESEncrypt_encrypt256(const char* raw, int raw_len, const char* key, int* outLen)
{
	aes_context context;
	if (!raw)
	{
		return 0;
	}
	char* new_raw_buf = (char*)malloc(raw_len+16);
	
	memset(&context,0,sizeof(aes_context));
	memcpy(new_raw_buf,raw,raw_len);
	int real_len = MTPKCS5Padding_padding(new_raw_buf,raw_len);

	char* new_raw_buf2= (char*)malloc(raw_len + 16);
	if (key == NULL)
	{
		key = "";
	}

	aes_setkey_enc(&context,(unsigned char*)key,256);
	char iv[0x20] = { 0 };
	aes_crypt_cbc(&context,AES_ENCRYPT, real_len,(unsigned char*)iv,(unsigned char*)new_raw_buf,(unsigned char*)new_raw_buf2);
	*outLen = real_len;
	free(new_raw_buf);
	return new_raw_buf2;
}	
std::string Base64Encode(std::string& in, bool with_new_line);
int rsa_encrpyted_content(const char* raw, char* outBuf, int* outBufSz)
{

	//char* tmp_raw = (char*)malloc(0x24);
	//memcpy(tmp_raw, raw, 0x20);
	const char* RSA_N = "00a3710b684db11c2d6202eecd7b004e1e79ff8d4f3b903d8c30743029b86c7f532e3bb4b528654dfdaecbd2c9d7e45288ae755f461be1c7e505d3c3af3b389bb5401bc5bacd1561d9aa4ff64ec76fdc66d00b482c2e91f2f067073d8e8d73b2a3d175fad29d8b8b5d2b067ebb2eb9b88fe9bedc842d594f93751bc560fcfe94b959f96adcd1aaf96bcba8069851f986c37864bc42d1d6bc6a898018d2c1c7aa03532d7fe42a7d579b8bf816d222da89d1c714821981d49188c405b126af403c54093823950c257c815a07937afe5846ea39287022deb49ca61e3725a918236ae8cf49afac71371f36bef9aacb2f5f81b374b18e928b8e72265d52338e15e66eb1";
	const char* RSA_E = "10001";

	rsa_context context;
	rsa_init(&context, 0, 0);
	context.len = 256;
	mpi_read_string(&context.N, 16, RSA_N);
	mpi_read_string(&context.E, 16, RSA_E);
	int err=rsa_pkcs1_encrypt(&context, rng, 0, 0, 32, (unsigned char*)raw, (unsigned char*)outBuf);
	rsa_free(&context);
	if (err)
	{
		//free(tmp_raw);
		return -1;
	}
	char basetmp[0x200];
	int  baselen = 0x200;

	std::string ss(outBuf,outBuf+0x100);
	std::string temp=Base64Encode(ss,0);

	//err=base64_encode((unsigned char*)basetmp,(size_t*)&baselen,(unsigned char*) outBuf,0x100);
	//if (err)
	//{
	//	return -1;
	//}
	//basetmp[baselen] = 0;
	//strcpy(outBuf,basetmp);
	//*outBufSz = baselen;


	strcpy(outBuf,temp.c_str());
	*outBufSz=temp.length();

	return 0;
}
int encrypt_request(char** out_enc, const char* uuid, const char* time_stamp, const char** in_array, int in_count)
{
	char hMacKey[0x200] = { 0 };
	strcpy(hMacKey, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo3ELaE2xHC1iAu7NewBOHnn/jU87kD2MMHQwKbhsf1MuO7S1KGVN/a7L0snX5");
	strcat(hMacKey,uuid);
	int hMacKeyLen = strlen(hMacKey);

	char hMacRaw[0x100] = { 0 };
	strcpy(hMacRaw,time_stamp);
	int hMacRawLen = strlen(hMacRaw);

	char hmac[0x20] = { 0 };
	sha2_hmac((unsigned char*)hMacKey,hMacKeyLen, (unsigned char*)hMacRaw,hMacRawLen,(unsigned char*)hmac,0);

	char hMacBase64[0x40] = { 0 };
	int hMacBase64Len = 0x40;
	base64_encode((unsigned char*)hMacBase64, (size_t*)&hMacBase64Len,(unsigned char*)hmac,0x20);
	hMacBase64[hMacBase64Len] = 0;

	out_enc[in_count + 1] = (char*)malloc(0x40);
	strcpy(out_enc[in_count+1],hMacBase64);
	out_enc[in_count + 0] = (char*)malloc(0x200);
	int  rsa_out_len=0x200;
	rsa_encrpyted_content(hmac, out_enc[in_count + 0],&rsa_out_len);

	int i;
	for (i = 0; i < in_count; i++)
	{
		char * aes_out = NULL;
		int    aes_len = 0;
		aes_out=MTAESEncrypt_encrypt256(in_array[i],strlen(in_array[i]),hmac,&aes_len);
		int    baselen = aes_len * 2;
		char * basetmp = (char*)malloc(baselen);
		base64_encode((unsigned char*)basetmp,(size_t*)&baselen, (unsigned char*)aes_out,aes_len);
		out_enc[i] = basetmp;
		free(aes_out);
	}
	return 0;
}

int encrypt_request_with_random(char** out_enc, const char* uuid, const char* time_stamp,const char* rand1,const char * rand2,const char* rand3, const char** in_array, int in_count)
{
	char hMacKey[0x200] = { 0 };
	strcpy(hMacKey, "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo3ELaE2xHC1iAu7NewBOHnn");
	strcat(hMacKey, uuid);
	strcat(hMacKey,rand1);
	int hMacKeyLen = strlen(hMacKey);

	char hMacRaw[0x100] = { 0 };
	strcpy(hMacRaw, rand2);
	strcat(hMacRaw, time_stamp);
	strcat(hMacRaw, rand3);
	
	int hMacRawLen = strlen(hMacRaw);

	char hmac[0x20] = { 0 };
	sha2_hmac((unsigned char*)hMacKey, hMacKeyLen, (unsigned char*)hMacRaw, hMacRawLen, (unsigned char*)hmac, 0);

	char hMacBase64[0x40] = { 0 };
	int hMacBase64Len = 0x40;
	base64_encode((unsigned char*)hMacBase64, (size_t*)&hMacBase64Len, (unsigned char*)hmac, 0x20);
	hMacBase64[hMacBase64Len] = 0;

	out_enc[in_count + 1] = (char*)malloc(0x40);
	strcpy(out_enc[in_count + 1], hMacBase64);
	out_enc[in_count + 0] = (char*)malloc(0x200);
	int  rsa_out_len = 0x200;
	rsa_encrpyted_content(hmac, out_enc[in_count + 0], &rsa_out_len);

	int i;
	for (i = 0; i < in_count; i++)
	{
		char * aes_out = NULL;
		int    aes_len = 0;
		aes_out = MTAESEncrypt_encrypt256(in_array[i], strlen(in_array[i]), hmac, &aes_len);
		int    baselen = aes_len * 2;
		char * basetmp = (char*)malloc(baselen);
		base64_encode((unsigned char*)basetmp, (size_t*)&baselen, (unsigned char*)aes_out, aes_len);
		out_enc[i] = basetmp;
		free(aes_out);
	}
	return 0;
}

void pay_request_sign(const char* uuid, const char* time_stamp, char* out_sign, int* out_sign_len)
{
	char* in_array[1] = {  };
	char* out_array[2];
	encrypt_request((char**)out_array, uuid,time_stamp, (const char**)in_array, 0);
	strcpy(out_sign,out_array[0]);
	*out_sign_len = strlen(out_sign);
	free(out_array[0]);
	free(out_array[1]);
}

void pay_encrypt_request_with_random( const char* uuid, const char* time_stamp, char* out_sign, int* out_sign_len)
{
	char rand1[0x11]="1234567890abcdef";
	char rand2[0x11]="1234567890abcdef";
	char rand3[0x11]="1234567890abcdef";
	char *out_array[2] = { NULL,NULL };
	encrypt_request_with_random((char**)out_array, uuid, time_stamp,rand1,rand2,rand3,NULL, 0);
	strcpy(out_sign, out_array[0]);
	*out_sign_len = strlen(out_sign);
	free(out_array[0]);
	free(out_array[1]);
}