#pragma once
extern "C"
{
	//char* MTAESEncrypt_encrypt256(const char* raw, int raw_len, const char* key, int* outLen);
	void pay_encrypt_request_with_random(const char* uuid, const char* time_stamp, char* out_sign, int* out_sign_len);
};