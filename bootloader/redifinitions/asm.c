#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char [200] concat(char * first,int number){

    char res[200],num[10];

    strcpy(res,first);
    sprintf(num,"%d",number);

    strcat(res,num);

    return res;

}

void __ggl_jump0(int line){

    __asm__(concat("jmp",line));

}