// rsaDemo.cpp : 定义控制台应用程序的入口点。
//

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "librsasign.h"


int main()
{
	int *out_buf;
	int out_size;
	for (int i = 0; i < 1; i++) {
		const char* src = "EOqXguubro6Q8LS3P7SYvw==";
		int ret = getSign(src, &out_buf, &out_size);
		std::string str = bin2hex((unsigned char*)out_buf, out_size * 4);
		free(out_buf);
		printf("src:%s\nresult = %d %s\n", src, out_size, str.c_str());
	}

    return 0;
}
