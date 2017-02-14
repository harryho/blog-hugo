#include<stdio.h>

int main( ) 
{
    char var[] = "World";   
    printf( "Hello %s \n Press any key to exit.", var );
    char key = getchar();
    return 0;
}