+++
tags = ["c"]
categories = ["code"]
title = "C Program Operator"
draft = true
+++

### Check whether the given Number is Palindrome or not using Bitwise Operator		

 Code Sample 
```c
/*
* C Program to Check whether the given Number is Palindrome 
* or not using Bitwise Operator
*/
#include <stdio.h>
#include <string.h>
#define SIZE 8
/* Function Prototype */
int is_palindrome(unsigned char[]);

void main()
{
    int num, num1 = 0, i = 0, j = SIZE - 1, res;
    unsigned char c[SIZE];
    printf("Enter a number(max 255)");
    scanf("%d", &num);
    num1 = num;
    while (num != 0)
    {
        c[j] = num&1;
        j--;
        num = num>>1; /* Shifting right the given number by 1 bit */
    }
    printf("The number %d in binary is:", num1);
    for (i = 0;i < SIZE;i++)
    {
        printf("%d", c[i]);
    }
    res = is_palindrome(c);    /* Calling Function */
    if (res == 0)
    {
        printf("\nNUMBER IS PALINDROME\n");
    }
    else
    {
        printf("\nNUMBER IS NOT PALINDROME\n");
    }
}

/* Code to check if the number is palindrome or not */
int is_palindrome(unsigned char c[])
{
    char temp[SIZE];
    int i, j, flag = 0;
    for (i = 0, j = SIZE - 1;i < SIZE, j >= 0;i++, j--)
    {
        temp[j] = c[i];
    }
    for (i = 0;i < SIZE;i++)
    {
        if (temp[i] != c[i])
        {
            flag = 1;
        }
    }
    return flag;
}
```

 Output 
```bash

$ cc  bits21.c
$ a.out
Enter a number
(
max 
255) 153

The number 
153 in  binary is:
10011001

NUMBER IS PALINDROME

$ a.out
Enter a number
(
max 
255) 24

The number 
24 in  binary is:00011000
NUMBER IS PALINDROME
```
### Perform Addition Operation Using Bitwise Operators		

 Code Sample 
```c
#include<stdio.h>

int bitwiseadd(int x, int y)
{
    while (y != 0)
    {
        int carry = x & y;
        x = x ^ y; 
        y = carry << 1;
    }
    return x;
}

int main()
{
    int num1, num2;
    printf("\nEnter two numbers to perform addition using bitwise operators: ");
    scanf("%d%d", &num1, &num2);
    printf("\nSum is %d", bitwiseadd(num1, num2));
    return 0;
}
```

 Output 
```bash

$ gcc  bitwiseadd.c 
-o
 bitwiseadd

$ ./bitwiseadd

Enter two numbers to perform addition using bitwise operators: 
20
 
12
 
Sum is 
32
```
### Swap two Numbers using Bitwise Operators		

 Code Sample 
```c
/*
* C Program to Swap two Numbers using Bitwise operators
*/
#include <stdio.h>
#include <string.h>

/* Function Prototype */
void swap(int*, int *);

void main()
{
    int num1, num2;
    printf("\nEnter two numbers:");
    scanf("%d %d", &num1, &num2);
    printf("\nThe numbers before swapping are Number1= %d Number2 = %d", num1, num2);
    swap(&num1, &num2);        /* Call by Reference to function swap */
    printf("\nThe numbers after swapping are Number1= %d Number2 = %d", num1, num2);
}

/* Code to swap two numbers using bitwise operator */
void swap(int *x, int *y)
{
    *x = *x ^ *y;
    *y = *x ^ *y;
    *x = *x ^ *y;
}
```

 Output 
```bash

$ cc  bit27.c
$ a.out

Enter two numbers:
45
 
76

The numbers before swapping are 
Number1 = 45
 
Number2 = 76

The numbers after swapping are 
Number1 = 76
 
Number2 = 45
```
