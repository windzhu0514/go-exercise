#include<cstdlib>
#include<cstdio>
#include<cstring>
#include"polarssl/sha1.h"
#include "polarssl/base64.h"
#include "sign.h"



int sign_internal(char *key, int key_len, const char *raw, int raw_len, char *out)
{
	char v_value; // r0
	char v_array_0[64] = { 0 }; // [sp-C0h] [bp-110h]
	char v_array_1[64] = { 0 }; // [sp-80h] [bp-D0h]
	char v_array_2[64] = { 0 }; // [sp-40h] [bp-90h]	
	char sha1_out[20] = { 0 }; // [sp+34h] [bp-1Ch]
	char last_data_to_sha1[88] = { 0 };

	if (key == 0 || key_len == 0 || raw == 0 || raw_len == 0 || out == 0) {
		return -1;
	}

	memset(v_array_2, 0x40, 0);
	memcpy(v_array_2, key, key_len);

	for (int v_array_idx = 0; v_array_idx < 0x40; v_array_idx++) {
		v_value = v_array_2[v_array_idx];
		v_array_1[v_array_idx] = (~v_value & 0xF2 | v_value & 0xD) ^ 0xAE;
		v_array_0[v_array_idx] = ~v_value & 0x36 | v_value & 0xC9;
	}
		
	int total_size = raw_len + 64;
	char* v_malloc = (char *)malloc(total_size);
	memcpy(v_malloc, v_array_0, 64);
	memcpy(v_malloc + 64, raw, raw_len);
	sha1((unsigned char*)v_malloc, total_size, (unsigned char*)sha1_out);
	free(v_malloc);

	memcpy(last_data_to_sha1, v_array_1, 0x40);
	memcpy(last_data_to_sha1 + 0x40, sha1_out, 20);
	sha1((unsigned char*)last_data_to_sha1, 84, (unsigned char*)out);

	return 0;
}
void sign(const char* in, size_t in_size, char* out, size_t* out_size) {
	char sha1[0x14] = { 0 };
	char key[] = "Tb6yTwgSEvbLgLtguw21Q80dR8atTLZ9gbOyX3m9FB0FMGWI60SALA==";
	//char key[] = "83c96b209eb9731bab61dd03dc34e1afY4yBJhR5whBO3j8lGOkXJQ==";
	int key_size = 0x38;
	sign_internal(key, key_size, in, in_size, sha1);
	base64_encode((unsigned char*)out, out_size, (unsigned char*)sha1, 0x14);
}