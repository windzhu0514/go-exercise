#pragma once

#if defined(__cplusplus)
extern "C" {
#endif
	/* 
	* @param in:  the input data.
	* @param in_size: size of input data.
	* @param out:  algorithm out data.  reserved 40 bytes.
	* @param out_size: size of out data.
	* @function: to get the sign of the url. 
	*/
	void sign(const char* in, size_t in_size, char* out, size_t* out_size);//skcy sign

	/*
	* @param in:  the input data (collected device info).
	* @param in_size: size of input data.
	* @param out:  algorithm out data.  reserved double length of input data.
	* @param out_size: size of out data.
	* @function: to get encrypted device info, named siua.
	*/
	void siua(const char* in, size_t  in_size,char* out, size_t * out_size); //siua encryption

	/*
	* @param in:  the input data (collected json info ).
	* @param in_size: size of input data.
	* @param out:  algorithm out data.  reserved double length of input data.
	* @param out_size: size of out data.
	* @function: to get encrypted json info, named siua.
	*/
	//https://appsec-mobile.meituan.com/sign/v2?__skck=8f5973b085446090f224af74e30e0181&__skts=1561614838&__skua=d41d8cd98f00b204e9800998ecf8427e&__skno=7c730d56-6708-4afb-a102-f5012d40cfe5&__skcy=4GUA3kMa99EpGAz2nnHm85UVMUA%3D
	void encyptDataAES2(const char* in, size_t  in_size, char* out, size_t * out_size);//dfp post data encryption
	
	/*
	* just for test 
	*/
	void decyptDataAES2(const char* in, size_t  in_size, char* out, size_t * out_size);//test for dec
	//void siuadec(const char* in, size_t  in_size, char* out, size_t * out_size); 
	//void pay_request_sign(const char* uuid, const char* time_stamp, char* out_sign, int* out_sign_len);// meituan trip pay sign

	/*
	* @param uuid  .
	* @param time_stamp.
	* @param out_sign:  algorithm out data. reserve 400 bytes.
	* @param out_sign_len: size of out data.
	* @function: to get encrypted json info, named siua.
	*/
	void pay_encrypt_request_with_random( const char* uuid, const char* time_stamp, char* out_sign, int* out_sign_len);//meituan app pay sign
	
#if defined(__cplusplus)
}
#endif
