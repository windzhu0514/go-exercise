#pragma once
#include <string>

int getSign(const char *in, int** out_buf, int* out_size);
std::string bin2hex(unsigned char *bin, size_t len);
