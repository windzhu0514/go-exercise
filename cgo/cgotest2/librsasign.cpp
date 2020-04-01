//#include "stdafx.h"
#include "idamacro.h"
#include <inttypes.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <string>
#include "librsasign.h"

int sum;
//uint32 unk_640C;

std::string bin2hex(unsigned char *bin, size_t len)
{
	std::string res;
	size_t  i;

	if (bin == NULL || len == 0)
		return NULL;

	for (i = 0; i < len; i++) {
	    printf("%d ",bin[i]);
		res += "0123456789ABCDEF"[bin[i] >> 4];
		res += "0123456789ABCDEF"[bin[i] & 0x0F];
	}

	printf("\n");

	return res;
}

signed __int64 rsa_modExp(__int64 a1, signed __int64 a2, signed __int64 a3)
{
	//printf("Before new_rsa_modExp %" PRIx64 " %" PRIx64 " %" PRIx64 "\n", a1, a2, a3);

	signed __int64 v3; // r6
	__int64 v4; // r2
	signed __int64 result; // r0
	__int64 v6; // r4
	signed __int64 v7; // r0
	signed __int64 v8; // r2
	int v9; // t2
	signed __int64 v10; // r0
	signed __int64 v11; // r0
	int v12; // r3

	v3 = a2;
	if (a2 >= 258)
		++sum;
	if ((HIDWORD(a1) | HIDWORD(a2)) < 0 || a3 < 1)
	{
		//_android_log_print(3, "myprint", "exit(1)");
		exit(1);
	}
	if (!a2) {
		//printf("After new_rsa_modExp %" PRIx64 "\n", 1LL);
		return 1LL;
	}

	v4 = a1 % a3;
	result = 1LL;
	v6 = v4;
	if (v3 == 1) {
		//printf("After new_rsa_modExp %" PRIx64 "\n", v6);
		return v6;
	}
	if (!(v3 & 1))
	{
		v7 = (unsigned int)v4 * (unsigned __int64)(unsigned int)v4;
		HIDWORD(v7) += 2 * v4 * HIDWORD(v4);
		HIDWORD(v8) = SHIDWORD(v3) >> 1;
		v9 = LODWORD(v3) >> 1;
		LODWORD(v8) = v9;
		v10 = rsa_modExp(v7 % a3, v8, a3);
		//printf("After new_rsa_modExp %" PRIx64 "\n", v10 % a3);
		return v10 % a3;
	}
	if (v3 & 1)
	{
		v11 = rsa_modExp(v4, v3 - 1, a3);
		v12 = v6 * HIDWORD(v11) + v11 * HIDWORD(v6);
		v10 = (unsigned int)v6 * (unsigned __int64)(unsigned int)v11;
		HIDWORD(v10) += v12;
		//printf("After new_rsa_modExp %" PRIx64 "\n", v10 % a3);
		return v10 % a3;
	}

	//printf("After new_rsa_modExp %" PRIx64 "\n", result);
	return result;
}


_QWORD *rsa_encrypt(unsigned __int8 *a1, unsigned int a2, _QWORD *a3)
{
	//printf("Before new_rsa_encrypt %d %s %" PRIx64 " %" PRIx64 "\n", a2, a1, a3[0], a3[1]);

	unsigned __int8 *v3; // r6
	unsigned int v4; // r7
	signed __int64 *v5; // r4
	_QWORD *v6; // r0
	_QWORD *v7; // r5
	signed __int64 *v8; // r8
	signed __int64 v9; // r10
	unsigned int v10; // t1

	v3 = a1;
	v4 = a2;
	v5 = (signed __int64 *)a3;
	v6 = (_QWORD *)malloc(8 * a2);
	v7 = v6;
	if (v6)
	{
		v8 = (signed __int64 *)v6;
		v9 = 0LL;
		while (v9 < v4)
		{
			v10 = *v3++;
			++v9;
			*v8 = rsa_modExp(v10, v5[1], *v5);
			++v8;
		}
	}
	//std::string str = bin2hex((unsigned char*)v7, a2 * 8);
	//printf("After new_rsa_encrypt %s\n", str.c_str());
	return v7;
}

int Envrypt2(const char *a1, const char *a2, int** out_buf, int* out_size)
{
	int v2; // r6
	const char *v3; // r11
	const char *v4; // r8
	size_t size_a1; // r9
	size_t size_a2; // r0
	size_t size; // r5
	char *buf; // r0
	int v10; // r1
	_QWORD *v11; // r6
	_QWORD *v12; // r0
	void *v13; // r9
	_QWORD *v14; // r0
	void *v15; // r5
	int *v16; // r11
	int *v17; // r8
	int v18; // r3
	int v19; // t1
	size_t v20; // r0
	signed int result; // r0
	//unsigned __int8 v22; // vf
	unsigned int v23; // [sp+0h] [bp+0h]
	//size_t n; // [sp+8h] [bp+8h]
	int c; // [sp+Ch] [bp+Ch]
	//int v27; // [sp+14h] [bp+14h]
	char buf2[1024] = { 0 };
	int* acSecret_Text = NULL;

//	v2 = 1 - (_DWORD)a1;
	v3 = a1;
//	if ((unsigned int)a1 > 1)
		v2 = 0;
	if (!a2)
		v2 |= 1u;
	v4 = a2;
	c = v2;
	if (v2)
		goto LABEL_16;
	size_a1 = strlen(a1);
	size_a2 = strlen(a2);
	size = size_a1 + size_a2 + 1;
	buf = (char *)malloc(size);
	memcpy(buf, a1, size_a1);
	memcpy(buf + size_a1, a2, size_a2);
	buf[size - 1] = 0;
	v10 = c;
	acSecret_Text = (int*)malloc(600);
	memset(acSecret_Text, v10, 600);
	memset(buf2, c, size);
	memcpy(buf2, buf, size);
	v11 = (_QWORD *)malloc(0x10u);
	v12 = (_QWORD *)malloc(0x10u);
	*v11 = (_QWORD )0xB5547;
	v11[1] = 257LL;
	*v12 = (_QWORD )0xB5547;
	v13 = v12;
	v12[1] = (_QWORD )0x98BF1;
	v14 = rsa_encrypt((unsigned __int8 *)buf2, size, v11);
	v15 = v14;
	if (v14) {
		v16 = (int *)v14;
		v17 = acSecret_Text;
		v18 = c;
		v20 = strlen(buf2);
		while (1) {
			v23 = v18;
			if (v23 >= v20)
				break;
			v19 = *v16;
			v16 += 2;
			v18 = v23 + 1;
			v17[0] = v19;
			++v17;
		}
		free(v15);
		free(buf);
		free(v11);
		free(v13);
		result = 0;
	}
	else
	{
	LABEL_16:
		result = 1;
	}
	//v22 = __OFSUB__(v27, *v24);
	*out_buf = acSecret_Text;
	*out_size = size - 1;

    //for (int i = 0; i < size - 1; i++) {
	//	printf("%d ",*(acSecret_Text+i) );
	//}

	//printf("\n");

	std::string str = bin2hex((unsigned char *)acSecret_Text, size - 1);
	printf("result = %d %s\n", size - 1, str.c_str());

	return result;
}

int getSign(const char *in, int** out_buf, int* out_size) {
	return Envrypt2(in, "com.htinns", out_buf, out_size);
}