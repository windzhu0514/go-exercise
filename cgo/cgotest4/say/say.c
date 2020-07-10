#include<stdio.h>
#include<string.h>
#include "say.h"
#include "cat/cat.h"

void saySomething(char* s){
    printf("i am c,i say:%s\n",s);
    char ss[] = " c has known. hahahhahaha";
    strcat( s,ss );

    miaomiao();
}